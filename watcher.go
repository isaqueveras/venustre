package venustre

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"syscall"
	"time"
)

// ID defines the type of identifier
type ID string

// ToString return the identifier as a string
func (id ID) ToString() string {
	return string(id)
}

// Watch defines the type of the watcher structure
type Watch struct {
	Id    ID        `json:"id"`
	Name  string    `json:"name"`
	RunAt time.Time `json:"run_at"`

	engine Engine
	log    ILogger
}

// Watcher initializes a new watcher
func Watcher(id, name string, opts ...WatcherOption) *Watch {
	watch := &Watch{
		Id:     ID(id),
		Name:   name,
		log:    setupLogger(),
		engine: newEngine(),
		RunAt:  time.Now(),
	}

	for _, opt := range opts {
		opt(watch)
	}

	return watch
}

// Wait method responsible for keeping routines running
func (watch *Watch) Wait() {
	if err := watch.engine.Wait(); err != nil {
		watch.log.Errorf("%s", err.Error())
		return
	}
}

// Wait responsible for keeping routines running
func Wait() {
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)

	for range wait {
		return
	}
}

// Go create a new routine in the watcher
func (watch *Watch) Go(opts ...Option) {
	watch.engine.Go(func() error {
		ctx := &Context{
			indicator: make([]*indicator, 0),
			metadata:  make(Metadata),
			log:       watch.log,
			Interval:  time.Minute,
			RunAt:     time.Now(),
			Watcher:   *watch,
			context:   context.Background(),
		}

		for _, opt := range opts {
			opt(ctx)
		}

		if err := ctx.validate(); err != nil {
			return err
		}

		info := runtime.FuncForPC(reflect.ValueOf(ctx.script).Pointer())
		file, line := info.FileLine(info.Entry())
		ctx.Path = fmt.Sprintf("%s:%v", file, line)

		if err := watch.engine.Init(ctx); err != nil {
			return err
		}

		defer func() {
			if r := recover(); r != nil {
				ctx.log.Panicf(fmt.Sprintf("%v", r))
			}
		}()

		if ctx.notUseLoop {
			return ctx.execute()
		}

		ticker := time.NewTicker(ctx.Interval)
		defer ticker.Stop()

		for range ticker.C {
			if err := ctx.execute(); err != nil {
				ctx.log.Errorf(err.Error())
				continue
			}
		}

		return nil
	})
}

func (ctx *Context) execute() error {
	now := time.Now()
	ctx.sleep(now)
	defer func() {
		if r := recover(); r != nil {
			ctx.log.Panicf(fmt.Sprintf("%v", r))
		}
	}()

	if err := ctx.Watcher.engine.Before(ctx); err != nil {
		return err
	}

	if err := ctx.script(ctx); err != nil {
		return err
	}

	ctx.latency = time.Since(now)
	if err := ctx.Watcher.engine.After(ctx); err != nil {
		return err
	}

	ctx.metrics(&ctx.Watcher, now)

	return nil
}
