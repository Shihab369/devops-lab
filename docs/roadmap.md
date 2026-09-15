# Roadmap

The roadmap is a direction, not a commitment to implement everything immediately.

## Current State

Implemented:

- CPU collector
- Disk collector
- Load collector
- Memory collector
- Network diagnostics collector
- Process collector
- System collector
- Uptime collector
- Human-readable output
- JSON output
- Unit tests
- Integration tests
- GitHub issue and PR workflow
- GitHub Actions CI

## Near-term Priorities

### 1. Documentation

Keep architecture, development, testing, and CI documentation aligned with the implementation.

### 2. Reliability

Improve error handling and test coverage where real system differences expose weaknesses.

### 3. Engineering Investigation

Use real Linux and DevOps problems to identify useful diagnostics and guide future development.

## Possible Future Features

Potential areas include:

- Filesystem health checks
- Additional network statistics
- More process diagnostics
- Resource thresholds
- Improved CLI options
- Configurable output
- More structured diagnostics
- Release versioning
- Automated binary builds

## Feature Selection Rule

Before starting a new feature, ask:

1. What real problem does it solve?
2. Is the scope small enough?
3. Can the existing architecture support it?
4. Can it be tested locally?
5. Does it add useful DevOps/Linux learning value?
6. What should remain out of scope?

## Avoid Premature Complexity

Do not introduce large systems before the project needs them.

Avoid prematurely adding:

- Web UI
- Database storage
- Kubernetes integration
- Distributed architecture
- Authentication
- Large frameworks
- Complex configuration systems

The immediate goal is to build a reliable Linux diagnostics CLI.
