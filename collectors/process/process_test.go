package process

import "testing"

func TestReadProcessStatus(t *testing.T) {
	info := readProcessStatus(1)

	if info.PID != 1 {
		t.Fatalf("expected PID 1, got %d", info.PID)
	}

	if info.Name == "" {
		t.Fatal("expected process name")
	}

	if info.State == "" {
		t.Fatal("expected process state")
	}
}

func TestProcessCollector(t *testing.T) {
	collector := ProcessCollector{
		PID: 1,
	}

	result := collector.Collect()

	if result.Name != "process" {
		t.Fatalf("expected result name process, got %s", result.Name)
	}

	if result.Data["pid"] != 1 {
		t.Fatalf("expected pid 1, got %v", result.Data["pid"])
	}

	if result.Data["name"] == "" {
		t.Fatal("expected process name")
	}

	if result.Data["state"] == "" {
		t.Fatal("expected process state")
	}
}
