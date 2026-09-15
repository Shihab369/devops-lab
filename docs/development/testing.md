# Testing

## Purpose

Tests verify collector behavior, parsing logic, and core execution.

## Test Location

Tests should normally live beside the implementation:

```text
collectors/network/
├── network.go
└── network_test.go
```

The core package also contains an integration test:

```internal/core/
├── runner_integration_test.go
└── runner_test.go
```

## Unit Tests

Unit tests should focus on one behavior at a time.

Examples:

- Parsing system data.
- Splitting IPv4 and IPv6 addresses.
- Handling empty data.
- Validating collected fields.
- Verifying collector results.

## Integration Tests

Integration tests verify that real collectors can execute through the core runner.

## Run All Tests

```bash
go test ./...
```

## Run One Package

```bash
go test ./collectors/network
```

## Run Verbose Tests

```bash
go test -v ./...
```

## Test Failure Process

When a test fails:

1. Find the first failing package.
2. Find the first failing test.
3. Read the exact assertion.
4. Reproduce locally.
5. Inspect the implementation and test.
6. Fix the root cause.
7. Run the failing package again.
8. Run the full test suite.

Do not assume that a later `PASS` means the whole test suite passed. `go test ./...` fails if any package fails.

## Definition of Done

```text
go test ./...
```

must complete successfully before opening a PR.
