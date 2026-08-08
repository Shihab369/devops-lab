package core

import "github.com/Shihab369/devops-lab/internal/models"

type Collector interface {
	Collect() models.Result
}
