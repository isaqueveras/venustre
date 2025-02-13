# How to Use

Using the Venüstre library is very simple, you just need to install the library and use the following example to execute your first routine using the default implementation.

```go
package main

import "github.com/isaqueveras/venustre"

func main() {
  w := venustre.Watcher("10a32703b39d", "v1/venustre/localtime")
  defer w.Wait()

  w.Go(
    venustre.WithID("b6158d1a2ae4"),
    venustre.WithName("Routine that say local time")
    venustre.WithInterval(time.Second)
    venustre.WithScript(SayTheTimeLocal)
  )
}
```

Let's create a function that will be responsible for just speaking the local time.

```go
func SayTheTimeLocal(ctx *venustre.Context) error {
  ctx.Info("The time now is %s", time.Now().String())
  return nil
}
```

---

# Indicator, Histogram and Metadata

Indicators, histograms and metadata can be used within the routine function, so you can count values ​​and use this data to save in a database to be used in reports.

The `Inc()` and `Add()` methods of indicators and histograms have the same functionality. The `Inc()` method simply increments the value by `1` and the `Add()` method adds or subtracts a value.

The only difference between an indicator and a histogram is in the composition of the data structure; the histogram records the date the data was captured each time the `Inc()` and `Add()` methods are used.

## Indicators

Creating an indicator is very simple, let's continue using the example function above:

```go
func SayTheTimeLocal(ctx *venustre.Context) error {
  total := ctx.NewIndicator("total")

  total.Add(10)
  total.Add(-6)

  for {
    ...
    total.Inc()
  }

  return nil
}
```

## Histograms

Creating a histogram is very simple, let's continue using the example function above:

```go
func SayTheTimeLocal(ctx *venustre.Context) error {
  activated := ctx.NewHistogram("user_activated")

  for {
    ...
    activated.Add(10)
    activated.Inc()
  }

  return nil
}
```

## Metadata

Creating a metadata is very simple, let's continue using the example function above:

```go
func SayTheTimeLocal(ctx *venustre.Context) error {
  ctx.Metadata("customer", venustre.Metadata{
    "contracts": []int64{724, 234, 212},
    "address": venustre.Metadata{
      "id":     123,
      "street": "Rua da minha sogra",
    },
  })
  return nil
}
```

## Get data in `Event` method

This data can be captured in the application's `Event` method like this:

```go
func (*VenustreImp) Event(ctx *venustre.Context, event venustre.Event) {
  metric, ok := event.(venustre.EventMetric)
  if !ok {
    return
  }

  // All indicators generated in the routines
  for _, indicator := range metric.Indicators {
    _ = indicator
  }

  // All histograms generated in the routines
  for _, histogram := range metric.Histograms {
    _ = histogram
  }

  // All metadata generated in routines
  if metric.Metadata != nil && len(metric.Metadata) != 0 {
    _ = metric.Metadata
  }
}
```

Metadata is an essential tool for storing data in JSON format so that it can be captured and stored dynamically by each programmed routine.
