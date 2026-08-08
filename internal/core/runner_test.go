package core

import (
	"testing"

	"github.com/Shihab369/devops-thinking-lab/internal/models"
)

type testCollector struct {
	name string
}

func (c testCollector) Collect() models.Result {
	return models.Result{
		Name: c.name,
		Data: map[string]interface{}{
			"test": true,
		},
	}
}

func TestRunner(t *testing.T) {
	runner := NewRunner(
		testCollector{name: "first"},
		testCollector{name: "second"},
	)

	results := runner.Run()

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}

	if results[0].Name != "first" {
		t.Fatalf("expected first result, got %q", results[0].Name)
	}

	if results[1].Name != "second" {
		t.Fatalf("expected second result, got %q", results[1].Name)
	}
}
