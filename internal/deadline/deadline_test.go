package deadline

import (
	"testing"
	"time"
)

func TestDeadline_IsMonthly(t *testing.T) {
	monthly := Deadline{Month: 0, Day: 20}
	yearly := Deadline{Month: 6, Day: 30}

	if !monthly.IsMonthly() {
		t.Error("month=0 should be monthly")
	}

	if yearly.IsMonthly() {
		t.Error("month=6 should not be monthly")
	}
}

func TestDeadline_HasTag(t *testing.T) {
	d := Deadline{Tags: []string{"iva", "trimestral"}}

	if !d.HasTag("iva") {
		t.Error("should have 'iva' tag")
	}

	if !d.HasTag("trimestral") {
		t.Error("should have 'trimestral' tag")
	}

	if d.HasTag("irs") {
		t.Error("should not have 'irs' tag")
	}
}

func TestAllDeadlines_Count(t *testing.T) {
	// Should have at least 10 deadlines defined
	if len(All) < 10 {
		t.Errorf("expected at least 10 deadlines, got %d", len(All))
	}
}

func TestAllDeadlines_Valid(t *testing.T) {
	for _, d := range All {
		if d.Name == "" {
			t.Error("deadline has empty name")
		}

		if d.Description == "" {
			t.Error("deadline has empty description")
		}

		if d.Day < 1 || d.Day > 31 {
			t.Errorf("deadline %s has invalid day: %d", d.Name, d.Day)
		}

		if d.Month < 0 || d.Month > 12 {
			t.Errorf("deadline %s has invalid month: %d", d.Name, d.Month)
		}

		if d.Priority == "" {
			t.Errorf("deadline %s has no priority", d.Name)
		}

		if len(d.Tags) == 0 {
			t.Errorf("deadline %s has no tags", d.Name)
		}
	}
}

func TestDaysUntil(t *testing.T) {
	now := time.Now()

	tomorrow := now.AddDate(0, 0, 1)
	days := DaysUntil(tomorrow)
	if days < 0 || days > 1 {
		t.Errorf("days until tomorrow should be 0-1, got %d", days)
	}

	nextWeek := now.AddDate(0, 0, 7)
	days = DaysUntil(nextWeek)
	if days < 6 || days > 7 {
		t.Errorf("days until next week should be 6-7, got %d", days)
	}
}

