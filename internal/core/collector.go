package core

import "github.com/Shihab369/devops-thinking-lab/internal/models"

type Collector interface {
	Collect() models.Result
}
