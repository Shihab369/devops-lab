package network

import "testing"

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
