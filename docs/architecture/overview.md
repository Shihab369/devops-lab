# Architecture Overview

## Purpose

DevOps Lab collects low-level Linux system information through independent collectors and presents the results as human-readable text or JSON.

## Runtime Flow

```text
cmd/devops-lab
    |
    v
Runner
    |
    +--> CPU Collector
    +--> Disk Collector
    +--> Load Collector
    +--> Memory Collector
    +--> Network Collector
    +--> Process Collector
    +--> System Collector
    +--> Uptime Collector
    |
    v
[]models.Result
    |
    +--------------------------+
    |                          |
    v                          v
Human-readable               JSON
    |                          |
    v                          v
output.PrintResults()       printJSON()
```

## Main Components

### cmd/devops-lab

Application entrypoint.

It:

- Registers concrete collectors with the runner.
- Executes the runner.
- Selects the output mode.
- Handles JSON encoding directly through `printJSON()`.

### collectors

Contains Linux system collection logic. Each diagnostic area has its own package.

### internal/core

Defines the collector contract and coordinates collector execution through the runner.

### internal/models

Contains the shared result model returned by collectors.

### internal/output

Formats results for human-readable terminal output.

JSON encoding is currently handled directly by `cmd/devops-lab/main.go`.

## Design Principle

Collectors collect data. Core executes collectors. Models represent results. Output formats results. The CLI wires the application together.

## Current Collectors

- CPU
- Disk
- Load
- Memory
- Network
- Process
- System
- Uptime

## Current Execution Model

The current CLI execution path is:

```text
main()
  |
  v
core.NewRunner(...)
  |
  v
runner.Run()
  |
  v
collector.Collect()
  |
  v
[]models.Result
  |
  +--> output.PrintResults()
  |
  +--> printJSON()
```

The `internal/core.Engine` component exists as a result storage abstraction but is not currently part of this CLI execution path.
