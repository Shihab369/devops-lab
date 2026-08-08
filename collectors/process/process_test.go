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

func TestListPIDs(t *testing.T) {
	pids := listPIDs()

	if len(pids) == 0 {
		t.Fatal("expected at least one process")
	}

	foundPID1 := false

	for _, pid := range pids {
		if pid == "1" {
			foundPID1 = true
			break
		}
	}

	if !foundPID1 {
		t.Fatal("expected to find PID 1")
	}
}
