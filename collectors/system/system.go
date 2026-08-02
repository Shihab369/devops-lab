package system

import (
	"github.com/Shihab369/devops-thinking-lab/internal/core"
	"github.com/Shihab369/devops-thinking-lab/internal/models"
)

type SystemCollector struct{}

var _ core.Collector = SystemCollector{}

func (s SystemCollector) Collect() models.Result {
	return models.Result{
		Name: "system",
	}
}
