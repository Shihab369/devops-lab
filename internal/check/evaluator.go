package check

import (
	"fmt"
	"math"

	"github.com/Shihab369/devops-lab/internal/models"
)

type Thresholds struct {
	CPUMaxPercent    float64
	MemoryMaxPercent float64
	DiskMaxPercent   float64
	Load1MMax        float64
}

type Item struct {
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	Threshold float64 `json:"threshold"`
	Unit      string  `json:"unit"`
	Status    string  `json:"status"`
	Message   string  `json:"message"`
}

type Summary struct {
	Status string `json:"status"`
	Passed int    `json:"passed"`
	Failed int    `json:"failed"`
	Items  []Item `json:"items"`
}

func Evaluate(results []models.Result, thresholds Thresholds) Summary {
	items := []Item{
		evaluateCPU(results, thresholds.CPUMaxPercent),
		evaluateMemory(results, thresholds.MemoryMaxPercent),
		evaluateDisk(results, thresholds.DiskMaxPercent),
		evaluateLoad(results, thresholds.Load1MMax),
	}

	summary := Summary{Status: "healthy", Items: items}

	for _, item := range items {
		if item.Status == "pass" {
			summary.Passed++
			continue
		}

		summary.Failed++
		summary.Status = "unhealthy"
	}

	return summary
}

func ToResult(summary Summary) models.Result {
	return models.Result{
		Name: "check",
		Data: map[string]interface{}{
			"status": summary.Status,
			"passed": summary.Passed,
			"failed": summary.Failed,
			"items":  summary.Items,
		},
	}
}

func evaluateCPU(results []models.Result, max float64) Item {
	value, ok := metricFloat(results, "cpu", "usage_percent")
	if !ok {
		return missing("cpu", max, "%")
	}

	return thresholdItem("cpu", value, max, "%")
}

func evaluateMemory(results []models.Result, max float64) Item {
	total, totalOK := metricFloat(results, "memory", "total_bytes")
	available, availableOK := metricFloat(results, "memory", "available_bytes")
	if !totalOK || !availableOK || total <= 0 {
		return missing("memory", max, "%")
	}

	usedPercent := (total - available) / total * 100
	return thresholdItem("memory", usedPercent, max, "%")
}

func evaluateDisk(results []models.Result, max float64) Item {
	total, totalOK := metricFloat(results, "disk", "total_bytes")
	used, usedOK := metricFloat(results, "disk", "used_bytes")
	if !totalOK || !usedOK || total <= 0 {
		return missing("disk", max, "%")
	}

	usedPercent := used / total * 100
	return thresholdItem("disk", usedPercent, max, "%")
}

func evaluateLoad(results []models.Result, max float64) Item {
	value, ok := metricFloat(results, "load", "1m")
	if !ok {
		return missing("load_1m", max, "")
	}

	return thresholdItem("load_1m", value, max, "")
}

func thresholdItem(name string, value, threshold float64, unit string) Item {
	if value <= threshold {
		return Item{
			Name:      name,
			Value:     round2(value),
			Threshold: threshold,
			Unit:      unit,
			Status:    "pass",
			Message:   fmt.Sprintf("%s within threshold", name),
		}
	}

	return Item{
		Name:      name,
		Value:     round2(value),
		Threshold: threshold,
		Unit:      unit,
		Status:    "fail",
		Message:   fmt.Sprintf("%s exceeds threshold", name),
	}
}

func missing(name string, threshold float64, unit string) Item {
	return Item{
		Name:      name,
		Value:     0,
		Threshold: threshold,
		Unit:      unit,
		Status:    "fail",
		Message:   fmt.Sprintf("%s metric unavailable", name),
	}
}

func metricFloat(results []models.Result, resultName, key string) (float64, bool) {
	for _, result := range results {
		if result.Name != resultName {
			continue
		}

		value, ok := result.Data[key]
		if !ok {
			return 0, false
		}

		switch number := value.(type) {
		case float64:
			return number, true
		case float32:
			return float64(number), true
		case uint64:
			return float64(number), true
		case uint32:
			return float64(number), true
		case uint:
			return float64(number), true
		case int64:
			return float64(number), true
		case int32:
			return float64(number), true
		case int:
			return float64(number), true
		}

		return 0, false
	}

	return 0, false
}

func round2(value float64) float64 {
	return math.Round(value*100) / 100
}
