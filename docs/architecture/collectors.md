# Collector Architecture

## Purpose

The `collectors` directory contains the Linux system diagnostics implemented by DevOps Lab.

Each collector is responsible for one diagnostic area and returns its result through the shared `core.Collector` interface.

Collectors do not control application flow and do not format terminal output. Their responsibility is to collect system data and return a `models.Result`.

## Collector Contract

All collectors implement the following interface:

```go
type Collector interface {
    Collect() models.Result
}
```

## Current Collectors

The project currently contains the following collectors:

```text
collectors/
├── cpu/
├── disk/
├── load/
├── memory/
├── network/
├── process/
├── system/
└── uptime/
```

## CPU Collector

Location:

```text
collectors/cpu/
```

The CPU collector gathers CPU usage information from the Linux system. It returns the CPU usage percentage as part of the shared result model.

## Disk Collector

Location:

```text
collectors/disk/
```

The disk collector gathers filesystem and disk usage information available through the system.

## Load Collector

Location:

```text
collectors/load/
```

The load collector gathers Linux load-average information.

## Memory Collector

Location:

```text
collectors/memory/
```

The memory collector gathers memory statistics such as:

- Total memory.
- Available memory.
- Free memory.

## Network Collector

Location:

```text
collectors/network/
```

The network collector gathers information about network interfaces.

### Current diagnostics include

- Interface name.
- Interface state.
- MAC address.
- MTU.
- IPv4 addresses.
- IPv6 addresses.
- Received bytes.
- Received packets.
- Transmitted bytes.
- Transmitted packets.

The collector uses Go's networking APIs together with Linux sysfs data under:

```text
/sys/class/net/
```

For interfaces without IPv4 or IPv6 addresses, the collector returns empty address lists instead of nil slices. This keeps the result structure predictable and makes JSON output consistent.

## Process Collector

Location:

```text
collectors/process/
```

The process collector gathers information about running processes on the Linux host.

## System Collector

Location:

```text
collectors/system/
```

The system collector gathers general host information.

## Uptime Collector

Location:

```text
collectors/uptime/
```

The uptime collector reports how long the Linux system has been running.

## Collector Execution

Collectors are registered by `cmd/devops-lab/main.go` and passed to the runner:

```text
main()
  |
  v
core.NewRunner(
    CPU,
    Disk,
    Load,
    Memory,
    Network,
    Process,
    System,
    Uptime,
)
  |
  v
runner.Run()
  |
  v
collector.Collect()
  |
  v
models.Result
```

The runner executes collectors sequentially in the order in which they are registered.

## Result Ownership

Collectors create and return `models.Result` values.

They do not:

- Print directly to stdout.
- Decide the output format.
- Control other collectors.
- Manage CLI flags.
- Execute the application lifecycle.

The collector layer is responsible only for collecting diagnostic data and returning the result.

## Error Handling

The current collector contract returns only `models.Result`:

```go
type Collector interface {
    Collect() models.Result
}
```

There is no separate `error` return value in the current collector interface.

Therefore, collector-specific failures are represented through the returned result model rather than being propagated through a separate error value.

This is a current design constraint of the collector architecture.

## Testing

Each collector package contains tests alongside its implementation:

```text
collectors/<name>/<name>_test.go
```

Collector tests verify the behavior of the individual collector without depending on the CLI entrypoint.

The complete test suite can be executed with:

```text
go test ./...
```

## Design Boundary

Each collector should have one clear responsibility:

```text
One diagnostic area
        |
        v
One collector
        |
        v
One models.Result
```

The collector does not decide what happens to the result after collection.

Execution is handled by `internal/core`, while human-readable presentation is handled by `internal/output`. JSON encoding is currently handled by `cmd/devops-lab/main.go`.

## Current Architecture Summary

The current collector architecture can be summarized as:

```text
cmd/devops-lab/main.go
          |
          v
    core.NewRunner(...)
          |
          v
      runner.Run()
          |
          +----> CPU Collector ------> models.Result
          |
          +----> Disk Collector -----> models.Result
          |
          +----> Load Collector -----> models.Result
          |
          +----> Memory Collector ----> models.Result
          |
          +----> Network Collector ---> models.Result
          |
          +----> Process Collector ---> models.Result
          |
          +----> System Collector ----> models.Result
          |
          +----> Uptime Collector ----> models.Result
          |
          v
     []models.Result
```

The collectors are independent diagnostic components that share a common contract and result model.

The current architecture favors simple composition: new collectors can implement the core.Collector interface and be registered with the runner.
