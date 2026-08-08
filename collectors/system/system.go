package system

import (
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/Shihab369/devops-thinking-lab/internal/core"
	"github.com/Shihab369/devops-thinking-lab/internal/models"
)

type SystemCollector struct{}

var _ core.Collector = SystemCollector{}

func (s SystemCollector) Collect() models.Result {
	return models.Result{
		Name: "system",
		Data: map[string]interface{}{
			"hostname":     readHostname(),
			"kernel":       readKernel(),
			"os":           readOSRelease(),
			"architecture": runtime.GOARCH,
			"logical_cpus": runtime.NumCPU(),
			"memory":       readMemoryStats(),
		},
	}
}

func readHostname() string {
	data, err := os.ReadFile("/proc/sys/kernel/hostname")
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}

func readKernel() string {
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}

func readOSRelease() map[string]string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return map[string]string{}
	}

	result := make(map[string]string)

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.Trim(strings.TrimSpace(parts[1]), `"`)

		result[key] = value
	}

	return result

}

func readMemoryStats() map[string]uint64 {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return map[string]uint64{}
	}

	result := make(map[string]uint64)

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		key, value, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}

		key = strings.TrimSpace(key)

		if key != "MemTotal" &&
			key != "MemAvailable" &&
			key != "MemFree" {
			continue
		}

		fields := strings.Fields(value)
		if len(fields) < 2 {
			continue
		}

		kb, err := strconv.ParseUint(fields[0], 10, 64)
		if err != nil {
			continue
		}

		switch key {
		case "MemTotal":
			result["total_bytes"] = kb * 1024
		case "MemAvailable":
			result["available_bytes"] = kb * 1024
		case "MemFree":
			result["free_bytes"] = kb * 1024
		}
	}

	return result
}
