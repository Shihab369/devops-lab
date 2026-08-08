package load

import "testing"

func TestReadLoadAverage(t *testing.T) {
	stats := readLoadAverage()

	if stats.oneMinute < 0 {
		t.Fatal("1-minute load average cannot be negative")
	}

	if stats.fiveMinutes < 0 {
		t.Fatal("5-minute load average cannot be negative")
	}

	if stats.fifteenMinutes < 0 {
		t.Fatal("15-minute load average cannot be negative")
	}
}

func TestLoadCollector(t *testing.T) {
	collector := LoadCollector{}

	result := collector.Collect()

	if result.Name != "load" {
		t.Fatalf("expected result name load, got %s", result.Name)
	}

	if result.Data == nil {
		t.Fatal("expected result data")
	}
}
