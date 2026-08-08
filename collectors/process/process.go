package process

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Shihab369/devops-thinking-lab/internal/core"
	"github.com/Shihab369/devops-thinking-lab/internal/models"
)

type ProcessCollector struct{}

var _ core.Collector = ProcessCollector{}

func (p ProcessCollector) Collect() models.Result {
	process := readProcessStatus(1)

	return models.Result{
		Name: "process",
		Data: map[string]interface{}{
			"pid":   process.pid,
			"ppid":  process.ppid,
			"name":  process.name,
			"state": process.state,
		},
	}
}

type processStatus struct {
	pid   int
	ppid  int
	name  string
	state string
}

func readProcessStatus(pid int) processStatus {
	path := filepath.Join("/proc", strconv.Itoa(pid), "status")

	data, err := os.ReadFile(path)
	if err != nil {
		return processStatus{}
	}

	var result processStatus

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		fields := strings.SplitN(line, ":", 2)

		if len(fields) != 2 {
			continue
		}

		key := strings.TrimSpace(fields[0])
		value := strings.TrimSpace(fields[1])

		switch key {
		case "Name":
			result.name = value

		case "State":
			result.state = value

		case "Pid":
			result.pid, _ = strconv.Atoi(value)

		case "PPid":
			result.ppid, _ = strconv.Atoi(value)
		}
	}

	return result
}
