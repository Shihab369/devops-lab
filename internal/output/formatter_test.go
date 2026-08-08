package output

import (
	"testing"

	"github.com/Shihab369/devops-thinking-lab/internal/models"
)

func TestInterfaceTypes(t *testing.T) {
	var value interface{}

	value = "hello"

	t.Logf("value: %v, type: %T", value, value)

	value = 42

	t.Logf("value: %v, type: %T", value, value)

	value = 3.14

	t.Logf("value: %v, type: %T", value, value)
}

func TestTypeAssertion(t *testing.T) {
	var value interface{} = "hello"

	text, ok := value.(string)

	if !ok {
		t.Fatal("expected string")
	}

	t.Logf("text: %s", text)
}

func TestNestedMap(t *testing.T) {
	data := map[string]interface{}{
		"hostname": "banded",
		"memory": map[string]interface{}{
			"total_bytes":     uint64(16000000000),
			"available_bytes": uint64(10000000000),
		},
	}

	memory, ok := data["memory"].(map[string]interface{})

	if !ok {
		t.Fatal("expected memory to be a nested map")
	}

	total, ok := memory["total_bytes"].(uint64)

	if !ok {
		t.Fatal("expected total_bytes to be uint64")
	}

	t.Logf("memory total: %d", total)
}

func TestMapTypes(t *testing.T) {
	data := map[string]interface{}{
		"memory": map[string]interface{}{
			"total": uint64(100),
		},
		"os": map[string]string{
			"name": "Ubuntu",
		},
	}

	_, memoryOK := data["memory"].(map[string]interface{})
	_, osOK := data["os"].(map[string]string)

	t.Logf("memory is map[string]interface{}: %v", memoryOK)
	t.Logf("os is map[string]string: %v", osOK)

	if !memoryOK {
		t.Fatal("memory should be map[string]interface{}")
	}

	if !osOK {
		t.Fatal("os should be map[string]string")
	}
}

func TestPrintResults(t *testing.T) {
	results := []models.Result{
		{
			Name: "memory",
			Data: map[string]interface{}{
				"total_bytes":     uint64(16000000000),
				"available_bytes": uint64(10000000000),
			},
		},
	}

	PrintResults(results)
}
