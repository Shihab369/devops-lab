# Continuous Integration

GitHub Actions provides automated validation for the repository.

## Purpose

CI should prevent changes from entering `main` when the project does not build or tests fail.

## Current Project Version

The project declares Go 1.26.5 in `go.mod`.

CI should use Go 1.26.x or another explicitly compatible Go 1.26+ version.

Do not configure CI with an unrelated older version such as Go 1.20.

## Basic CI Pipeline

The minimum useful pipeline is:

```text
Checkout
   |
   v
Set up Go
   |
   v
 Build
   |
   v
  Test
   |
   v
  Vet
```

Recommended commands:

```bash
go build ./...
go test ./...
go vet ./...
```

## Trigger

CI should run for:

- Pushes to `main`.
- Pull requests targeting `main`.

This ensures both direct changes and proposed changes are validated.

## When CI Fails

Do not ignore the failure.

Process:

1. Open the failed workflow run.
2. Identify the failed step.
3. Read the first relevant error.
4. Reproduce the command locally.
5. Fix the root cause.
6. Run the same command locally.
7. Run the full quality checks.
8. Push the fix.
9. Confirm CI passes.

## CI and Local Development

CI should reproduce the same checks developers run locally.

Local:

```bash
go build ./...
go test ./...
go vet ./...
```

CI:

```bash
go build ./...
go test ./...
go vet ./...
```

The goal is to catch failures before merge.
