package cpu

import "testing"

func TestReadCPUStats(t *testing.T) {
	stats := readCPUStats()

	if stats.user == 0 {
		t.Fatal("expected user CPU time")
	}

	if stats.system == 0 {
		t.Fatal("expected system CPU time")
	}

	if stats.idle == 0 {
		t.Fatal("expected idle CPU time")
	}

	total := stats.total()
	idle := stats.idleTime()

	if total == 0 {
		t.Fatal("expected total CPU time")
	}

	if idle == 0 {
		t.Fatal("expected idle CPU time")
	}

	if idle > total {
		t.Fatal("idle CPU time cannot exceed total CPU time")
	}
}

func TestCPUUsagePercent(t *testing.T) {
	previous := cpuStats{
		user:   100,
		system: 100,
		idle:   700,
		iowait: 0,
	}

	current := cpuStats{
		user:   150,
		system: 100,
		idle:   750,
		iowait: 0,
	}

	usage := cpuUsagePercent(previous, current)

	if usage != 50 {
		t.Fatalf("expected CPU usage 50%%, got %f%%", usage)
	}
}

func TestCPUCollector(t *testing.T) {
	collector := CPUCollector{}

	result := collector.Collect()

	if result.Name != "cpu" {
		t.Fatalf("expected collector name cpu, got %s", result.Name)
	}

	usage, ok := result.Data["usage_percent"].(float64)
	if !ok {
		t.Fatal("expected usage_percent to be float64")
	}

	if usage < 0 || usage > 100 {
		t.Fatalf("expected CPU usage between 0 and 100, got %f", usage)
	}
}
