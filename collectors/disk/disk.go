package disk

import (
	"github.com/Shihab369/devops-thinking-lab/internal/core"
	"github.com/Shihab369/devops-thinking-lab/internal/models"
	"golang.org/x/sys/unix"
)

type DiskCollector struct{}

var _ core.Collector = DiskCollector{}

func (d DiskCollector) Collect() models.Result {
	stats := readDiskStats("/")

	return models.Result{
		Name: "disk",
		Data: map[string]interface{}{
			"mount_point": "/",
			"total_bytes": stats["total_bytes"],
			"used_bytes":  stats["used_bytes"],
			"free_bytes":  stats["free_bytes"],
		},
	}
}

func readDiskStats(path string) map[string]uint64 {
	var stat unix.Statfs_t

	err := unix.Statfs(path, &stat)
	if err != nil {
		return map[string]uint64{}
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bavail * uint64(stat.Bsize)
	used := total - free

	return map[string]uint64{
		"total_bytes": total,
		"used_bytes":  used,
		"free_bytes":  free,
	}
}
