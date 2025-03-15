package venustre

import (
	"log"
	"sync"
)

type Worker interface {
	Kind() string
	Work(*Context)
}

type Enginex struct {
	workers map[string]GroupInfo
	wg      sync.WaitGroup
}

type GroupInfo struct {
	id    string
	group func(*Group)
}

type WorkerInfo struct {
	id     string
	name   string
	worker Worker
}

type Group struct {
	workers map[string]WorkerInfo
}

// NewEnginex initializes and starts a set of WorkerBuilder instances concurrently.
// It accepts a variadic number of WorkerBuilder instances, adds them to a wait group,
// and starts each one in a separate goroutine. The function waits for all goroutines
// to complete before returning.
func NewEngine(opts ...WatcherOption) *Enginex {
	// var group = &Group{}

	// var engx []func(*Group)

	// var wg sync.WaitGroup
	// for _, fn := range engx {
	// 	fn := fn
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		fn(group)
	// 		// e.Work(&Context{})
	// 	}()
	// }
	// wg.Wait()

	return &Enginex{}
}

func (e *Enginex) Start() {
	e.wg.Add(1)

	for id, group := range e.workers {
		log.Printf("id: %s\t", id)
		group.group(&Group{})
	}
}

func (e *Enginex) Group(id string, group func(g *Enginex)) *Enginex {
	if id == "" {
		panic("id is null")
	}

	if group == nil {
		panic("group function is null")
	}

	if e.workers == nil {
		e.workers = make(map[string]GroupInfo)
	}

	e.workers[id] = GroupInfo{
		id: id,
		// group: group,
	}

	return e
}

func (g *Enginex) Work(id, name string, worker Worker) *Group {
	// g.workers[id] = WorkerInfo{
	// 	id:     id,
	// 	name:   name,
	// 	worker: worker,
	// }

	// if worker == nil {
	// panic("Worker null")
	// }

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()

	// 	ticker := time.NewTicker(time.Second)
	// 	defer ticker.Stop()

	// 	for range ticker.C {
	// 		// wg.Add(1)
	// 		// go func() {
	// 		// defer wg.Done()
	// 		// worker.Work(&Context{})
	// 		log.Printf("(%s) name: %s -> kind: %s:%d", id, name, worker.Kind(), 12)
	// 		// }()
	// 	}
	// }()
	// wg.Wait()

	return nil
}
