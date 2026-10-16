package check

import (
	"testing"

	"github.com/Shihab369/devops-lab/internal/models"
)

func TestEvaluateHealthy(t *testing.T) {
	results := []models.Result{
		{Name: "cpu", Data: map[string]interface{}{"usage_percent": 20.0}},
		{Name: "memory", Data: map[string]interface{}{"total_bytes": uint64(1000), "available_bytes": uint64(300)}},
		{Name: "disk", Data: map[string]interface{}{"total_bytes": uint64(2000), "used_bytes": uint64(1000)}},
		{Name: "load", Data: map[string]interface{}{"1m": 0.8}},
	}

	summary := Evaluate(results, Thresholds{
		CPUMaxPercent:    80,
		MemoryMaxPercent: 85,
		DiskMaxPercent:   90,
		Load1MMax:        2,
	})

	if summary.Status != "healthy" {
		t.Fatalf("expected healthy status, got %s", summary.Status)
	}

	if summary.Passed != 4 || summary.Failed != 0 {
		t.Fatalf("expected 4 passed and 0 failed, got %d passed and %d failed", summary.Passed, summary.Failed)
	}
}

func TestEvaluateUnhealthy(t *testing.T) {
	results := []models.Result{
		{Name: "cpu", Data: map[string]interface{}{"usage_percent": 95.0}},
		{Name: "memory", Data: map[string]interface{}{"total_bytes": uint64(1000), "available_bytes": uint64(50)}},
		{Name: "disk", Data: map[string]interface{}{"total_bytes": uint64(2000), "used_bytes": uint64(1900)}},
		{Name: "load", Data: map[string]interface{}{"1m": 6.1}},
	}

	summary := Evaluate(results, Thresholds{
		CPUMaxPercent:    80,
		MemoryMaxPercent: 85,
		DiskMaxPercent:   90,
		Load1MMax:        2,
	})

	if summary.Status != "unhealthy" {
		t.Fatalf("expected unhealthy status, got %s", summary.Status)
	}

	if summary.Passed != 0 || summary.Failed != 4 {
		t.Fatalf("expected 0 passed and 4 failed, got %d passed and %d failed", summary.Passed, summary.Failed)
	}
}

func TestEvaluateMissingDataFails(t *testing.T) {
	results := []models.Result{
		{Name: "cpu", Data: map[string]interface{}{"usage_percent": 25.0}},
	}

	summary := Evaluate(results, Thresholds{
		CPUMaxPercent:    80,
		MemoryMaxPercent: 85,
		DiskMaxPercent:   90,
		Load1MMax:        2,
	})

	if summary.Status != "unhealthy" {
		t.Fatalf("expected unhealthy status, got %s", summary.Status)
	}

	if summary.Passed != 1 || summary.Failed != 3 {
		t.Fatalf("expected 1 passed and 3 failed, got %d passed and %d failed", summary.Passed, summary.Failed)
	}
}

func TestToResult(t *testing.T) {
	summary := Summary{Status: "healthy", Passed: 4, Failed: 0, Items: []Item{{Name: "cpu", Status: "pass"}}}
	result := ToResult(summary)

	if result.Name != "check" {
		t.Fatalf("expected check result name, got %s", result.Name)
	}

	status, ok := result.Data["status"].(string)
	if !ok || status != "healthy" {
		t.Fatal("expected healthy status in result data")
	}
}
