package memory

import (
	"strings"
	"testing"
)

func TestParseMemInfo(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		want        memoryStats
		wantErr     bool
		errContains string
	}{
		{
			name: "valid input",
			input: strings.Join([]string{
				"MemTotal:       1000 kB",
				"MemAvailable:    400 kB",
				"MemFree:         300 kB",
				"Buffers:         200 kB",
			}, "\n"),
			want: memoryStats{
				totalBytes:     1000 * 1024,
				availableBytes: 400 * 1024,
				freeBytes:      300 * 1024,
			},
		},
		{
			name: "valid zero value",
			input: strings.Join([]string{
				"MemTotal:       1 kB",
				"MemAvailable:   1 kB",
				"MemFree:        0 kB",
			}, "\n"),
			want: memoryStats{
				totalBytes:     1024,
				availableBytes: 1024,
				freeBytes:      0,
			},
		},
		{
			name:        "empty input",
			input:       "   \n  ",
			wantErr:     true,
			errContains: "empty",
		},
		{
			name: "missing required field",
			input: strings.Join([]string{
				"MemTotal:       1000 kB",
				"MemFree:        200 kB",
			}, "\n"),
			wantErr:     true,
			errContains: "missing required fields",
		},
		{
			name: "malformed numeric value",
			input: strings.Join([]string{
				"MemTotal:       abc kB",
				"MemAvailable:   100 kB",
				"MemFree:        50 kB",
			}, "\n"),
			wantErr:     true,
			errContains: "invalid numeric value",
		},
		{
			name: "incomplete required value",
			input: strings.Join([]string{
				"MemTotal:",
				"MemAvailable:   100 kB",
				"MemFree:        50 kB",
			}, "\n"),
			wantErr:     true,
			errContains: "incomplete required field MemTotal",
		},
		{
			name: "incomplete required field without separator",
			input: strings.Join([]string{
				"MemTotal 1000 kB",
				"MemAvailable:   100 kB",
				"MemFree:        50 kB",
			}, "\n"),
			wantErr:     true,
			errContains: "incomplete required field MemTotal",
		},
		{
			name: "invalid unit",
			input: strings.Join([]string{
				"MemTotal:       1000 MB",
				"MemAvailable:   100 kB",
				"MemFree:        50 kB",
			}, "\n"),
			wantErr:     true,
			errContains: "unexpected unit",
		},
		{
			name: "kb to bytes conversion",
			input: strings.Join([]string{
				"MemTotal:       2 kB",
				"MemAvailable:   1 kB",
				"MemFree:        1 kB",
			}, "\n"),
			want: memoryStats{
				totalBytes:     2048,
				availableBytes: 1024,
				freeBytes:      1024,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseMemInfo(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.errContains)
				}

				if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Fatalf("expected error containing %q, got %q", tt.errContains, err.Error())
				}

				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if got != tt.want {
				t.Fatalf("expected %+v, got %+v", tt.want, got)
			}
		})
	}
}

func TestMemoryCollectorCollect(t *testing.T) {
	collector := MemoryCollector{}
	result := collector.Collect()

	if result.Name != "memory" {
		t.Fatalf("expected collector name memory, got %s", result.Name)
	}

	for _, key := range []string{"total_bytes", "available_bytes", "free_bytes"} {
		value, ok := result.Data[key]
		if !ok {
			t.Fatalf("expected key %s", key)
		}

		if _, ok := value.(uint64); !ok {
			t.Fatalf("expected %s to be uint64", key)
		}
	}

	if errValue, ok := result.Data["collection_error"]; ok {
		errMessage, ok := errValue.(string)
		if !ok || strings.TrimSpace(errMessage) == "" {
			t.Fatal("expected collection_error to be a non-empty string")
		}
	}
}
