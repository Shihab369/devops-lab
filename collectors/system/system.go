package system

import (
	"os"
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
			"hostname": readHostname(),
			"kernel":   readKernel(),
			"os":       readOSRelease(),
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
