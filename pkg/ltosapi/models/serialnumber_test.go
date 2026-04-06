package models

import (
	"encoding/json"
	"testing"
)

func TestSerialNumber_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected SerialNumber
	}{
		{"valid serial", `"ABC123"`, "ABC123"},
		{"whitespace trimmed", `"  ABC123  "`, "ABC123"},
		{"empty string", `""`, ""},
		{"unknown", `"unknown"`, ""},
		{"Unknown mixed case", `"Unknown"`, ""},
		{"n/a", `"n/a"`, ""},
		{"na", `"na"`, ""},
		{"none", `"none"`, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sn SerialNumber
			if err := json.Unmarshal([]byte(tt.input), &sn); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if sn != tt.expected {
				t.Errorf("got %q, want %q", sn, tt.expected)
			}
		})
	}
}
