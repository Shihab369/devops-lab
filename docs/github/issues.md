# GitHub Issues

Issues define planned work before implementation.

## When to Create an Issue

Create an issue for:

- New features.
- Significant improvements.
- Bugs that need tracking.
- Architectural changes.

Small obvious fixes may not require a separate issue.

## Issue Structure

Use:

```text
Objective

Scope

Requirements

Acceptance Criteria

Out of Scope
```

## Acceptance Criteria

Acceptance criteria should be testable.

Example:

```text
- Collector package exists.
- Collector is registered with the runner.
- Required data is collected.
- Unit tests are added.
- go test ./... passes.
- go vet ./... passes.
```

## Issue Lifecycle

```text
Open
  |
  v
Implementation
  |
  v
Pull Request
  |
  v
Merge
  |
  v
Closed
```

The network diagnostics work followed this process through Issue #1 and PR #3.
