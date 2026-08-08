package load

import (
	"os"
	"strconv"
	"strings"

	"github.com/Shihab369/devops-lab/internal/core"
	"github.com/Shihab369/devops-lab/internal/models"
)

type LoadCollector struct{}

var _ core.Collector = LoadCollector{}

type loadStats struct {
	oneMinute      float64
	fiveMinutes    float64
	fifteenMinutes float64
}

func readLoadAverage() loadStats {
	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return loadStats{}
	}

	fields := strings.Fields(string(data))

	if len(fields) < 3 {
		return loadStats{}
	}

	oneMinute, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return loadStats{}
	}

	fiveMinutes, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return loadStats{}
	}

	fifteenMinutes, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return loadStats{}
	}

	return loadStats{
		oneMinute:      oneMinute,
		fiveMinutes:    fiveMinutes,
		fifteenMinutes: fifteenMinutes,
	}
}

func (l LoadCollector) Collect() models.Result {
	stats := readLoadAverage()

	return models.Result{
		Name: "load",
		Data: map[string]interface{}{
			"1m":  stats.oneMinute,
			"5m":  stats.fiveMinutes,
			"15m": stats.fifteenMinutes,
		},
	}
}
