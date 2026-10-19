# DevOps Lab

A lightweight Linux system diagnostics CLI written in Go.

DevOps Lab collects low-level system information from a Linux host and presents the results in either a human-readable format or structured JSON. It also provides threshold-based health checks for automation and monitoring.

The project is designed as a small, modular foundation for experimenting with system diagnostics, collectors, observability concepts, and practical DevOps tooling.

## Features

Currently implemented collectors:

- CPU usage
- Disk usage
- Load average
- Memory statistics
- Network diagnostics
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

[network]
  eth0:
    state: up
    mac_address: 00:11:22:33:44:55
    mtu: 1500
    ipv4_addresses: [192.168.1.10/24]
    ipv6_addresses: [fe80::1/64]
    rx_bytes: 123456
    tx_bytes: 654321
````

The network collector reports interface state, MAC address, MTU, IPv4 and IPv6 addresses, and RX/TX byte and packet counters when available.

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
```

### Health check mode

Use `--check` to evaluate core system metrics against configurable thresholds. The check evaluates CPU usage, memory usage, disk usage, and 1-minute load average.

Default thresholds are:

| Flag | Default | Metric |
| --- | ---: | --- |
| `--cpu-max` | `80` | Maximum CPU usage percentage |
| `--mem-max` | `85` | Maximum memory usage percentage |
| `--disk-max` | `90` | Maximum disk usage percentage |
| `--load-max` | `runtime.NumCPU()` | Maximum 1-minute load average |

The command exits with status `0` when all checks pass and status `1` when any check fails. Missing or invalid metric data is treated as a failed check.

Human-readable health check:

```bash
go run ./cmd/devops-lab --check
```

Custom thresholds:

```bash
go run ./cmd/devops-lab --check --cpu-max 80 --mem-max 85 --disk-max 90 --load-max 4
```

JSON output with a machine-readable `check` result:

```bash
go run ./cmd/devops-lab --check --json
```

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

The project is organized around independent collectors, a shared execution pipeline, and a health-evaluation layer.

```text
cmd/
└── devops-lab/
    └── main.go

collectors/
├── cpu/
├── disk/
├── load/
├── memory/
├── network/
├── process/
├── system/
└── uptime/

internal/
├── check/
├── core/
├── models/
└── output/
```

### Core

The `internal/core` package defines the collector contract and coordinates collector execution.

This keeps collection logic independent from the CLI and output formatting.

### Models

The `internal/models` package contains the common result model shared across collectors.

### Check

The `internal/check` package evaluates CPU, memory, disk, and load results against configured thresholds. It produces a normalized health summary containing overall status, passed and failed counts, and per-metric details. The CLI uses this summary for human-readable output, JSON output, and process exit signaling.

### Output

The `internal/output` package is responsible for formatting collected results into human-readable output. JSON encoding for the CLI is handled in `cmd/devops-lab/main.go`.

## Development

Format all Go files recursively:

```bash
find . -name '*.go' -not -path './vendor/*' -print0 | xargs -0 gofmt -w
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
find . -name '*.go' -not -path './vendor/*' -print0 | xargs -0 gofmt -w
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
find . -name '*.go' -not -path './vendor/*' -print0 | xargs -0 gofmt -w
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

- Filesystem health checks
- Additional network statistics
- More process diagnostics
- Configurable output formats
- More diagnostic capabilities
- Cross-platform support where practical
- Better documentation and contributor tooling
- Release versioning
- Automated binary builds
