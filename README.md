# DevOps Lab

A lightweight Linux system diagnostics CLI written in Go.

DevOps Lab collects low-level system information from a Linux host and presents the results in either a human-readable format or structured JSON.

The project is designed as a small, modular foundation for experimenting with system diagnostics, collectors, observability concepts, and practical DevOps tooling.

## Features

Currently implemented collectors:

- CPU usage
- Disk usage
- Load average
- Memory statistics
- Process information
- System information
- System uptime

### Output formats

Human-readable output:

````text
[cpu]
usage_percent: 5.04

[memory]
total_bytes: 16580902912
available_bytes: 9771622400
free_bytes: 2123935744

JSON output:

```json
[
  {
    "name": "cpu",
    "data": {
      "usage_percent": 5.04
    }
  },
  {
    "name": "uptime",
    "data": {
      "uptime_seconds": 203839.03
    }
  }
]
````

## Installation

### Requirements

- Linux
- Go 1.26+

### Clone the repository

```bash
git clone https://github.com/Shihab369/devops-lab.git
cd devops-lab
```

### Run

Run the CLI with:

```bash
go run ./cmd/devops-lab
```

For JSON output:

```bash
go run ./cmd/devops-lab --json
```

To see available options:

```bash
go run ./cmd/devops-lab --help
```

## Architecture

The project is organized around independent collectors and a shared execution pipeline.

```text
cmd/
└── devops-lab/
    └── main.go

collectors/
├── cpu/
├── disk/
├── load/
├── memory/
├── process/
├── system/
└── uptime/

internal/
├── core/
├── models/
└── output/
```

### Core

The `internal/core` package defines the collector contract and coordinates collector execution.

This keeps collection logic independent from the CLI and output formatting.

### Models

The `internal/models` package contains the common result model shared across collectors.

### Output

The `internal/output` package is responsible for formatting collected results into the supported output formats.

## Development

Format the code:

```bash
gofmt -w .
```

Run the test suite:

```bash
go test ./...
```

Run static analysis:

```bash
go vet ./...
```

Check for whitespace errors:

```bash
git diff --check
```

Before opening a pull request, make sure all checks pass:

```bash
gofmt -w .
go test ./...
go vet ./...
git diff --check
```

## Contributing

Contributions are welcome.

Before contributing, please open an issue to discuss significant changes or new features.

For smaller fixes, documentation improvements, tests, or collector additions, you can open a pull request directly.

### Development workflow

1. Fork the repository.
2. Clone your fork.
3. Create a feature branch from `main`.
4. Make your changes.
5. Run the development checks.
6. Commit your changes.
7. Push your branch.
8. Open a pull request against `main`.

Example:

```bash
git checkout -b feat/add-network-collector
```

After making changes, run:

```bash
gofmt -w .
go test ./...
go vet ./...
git diff --check
```

Use clear, focused commits that describe the change:

```bash
git add .
git commit -m "feat: add network collector"
```

Push your branch:

```bash
git push origin feat/add-network-collector
```

Then open a pull request against the `main` branch.

### Pull request guidelines

- Keep pull requests focused on one change.
- Add or update tests when applicable.
- Update documentation when behavior changes.
- Avoid unrelated changes.
- Make sure all development checks pass before requesting review.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Roadmap

Planned improvements include:

- Additional Linux system collectors
- Network diagnostics
- Disk and filesystem health checks
- Configurable output formats
- More diagnostic capabilities
- Improved CLI options
- Cross-platform support where practical
- Better documentation and contributor tooling
