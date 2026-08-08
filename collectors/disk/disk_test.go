package disk

import "testing"

func TestReadDiskStats(t *testing.T) {
	stats := readDiskStats("/")

	if stats["total_bytes"] == 0 {
		t.Fatal("expected total disk space")
	}

	if stats["free_bytes"] == 0 {
		t.Fatal("expected free disk space")
	}

	if stats["used_bytes"] == 0 {
		t.Fatal("expected used disk space")
	}

	if stats["used_bytes"] > stats["total_bytes"] {
		t.Fatal("used disk space cannot exceed total disk space")
	}
}
