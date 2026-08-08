package cpu

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Shihab369/devops-thinking-lab/internal/core"
	"github.com/Shihab369/devops-thinking-lab/internal/models"
)

type CPUCollector struct{}

var _ core.Collector = CPUCollector{}

func (c CPUCollector) Collect() models.Result {
	previous := readCPUStats()

	time.Sleep(100 * time.Millisecond)

	current := readCPUStats()

	usage := cpuUsagePercent(previous, current)

	return models.Result{
		Name: "cpu",
		Data: map[string]interface{}{
			"usage_percent": usage,
		},
	}
}

type cpuStats struct {
	user      uint64
	nice      uint64
	system    uint64
	idle      uint64
	iowait    uint64
	irq       uint64
	softirq   uint64
	steal     uint64
	guest     uint64
	guestNice uint64
}

func readCPUStats() cpuStats {
	data, err := os.ReadFile("/proc/stat")
	if err != nil {
		return cpuStats{}
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if !strings.HasPrefix(line, "cpu ") {
			continue
		}

		fields := strings.Fields(line)

		if len(fields) < 11 {
			return cpuStats{}
		}

		values := make([]uint64, 10)

		for i := 0; i < 10; i++ {
			value, err := strconv.ParseUint(fields[i+1], 10, 64)
			if err != nil {
				return cpuStats{}
			}

			values[i] = value
		}

		return cpuStats{
			user:      values[0],
			nice:      values[1],
			system:    values[2],
			idle:      values[3],
			iowait:    values[4],
			irq:       values[5],
			softirq:   values[6],
			steal:     values[7],
			guest:     values[8],
			guestNice: values[9],
		}
	}

	return cpuStats{}
}

func (s cpuStats) total() uint64 {
	return s.user +
		s.nice +
		s.system +
		s.idle +
		s.iowait +
		s.irq +
		s.softirq +
		s.steal +
		s.guest +
		s.guestNice
}

func (s cpuStats) idleTime() uint64 {
	return s.idle + s.iowait
}

func cpuUsagePercent(previous, current cpuStats) float64 {
	previousTotal := previous.total()
	currentTotal := current.total()

	previousIdle := previous.idleTime()
	currentIdle := current.idleTime()

	totalDelta := currentTotal - previousTotal
	idleDelta := currentIdle - previousIdle

	if totalDelta == 0 {
		return 0
	}

	busyDelta := totalDelta - idleDelta

	return float64(busyDelta) / float64(totalDelta) * 100
}
