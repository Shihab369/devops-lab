package memory

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Shihab369/devops-lab/internal/core"
	"github.com/Shihab369/devops-lab/internal/models"
)

type MemoryCollector struct{}

var _ core.Collector = MemoryCollector{}

type memoryStats struct {
	totalBytes     uint64
	availableBytes uint64
	freeBytes      uint64
}

func (m MemoryCollector) Collect() models.Result {
	stats, err := readMemoryStats()

	data := map[string]interface{}{
		"total_bytes":     stats.totalBytes,
		"available_bytes": stats.availableBytes,
		"free_bytes":      stats.freeBytes,
	}

	if err != nil {
		data["collection_error"] = err.Error()
	}

	return models.Result{
		Name: "memory",
		Data: data,
	}
}

func readMemoryStats() (memoryStats, error) {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return memoryStats{}, fmt.Errorf("failed to read /proc/meminfo: %w", err)
	}

	return parseMemInfo(string(data))
}

func parseMemInfo(input string) (memoryStats, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return memoryStats{}, fmt.Errorf("meminfo input is empty")
	}

	required := map[string]bool{
		"MemTotal":     false,
		"MemAvailable": false,
		"MemFree":      false,
	}

	stats := memoryStats{}

	for _, line := range strings.Split(trimmed, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		key, value, ok := strings.Cut(line, ":")
		if !ok {
			if requiredKey, isRequired := requiredKeyPrefix(line); isRequired {
				return memoryStats{}, fmt.Errorf("incomplete required field %s", requiredKey)
			}

			continue
		}

		key = strings.TrimSpace(key)

		if _, isRequired := required[key]; !isRequired {
			continue
		}

		fields := strings.Fields(value)
		if len(fields) < 2 {
			return memoryStats{}, fmt.Errorf("incomplete required field %s", key)
		}

		kb, err := strconv.ParseUint(fields[0], 10, 64)
		if err != nil {
			return memoryStats{}, fmt.Errorf("invalid numeric value for %s: %w", key, err)
		}

		if fields[1] != "kB" {
			return memoryStats{}, fmt.Errorf("unexpected unit for %s: %s", key, fields[1])
		}

		sizeBytes := kb * 1024

		switch key {
		case "MemTotal":
			stats.totalBytes = sizeBytes
		case "MemAvailable":
			stats.availableBytes = sizeBytes
		case "MemFree":
			stats.freeBytes = sizeBytes
		}

		required[key] = true
	}

	missing := missingRequiredFields(required)
	if len(missing) > 0 {
		return memoryStats{}, fmt.Errorf("missing required fields: %s", strings.Join(missing, ", "))
	}

	return stats, nil
}

func requiredKeyPrefix(line string) (string, bool) {
	for _, key := range []string{"MemTotal", "MemAvailable", "MemFree"} {
		if strings.HasPrefix(line, key) {
			return key, true
		}
	}

	return "", false
}

func missingRequiredFields(required map[string]bool) []string {
	missing := make([]string, 0)

	for key, present := range required {
		if !present {
			missing = append(missing, key)
		}
	}

	sort.Strings(missing)

	return missing
}
