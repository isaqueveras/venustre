<p align="center">
  <img src="https://github.com/user-attachments/assets/158c4190-85a9-4154-b29a-4c3c2f00bb41" width="200">
</p>

<p align="center">
<strong>A documentation for the <a href="https://github.com/isaqueveras/outis">Outis</a> library</strong> <br/>
The outis library helps you create and manage routines,<br/>schedule execution time, and control competition between other processes.
</p>

<p align="center">
  <a href="https://github.com/sponsors/isaqueveras"><img src="https://img.shields.io/github/sponsors/isaqueveras" alt="Sponsors"/></a>
  <img src="https://img.shields.io/github/go-mod/go-version/isaqueveras/outis" alt="GitHub go.mod Go version"/>
  <img alt="GitHub commit activity" src="https://img.shields.io/github/commit-activity/y/isaqueveras/outis">
</p>

---

## How to use.

```go
package main

import (
  "time"

  "github.com/isaqueveras/outis"
)

func main() {
  // Initialize Outis to be able to add routines
  watch := outis.Watcher("35d3965c8783", "v1/example",
    outis.WithLogger(nil),         // Option to implement logs interface
    outis.WithOutisInterface(nil), // Option to implement outis interface
  )

  watch.Go(
    // Routine identifier to perform concurrency control
    outis.WithID("8cf7e174fb4f"),

    outis.WithName("Here is the name of my routine"),
    outis.WithDesc("Here is the description of my routine"),

    // It will run every 10 second
    outis.WithInterval(time.Second * 10),

    // It will run from 12pm to 4pm.
    // by default, there are no time restrictions.
    outis.WithHours(12, 16),

    // Time when routine information will be updated
    outis.WithLoadInterval(time.Second * 30),

    // Here the script function that will be executed will be passed
    outis.WithScript(func(ctx *outis.Context) {
      ctx.Info("this is an information message")
      ctx.Error("error: %v", errors.New("this is an error message"))

      ctx.Metric("client_ids", []int64{234234})
      ctx.Metric("notification", outis.Metric{
        "client_id": 234234,
        "message":   "Hi, we are notifying you.",
        "fcm":       "3p2okrmionfiun2uni3nfin2i3f",
      })

      ctx.Debug("Hello")
    }),
  )

  // Method that maintains routine in the process
  watch.Wait()
}
```

## Reporting issues

The issue tracker for this project [is located here](https://github.com/isaqueveras/outis/issues).

Please report any issues with a sufficient description of the bug or feature request. Ideally, bug reports should be accompanied by a minimal reproduction of the issue.
Bug reports must specify the [version](https://github.com/isaqueveras/outis/releases).

## Contributing
This project is open-source and accepts contributions. See the [contribution guide](https://github.com/isaqueveras/outis/blob/main/CONTRIBUTING.md) for more information.


