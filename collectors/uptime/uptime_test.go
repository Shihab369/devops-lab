package uptime

import "testing"

func TestReadUptime(t *testing.T) {
	uptime := readUptime()

	if uptime <= 0 {
		t.Fatalf("expected positive uptime, got %f", uptime)
	}
}

func TestUptimeCollector(t *testing.T) {
	collector := UptimeCollector{}

	result := collector.Collect()

	if result.Name != "uptime" {
		t.Fatalf("expected result name uptime, got %s", result.Name)
	}

	value, ok := result.Data["uptime_seconds"].(float64)
	if !ok {
		t.Fatal("expected uptime_seconds to be float64")
	}

	if value <= 0 {
		t.Fatalf("expected positive uptime, got %f", value)
	}
}
