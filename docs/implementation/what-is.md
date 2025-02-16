# Whats is Implementation

The library has a basic implementation of the `Engine` interface so you can use it for personal projects, it is not recommended for script server projects with many routines.

Use this default implementation for study and learning purposes only.

## Own Implementation

To have greater control, we strongly recommend that you use or create an implementation of the main **`Engine`** interface so that you can control the methods that are used within the routine executions.

With this proprietary implementation, you will be able to receive events emitted in routines and save or send this data to any database or use it in Grafana.

You can implement the logs `ILogger` interface to have control over which library you want to use for logging.

See how to create your own implementation [here](./how-it-works.md).
