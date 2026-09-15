# Output Architecture

## Purpose

The `internal/output` package formats collected results for human-readable terminal output.

## Human-readable Output

The formatter converts `models.Result` values into readable terminal output:

```text
[cpu]
usage_percent: 5.04
```

## Responsibility

The output package:

- Formats results for terminal display.
- Does not collect system information.
- Does not execute collectors.
- Does not manage CLI flags.

JSON encoding is currently handled directly by `cmd/devops-lab/main.go`.

## Current Flow

```text
[]models.Result
       |
       v
output.PrintResults()
       |
       v
Human-readable terminal output
```
