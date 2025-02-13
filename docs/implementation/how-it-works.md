# How It Works Implementation

To create your own implementation of the main interface, you must be aware of how the interface methods work.

## Implementation of the main interface

```go
type IVenustre interface {
  Go(fn func() error)
  Wait() error

  Init(ctx *Context) error
  Event(ctx *Context, event Event)

  Before(ctx *Context) error
  After(ctx *Context) error
}
```

The `Go` and `Wait` methods are necessary so that you can have control over the asynchronous execution at startup.

> Disclaimer: these two methods (`Go` and `Wait`) are being evaluated for discontinuation. See more in this [issue](https://github.com/isaqueveras/venustre/issues/18).

The `Init` method is used as soon as the application starts, with this you can create a mechanism to record this information in a database for auditing purposes.

The `Event` method is where you can capture events recorded during the execution of routines, such as indicators, histograms, metadata and data related to the execution of the routine, such as start and end dates and execution latency.

The `Before` and `After` methods are always used before starting the execution of the routine and after its execution, this is useful in case you want to store information about the routine before and after completion.

With these methods it is possible to use integration of observability libraries such as [OpenTelemetry](https://en.wikipedia.org/wiki/Cloud_Native_Computing_Foundation#OpenTelemetry) to capture data related to execution.


## Implementação da interface de Logger

The logger interface is very simple, each method represents the type of log that will be used within the routines.

> The logger interface was prototyped to fit the methods of the [zap](https://github.com/uber-go/zap) library, but it needs to be refactored to follow a simpler definition so that implementations can be made using other libraries.

```go
type ILogger interface {
  Infof(format string, v ...interface{})
  Warnf(format string, v ...interface{})
  Errorf(format string, v ...interface{})
  Debugf(format string, v ...interface{})
  Panicf(format string, v ...interface{})
}
```
