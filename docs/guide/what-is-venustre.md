# What is Venüstre?

Venüstre is a simple, dependency-free framework for [creating]() and [managing]() routines (scripts) for the [Golang]() programming language.

Its use is intended for those who need to create and manage scripts using goroutines, without any external dependencies.

If you need some of these points, then this framework is for you:
- Need to schedule a script to run at a specific time of day?
- Do you need to know/have indicators and histograms in data processing?
- Need to have control over which scripts are running in real time?
- Do you need to organize many routines (scripts) in a simple way?

<div class="tip custom-block" style="padding-top: 8px">

Just want to try it out? Skip to the [Quickstart](./getting-started.md).

</div>

## Developer Experience

Venüstre aims to provide a great developer experience when working with routines (scripts) with the ease of using just one function and defining the interval at which it should be executed.

- **Watcher:** it is the main function that initializes a group of routines, being their parent. This allows creating routines with many domains (contexts), just like the grouping of http routes works.

- **Go:** it is the main method that initializes a routine within a watcher, not only the identifier, name (for humans to understand), interval and the function that should be executed are required.

  The function to be executed must comply with the typing.
  ```go
  func (ctx *venustre.Context) error
  ```

- **Wait** this method only causes the routines to be executed and the process is not finalized, its use is mandatory.
