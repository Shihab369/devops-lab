package memory

import (
	"os"
	"strconv"
	"strings"

	"github.com/Shihab369/devops-lab/internal/core"
	"github.com/Shihab369/devops-lab/internal/models"
)

type MemoryCollector struct{}

var _ core.Collector = MemoryCollector{}

func (m MemoryCollector) Collect() models.Result {
	stats := readMemoryStats()

	return models.Result{
		Name: "memory",
		Data: map[string]interface{}{
			"total_bytes":     stats["total_bytes"],
			"available_bytes": stats["available_bytes"],
			"free_bytes":      stats["free_bytes"],
		},
	}
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
