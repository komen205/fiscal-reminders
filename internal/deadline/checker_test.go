package deadline

import (
	"testing"
	"time"

	"github.com/komen205/fiscal-reminders/internal/config"
)

func TestChecker_Creation(t *testing.T) {
	cfg := &config.Config{
		NtfyTopic:       "test",
		NtfyServer:      "https://ntfy.sh",
		DaysBeforeAlert: []int{7, 3, 1, 0},
	}

	checker := NewChecker(cfg)

	if checker == nil {
		t.Fatal("NewChecker returned nil")
	}

	if checker.config != cfg {
		t.Error("config not set correctly")
	}

	if checker.notifier == nil {
		t.Error("notifier not initialized")
	}
}

func TestChecker_YearlyDeadlineCalculation(t *testing.T) {
	// Test that yearly deadlines are calculated correctly
	now := time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)

	// Deadline on Jan 31 - should be 16 days away
	deadline := Deadline{Month: 1, Day: 31}
	deadlineDate := time.Date(now.Year(), time.Month(deadline.Month), deadline.Day, 23, 59, 59, 0, now.Location())

	daysUntil := int(deadlineDate.Sub(now).Hours() / 24)

	if daysUntil < 15 || daysUntil > 17 {
		t.Errorf("expected ~16 days until Jan 31 from Jan 15, got %d", daysUntil)
	}
}

func TestChecker_PastDeadlineRollsToNextYear(t *testing.T) {
	// If deadline passed this year, should calculate for next year
	now := time.Date(2026, 2, 15, 12, 0, 0, 0, time.UTC)

	// Deadline on Jan 31 - already passed
	deadline := Deadline{Month: 1, Day: 31}
	year := now.Year()

	deadlineDate := time.Date(year, time.Month(deadline.Month), deadline.Day, 23, 59, 59, 0, now.Location())

	// Should be in the past
	if !deadlineDate.Before(now) {
		t.Error("Jan 31 should be before Feb 15")
	}

	// Roll to next year
	deadlineDate = time.Date(year+1, time.Month(deadline.Month), deadline.Day, 23, 59, 59, 0, now.Location())

	// Should be in the future now
	if !deadlineDate.After(now) {
		t.Error("Jan 31 next year should be after Feb 15 this year")
	}
}

func TestChecker_MonthlyDeadlineCalculation(t *testing.T) {
	now := time.Date(2026, 3, 10, 12, 0, 0, 0, time.UTC)

	// Monthly deadline on 20th
	deadline := Deadline{Month: 0, Day: 20}
	deadlineDate := time.Date(now.Year(), now.Month(), deadline.Day, 23, 59, 59, 0, now.Location())

	daysUntil := int(deadlineDate.Sub(now).Hours() / 24)

	if daysUntil < 9 || daysUntil > 11 {
		t.Errorf("expected ~10 days until 20th from 10th, got %d", daysUntil)
	}
}

func TestChecker_MonthlyDeadlinePastRollsToNextMonth(t *testing.T) {
	now := time.Date(2026, 3, 25, 12, 0, 0, 0, time.UTC)

	// Monthly deadline on 20th - already passed
	deadline := Deadline{Month: 0, Day: 20}
	deadlineDate := time.Date(now.Year(), now.Month(), deadline.Day, 23, 59, 59, 0, now.Location())

	// Should be in the past
	if !deadlineDate.Before(now) {
		t.Error("20th should be before 25th")
	}

	// Roll to next month
	deadlineDate = deadlineDate.AddDate(0, 1, 0)

	// Should be in April now
	if deadlineDate.Month() != time.April {
		t.Errorf("expected April, got %v", deadlineDate.Month())
	}
}

func TestChecker_AlertDayMatching(t *testing.T) {
	alertDays := []int{7, 3, 1, 0}

	tests := []struct {
		daysUntil   int
		shouldAlert bool
	}{
		{7, true},
		{3, true},
		{1, true},
		{0, true},
		{5, false},
		{2, false},
		{10, false},
	}

	for _, tt := range tests {
		found := false
		for _, alertDay := range alertDays {
			if alertDay == tt.daysUntil {
				found = true
				break
			}
		}

		if found != tt.shouldAlert {
			t.Errorf("daysUntil=%d: expected shouldAlert=%v, got %v",
				tt.daysUntil, tt.shouldAlert, found)
		}
	}
}

func TestDaysUntil_Precision(t *testing.T) {
	// Test that DaysUntil handles edge cases
	now := time.Now()

	tests := []struct {
		name     string
		deadline time.Time
		minDays  int
		maxDays  int
	}{
		{"1 hour from now", now.Add(1 * time.Hour), 0, 0},
		{"23 hours from now", now.Add(23 * time.Hour), 0, 1},
		{"25 hours from now", now.Add(25 * time.Hour), 1, 1},
		{"7 days from now", now.AddDate(0, 0, 7), 6, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			days := DaysUntil(tt.deadline)
			if days < tt.minDays || days > tt.maxDays {
				t.Errorf("expected days in range [%d, %d], got %d",
					tt.minDays, tt.maxDays, days)
			}
		})
	}
}
