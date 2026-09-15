# Code Quality

Run these checks before opening a pull request.

## Formatting

```bash
gofmt -w .
```

Go source files should be formatted with `gofmt`.

## Tests

```bash
go test ./...
```

All packages must pass.

## Static Analysis

```bash
go vet ./...
```

This checks for suspicious or incorrect Go constructs.

## Diff Validation

```bash
git diff --check
```

This detects whitespace errors in the diff.

## Complete Local Check

```bash
gofmt -w .
go test ./...
go vet ./...
git diff --check
```

All commands should succeed.

## Before Commit

Check the working tree:

```bash
git status
git diff
```

Only intended changes should be committed.
