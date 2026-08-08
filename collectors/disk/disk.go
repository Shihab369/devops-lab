package disk

import (
	"os"
	"strings"

	"github.com/Shihab369/devops-thinking-lab/internal/core"
	"github.com/Shihab369/devops-thinking-lab/internal/models"
)

type DiskCollector struct{}

var _ core.Collector = DiskCollector{}

func (d DiskCollector) Collect() models.Result {
	return models.Result{
		Name: "disk",
		Data: map[string]interface{}{
			"mount_points": readMountPoints(),
		},
	}
}

func readMountPoints() []string {
	data, err := os.ReadFile("/proc/mounts")
	if err != nil {
		return []string{}
	}

	var mounts []string

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) < 2 {
			continue
		}

		mounts = append(mounts, fields[1])
	}

	return mounts
}
