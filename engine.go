package venustre

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
)

type engine struct{}

func newEngine() Engine { return &engine{} }

func (engine) Go(fn func() error) { fn() }
func (engine) Wait() error        { return nil }

// Init implements a business rule when initializing a routine
func (engine) Init(ctx *Context) error {
	ctx.Info(fmt.Sprintf("script '%s' (rid: %s) initialized", ctx.Name, ctx.RoutineID))
	return nil
}

// Before implements a business rule before initializing script execution
func (engine) Before(ctx *Context) error {
	ctx.Id = ID(strconv.FormatInt(rand.Int63(), 10))
	ctx.Info(fmt.Sprintf("script '%s' (rid: %s, id: %s) initialized", ctx.Name, ctx.RoutineID, ctx.Id))
	return nil
}

// After implements a business rule after initializing script execution
func (engine) After(ctx *Context) error {
	ctx.Info(fmt.Sprintf("script '%s' (rid: %s, id: %s) finished", ctx.Name, ctx.RoutineID, ctx.Id))
	return nil
}

// Event implements a business rule for event handling
func (engine) Event(ctx *Context, event Event) {
	if metric, ok := event.(EventMetric); ok {
		value, _ := json.Marshal(metric)
		ctx.Debug(string(value))
	}
}
