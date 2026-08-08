package core

import "github.com/Shihab369/devops-lab/internal/models"

type Runner struct {
	collectors []Collector
}

func NewRunner(collectors ...Collector) Runner {
	return Runner{
		collectors: collectors,
	}
}

func (r Runner) Run() []models.Result {
	results := make([]models.Result, 0, len(r.collectors))

	for _, collector := range r.collectors {
		results = append(results, collector.Collect())
	}

	return results
}
