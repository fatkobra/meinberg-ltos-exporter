package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestEvent_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		expectedType     string
		expectedName     string
		expectedLastUnix float64
		expectErr        bool
	}{
		{
			"triggered event",
			`{"type":"warning","object-id":"ntp-not-synchronized","last-triggered":"2025-03-15T12:30:00"}`,
			"warning", "ntp-not-synchronized",
			float64(time.Date(2025, 3, 15, 12, 30, 0, 0, time.UTC).Unix()),
			false,
		},
		{
			"never triggered",
			`{"type":"error","object-id":"antenna-fault","last-triggered":"never"}`,
			"error", "antenna-fault", 0, false,
		},
		{
			"invalid timestamp",
			`{"type":"warning","object-id":"test","last-triggered":"not-a-time"}`,
			"", "", 0, true,
		},
		{
			"invalid json",
			`{broken}`,
			"", "", 0, true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var e Event
			err := json.Unmarshal([]byte(tt.input), &e)
			if tt.expectErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if e.Type != tt.expectedType {
				t.Errorf("Type: got %q, want %q", e.Type, tt.expectedType)
			}
			if e.Name != tt.expectedName {
				t.Errorf("Name: got %q, want %q", e.Name, tt.expectedName)
			}
			if e.LastTriggeredUnix != tt.expectedLastUnix {
				t.Errorf("LastTriggeredUnix: got %f, want %f", e.LastTriggeredUnix, tt.expectedLastUnix)
			}
		})
	}
}
