package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTimeQuality_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  time.Duration
		expectErr bool
	}{
		{"microseconds", `"less-than-100µs"`, 100 * time.Microsecond, false},
		{"milliseconds", `"less-than-25ms"`, 25 * time.Millisecond, false},
		{"uppercase prefix", `"Less-Than-1s"`, 1 * time.Second, false},
		{"no prefix", `"500ms"`, 500 * time.Millisecond, false},
		{"invalid duration", `"not-a-duration"`, 0, true},
		{"not a string", `123`, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tq TimeQuality
			err := json.Unmarshal([]byte(tt.input), &tq)
			if tt.expectErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if time.Duration(tq) != tt.expected {
				t.Errorf("got %v, want %v", time.Duration(tq), tt.expected)
			}
		})
	}
}

func TestTimeQuality_Seconds(t *testing.T) {
	tq := TimeQuality(100 * time.Microsecond)
	got := tq.Seconds()
	want := 0.0001
	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

func TestClockStatus_IsSynchronized(t *testing.T) {
	if !(ClockStatus{Clock: "synchronized"}).IsSynchronized() {
		t.Error("expected true for 'synchronized'")
	}
	if (ClockStatus{Clock: "not-synchronized"}).IsSynchronized() {
		t.Error("expected false for 'not-synchronized'")
	}
}

func TestClockStatus_IsOscillatorWarmedUp(t *testing.T) {
	if !(ClockStatus{Oscillator: "warmed-up"}).IsOscillatorWarmedUp() {
		t.Error("expected true for 'warmed-up'")
	}
	if (ClockStatus{Oscillator: "warming-up"}).IsOscillatorWarmedUp() {
		t.Error("expected false for 'warming-up'")
	}
}
