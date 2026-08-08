package system

import "testing"

func TestReadOSRelease(t *testing.T) {
	result := readOSRelease()

	if result["ID"] == "" {
		t.Fatal("expected OS ID")
	}

	if result["NAME"] == "" {
		t.Fatal("expected OS name")
	}

	if result["VERSION_ID"] == "" {
		t.Fatal("expected OS version")
	}
}

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
		t.Fatal("expected total memory in bytes")
	}
}
