# Network Collector

## Purpose

The network collector gathers basic Linux network interface diagnostics.

## Collected Information

For each interface, the collector can report:

- Interface name
- Interface state
- IPv4 addresses
- IPv6 addresses
- MAC address
- MTU
- RX bytes
- TX bytes
- RX packets
- TX packets

## Data Sources

Go's `net` package provides:

```go
net.Interfaces()
```

and interface addresses through:

```go
Interface.Addrs()
```

Linux interface metadata and statistics are read from:

```text
/sys/class/net/<interface>/
```

Relevant files include:

```text
operstate
address
mtu
statistics/rx_bytes
statistics/tx_bytes
statistics/rx_packets
statistics/tx_packets
```

## Address Handling

Addresses are separated into IPv4 and IPv6.

The collector should safely handle interfaces that have no address of a particular type.

## Scope

The initial implementation intentionally remains small.

It does not attempt to provide:

- Routing diagnostics.
- DNS diagnostics.
- Connection inspection.
- Bandwidth measurement.
- Packet capture.
- Advanced link diagnostics.

Those can be separate future features.

## Testing

Network tests cover:

- Interface discovery.
- Interface information.
- IPv4/IPv6 address separation.
- Collector execution.

Run:

```bash
go test ./collectors/network
```

Run the complete suite:

```bash
go test ./...
```
