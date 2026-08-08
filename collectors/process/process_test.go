package process

import "testing"

func TestReadProcessStatus(t *testing.T) {
	result := readProcessStatus(1)

	if result.pid != 1 {
		t.Fatalf("expected PID 1, got %d", result.pid)
	}

	if result.name == "" {
		t.Fatal("expected process name")
	}

	if result.state == "" {
		t.Fatal("expected process state")
	}
}

func TestProcessCollector(t *testing.T) {
	collector := ProcessCollector{}

	result := collector.Collect()

	if result.Name != "process" {
		t.Fatalf("expected result name process, got %s", result.Name)
	}

	if result.Data["pid"] != 1 {
		t.Fatalf("expected PID 1, got %v", result.Data["pid"])
	}
}
