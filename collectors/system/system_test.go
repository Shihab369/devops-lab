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
