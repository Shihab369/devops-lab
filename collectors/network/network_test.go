package network

import (
	"net"
	"testing"
)

func TestListInterfaces(t *testing.T) {
	interfaces := listInterfaces()

	if len(interfaces) == 0 {
		t.Fatal("expected at least one network interface")
	}
}

func TestReadInterfaceInfo(t *testing.T) {
	interfaces := listInterfaces()

	if len(interfaces) == 0 {
		t.Fatal("expected at least one network interface")
	}

	info := readInterfaceInfo(interfaces[0])

	if info.name == "" {
		t.Fatal("expected interface name")
	}

	if info.mtu == 0 {
		t.Fatal("expected positive MTU")
	}

	if info.ipv4Addresses == nil {
		t.Fatal("expected IPv4 addresses slice")
	}

	if info.ipv6Addresses == nil {
		t.Fatal("expected IPv6 addresses slice")
	}
}

func TestNetworkCollector(t *testing.T) {
	collector := NetworkCollector{}

	result := collector.Collect()

	if result.Name != "network" {
		t.Fatalf("expected result name network, got %s", result.Name)
	}

	if len(result.Data) == 0 {
		t.Fatal("expected network interface data")
	}

	for name, value := range result.Data {
		if name == "" {
			t.Fatal("expected non-empty interface name")
		}

		info, ok := value.(map[string]interface{})
		if !ok {
			t.Fatalf("expected interface data to be map[string]interface{}, got %T", value)
		}

		if _, ok := info["state"]; !ok {
			t.Fatalf("expected state for interface %s", name)
		}

		if _, ok := info["mtu"]; !ok {
			t.Fatalf("expected mtu for interface %s", name)
		}

		if _, ok := info["rx_bytes"]; !ok {
			t.Fatalf("expected rx_bytes for interface %s", name)
		}

		if _, ok := info["tx_bytes"]; !ok {
			t.Fatalf("expected tx_bytes for interface %s", name)
		}
	}
}

func TestSplitAddresses(t *testing.T) {
	addrs := []net.Addr{
		&net.IPNet{
			IP:   net.ParseIP("192.168.1.10"),
			Mask: net.CIDRMask(24, 32),
		},
		&net.IPNet{
			IP:   net.ParseIP("10.0.0.5"),
			Mask: net.CIDRMask(8, 32),
		},
		&net.IPNet{
			IP:   net.ParseIP("fe80::1"),
			Mask: net.CIDRMask(64, 128),
		},
		&net.IPNet{
			IP:   net.ParseIP("2001:db8::1"),
			Mask: net.CIDRMask(64, 128),
		},
	}

	ipv4, ipv6 := splitAddresses(addrs)

	if len(ipv4) != 2 {
		t.Fatalf("expected 2 IPv4 addresses, got %d", len(ipv4))
	}

	if len(ipv6) != 2 {
		t.Fatalf("expected 2 IPv6 addresses, got %d", len(ipv6))
	}

	if ipv4[0] != "192.168.1.10/24" {
		t.Fatalf("unexpected IPv4 address: %s", ipv4[0])
	}

	if ipv6[0] != "fe80::1/64" {
		t.Fatalf("unexpected IPv6 address: %s", ipv6[0])
	}
}
