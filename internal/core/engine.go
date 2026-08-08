package core

import "github.com/Shihab369/devops-lab/internal/models"

type Engine struct {
	results []models.Result
}

func NewEngine() *Engine {
	return &Engine{
		results: make([]models.Result, 0),
	}
}

func (e *Engine) AddResult(result models.Result) {
	e.results = append(e.results, result)
}

func (e *Engine) Results() []models.Result {
	return e.results
}
