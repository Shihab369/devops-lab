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
