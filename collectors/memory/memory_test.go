package memory

import "testing"

func TestReadMemoryStats(t *testing.T) {
	result := readMemoryStats()

	if result["total_bytes"] == 0 {
		t.Fatal("expected total memory")
	}

	if result["available_bytes"] == 0 {
		t.Fatal("expected available memory")
	}

	if result["free_bytes"] == 0 {
		t.Fatal("expected free memory")
	}

	if result["total_bytes"]%1024 != 0 {
		t.Fatal("expected memory values in bytes")
	}
}
