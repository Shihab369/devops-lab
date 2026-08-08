package uptime

import (
	"os"
	"strconv"
	"strings"

	"github.com/Shihab369/devops-thinking-lab/internal/core"
	"github.com/Shihab369/devops-thinking-lab/internal/models"
)

type UptimeCollector struct{}

var _ core.Collector = UptimeCollector{}

func (u UptimeCollector) Collect() models.Result {
	uptime := readUptime()

	return models.Result{
		Name: "uptime",
		Data: map[string]interface{}{
			"uptime_seconds": uptime,
		},
	}
}

func readUptime() float64 {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return 0
	}

	fields := strings.Fields(string(data))

	if len(fields) < 1 {
		return 0
	}

	uptime, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return 0
	}

	return uptime
}
