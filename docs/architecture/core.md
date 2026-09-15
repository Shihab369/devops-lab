# Core Architecture

## Purpose

The `internal/core` package defines the common collector contract and coordinates collector execution.

It provides the execution layer between the CLI entrypoint and the individual collectors.

## Components

```text
internal/core/
├── collector.go
├── engine.go
├── runner.go
├── runner_test.go
└── runner_integration_test.go
```

## Collector Contract

```go
type Collector interface {
    Collect() models.Result
}
```

This interface allows the runner to execute different collector implementations without depending on their internal implementation details.

## Runner

The Runner is responsible for storing registered collectors and executing them.

Its main responsibilities are:

- Accept collector implementations.
- Store the collectors in registration order.
- Execute each collector.
- Collect the returned models.Result values.
- Return the aggregated results as []models.Result.

The runner is created through:

```go
core.NewRunner(...)
```

Collectors are passed to the runner by `cmd/devops-lab/main.go`

## Runner Execution

The current execution flow is:

```cmd/devops-lab/main.go
          |
          v
    core.NewRunner(...)
          |
          v
      runner.Run()
          |
          v
    for each collector
          |
          v
    collector.Collect()
          |
          v
      models.Result
          |
          v
     []models.Result
```

The runner executes collectors sequentially in the order in which they were registered. The runner does not contain collector-specific diagnostic logic.

## Collector Independence

The core layer depends on the `core.Collector` interface rather than concrete collector implementations.

As a result, the runner does not need to know how an individual collector gathers its data.

For example, the runner does not contain logic such as:

```go
if collector == network {
    // special network behavior
}
```

Network-specific behavior remains inside the network collector.

This keeps collector-specific implementation details outside the execution layer.

## Engine

The `Engine` is another component in the `internal/core`package.

It currently provides a result storage abstraction:

```go
type Engine struct {
    results []models.Result
}
```

It can:

- Create an empty result collection.
- Add a models.Result.
- Return the stored results.
  The current CLI execution path does not use Engine.

The active execution path uses `Runner` directly and returns `[]models.Result`.

Therefore, `Engine` is currently present in the core package but is not part of the active CLI runtime flow.

## Core Boundaries

The current core layer is responsible for:

```Collector Contract
       |
       v
Collector Execution
       |
       v
Result Aggregation
```

The core layer does not:

- Implement diagnostic collection logic.
- Format human-readable output.
- Encode JSON output.
- Manage individual collector internals.

## Testing

The core package contains unit tests and integration tests:

```internal/core/
├── runner_test.go
└── runner_integration_test.go
```

Unit tests verify runner behavior in isolation.

Integration tests verify that real collector implementations can execute through the core runner.

The complete test suite can be executed with:

```bash
go test ./...
```

## Current Core Architecture

The current core architecture can be summarized as:

```core.Collector
                       |
                       v
               +---------------+
               |    Runner     |
               +---------------+
                       |
                       v
             collector.Collect()
                       |
                       v
                models.Result
                       |
                       v
                 []models.Result
```

The `Runner` is the active execution component of the current CLI.

The `Engine` exists as a result storage abstraction but is not currently connected to the CLI execution path.
