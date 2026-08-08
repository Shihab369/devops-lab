package core_test

import (
	"testing"

	"github.com/Shihab369/devops-thinking-lab/collectors/cpu"
	"github.com/Shihab369/devops-thinking-lab/collectors/disk"
	"github.com/Shihab369/devops-thinking-lab/collectors/memory"
	"github.com/Shihab369/devops-thinking-lab/collectors/system"
	"github.com/Shihab369/devops-thinking-lab/internal/core"
)

func TestRunnerWithRealCollectors(t *testing.T) {
	runner := core.NewRunner(
		system.SystemCollector{},
		memory.MemoryCollector{},
		disk.DiskCollector{},
		cpu.CPUCollector{},
	)

	results := runner.Run()

	if len(results) != 4 {
		t.Fatalf("expected 4 results, got %d", len(results))
	}

	expectedNames := []string{
		"system",
		"memory",
		"disk",
		"cpu",
	}

	for i, expected := range expectedNames {
		if results[i].Name != expected {
			t.Fatalf(
				"expected result %d to be %q, got %q",
				i,
				expected,
				results[i].Name,
			)
		}
	}
}
