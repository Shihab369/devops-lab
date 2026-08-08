package process

import (
	"os"
	"strconv"
	"strings"

	"github.com/Shihab369/devops-thinking-lab/internal/core"
	"github.com/Shihab369/devops-thinking-lab/internal/models"
)

type ProcessInfo struct {
	PID   int
	PPID  int
	Name  string
	State string
}

type ProcessCollector struct {
	PID int
}

var _ core.Collector = ProcessCollector{}

func (p ProcessCollector) Collect() models.Result {
	info := readProcessStatus(p.PID)

	return models.Result{
		Name: "process",
		Data: map[string]interface{}{
			"pid":   info.PID,
			"ppid":  info.PPID,
			"name":  info.Name,
			"state": info.State,
		},
	}
}

func readProcessStatus(pid int) ProcessInfo {
	path := "/proc/" + strconv.Itoa(pid) + "/status"

	data, err := os.ReadFile(path)
	if err != nil {
		return ProcessInfo{}
	}

	var info ProcessInfo

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		key, value, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}

		value = strings.TrimSpace(value)

		switch key {
		case "Name":
			info.Name = value

		case "State":
			fields := strings.Fields(value)
			if len(fields) > 0 {
				info.State = fields[0]
			}

		case "Pid":
			info.PID, _ = strconv.Atoi(value)

		case "PPid":
			info.PPID, _ = strconv.Atoi(value)
		}
	}

	return info
}
