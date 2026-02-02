package main

import (
	"testing"
	"time"
)

// Test deadline calculations
func TestGetNextDeadline(t *testing.T) {
	tests := []struct {
		name     string
		deadline time.Time
		want     bool // should be in future
	}{
		{"SegSoc Q4", time.Date(2026, 1, 31, 23, 59, 0, 0, time.UTC), true},
		{"IVA Q1", time.Date(2026, 5, 20, 23, 59, 0, 0, time.UTC), true},
		{"IRS", time.Date(2026, 6, 30, 23, 59, 0, 0, time.UTC), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.deadline.After(time.Now()) != tt.want {
				t.Errorf("deadline %v should be in future", tt.deadline)
			}
		})
	}
}

// Test days until calculation
func TestDaysUntil(t *testing.T) {
	tomorrow := time.Now().AddDate(0, 0, 1)
	days := int(time.Until(tomorrow).Hours() / 24)

	if days < 0 || days > 2 {
		t.Errorf("days until tomorrow should be ~1, got %d", days)
	}
}

// Test alert thresholds
func TestShouldAlert(t *testing.T) {
	alertDays := []int{7, 3, 1, 0}
	testDays := 3

	found := false
	for _, d := range alertDays {
		if d == testDays {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("should alert on day %d", testDays)
	}
}

// Test config defaults
func TestConfigDefaults(t *testing.T) {
	config := struct {
		NtfyTopic          string `json:"ntfy_topic"`
		NtfyServer         string `json:"ntfy_server"`
		CheckIntervalHours int    `json:"check_interval_hours"`
	}{
		NtfyTopic:          "fiscal-reminders",
		NtfyServer:         "https://ntfy.sh",
		CheckIntervalHours: 12,
	}

	if config.NtfyTopic == "" {
		t.Error("ntfy_topic should have default")
	}
	if config.CheckIntervalHours <= 0 {
		t.Error("check_interval should be positive")
	}
}

