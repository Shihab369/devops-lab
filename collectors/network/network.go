package network

import (
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Shihab369/devops-lab/internal/core"
	"github.com/Shihab369/devops-lab/internal/models"
)

type NetworkCollector struct{}

var _ core.Collector = NetworkCollector{}

type interfaceInfo struct {
	name          string
	state         string
	macAddress    string
	mtu           uint64
	ipv4Addresses []string
	ipv6Addresses []string
	rxBytes       uint64
	txBytes       uint64
	rxPackets     uint64
	txPackets     uint64
}

func (n NetworkCollector) Collect() models.Result {
	interfaces := listInterfaces()

	data := make(map[string]interface{})

	for _, name := range interfaces {
		info := readInterfaceInfo(name)

		data[name] = map[string]interface{}{
			"state":          info.state,
			"mac_address":    info.macAddress,
			"mtu":            info.mtu,
			"ipv4_addresses": info.ipv4Addresses,
			"ipv6_addresses": info.ipv6Addresses,
			"rx_bytes":       info.rxBytes,
			"tx_bytes":       info.txBytes,
			"rx_packets":     info.rxPackets,
			"tx_packets":     info.txPackets,
		}
	}

	return models.Result{
		Name: "network",
		Data: data,
	}
}

func listInterfaces() []string {
	entries, err := os.ReadDir("/sys/class/net")
	if err != nil {
		return []string{}
	}

	interfaces := make([]string, 0, len(entries))

	for _, entry := range entries {
		interfaces = append(interfaces, entry.Name())
	}

	return interfaces
}

func splitAddresses(addrs []net.Addr) ([]string, []string) {
	var ipv4Addresses []string
	var ipv6Addresses []string

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}

		if ipNet.IP.To4() != nil {
			ipv4Addresses = append(ipv4Addresses, ipNet.String())
		} else {
			ipv6Addresses = append(ipv6Addresses, ipNet.String())
		}
	}

	return ipv4Addresses, ipv6Addresses
}

func readInterfaceInfo(name string) interfaceInfo {
	basePath := filepath.Join("/sys/class/net", name)

	info := interfaceInfo{
		name:          name,
		state:         readString(filepath.Join(basePath, "operstate")),
		macAddress:    readString(filepath.Join(basePath, "address")),
		mtu:           readUint64(filepath.Join(basePath, "mtu")),
		rxBytes:       readUint64(filepath.Join(basePath, "statistics/rx_bytes")),
		txBytes:       readUint64(filepath.Join(basePath, "statistics/tx_bytes")),
		rxPackets:     readUint64(filepath.Join(basePath, "statistics/rx_packets")),
		txPackets:     readUint64(filepath.Join(basePath, "statistics/tx_packets")),
		ipv4Addresses: []string{},
		ipv6Addresses: []string{},
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return info
	}

	for _, iface := range interfaces {
		if iface.Name != name {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return info
		}

		info.ipv4Addresses, info.ipv6Addresses = splitAddresses(addrs)
		break
	}

	return info
}

func readString(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}

func readUint64(path string) uint64 {
	value := readString(path)

	number, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0
	}

	return number
}
