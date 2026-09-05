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
	if len(All) != 13 {
		t.Errorf("expected 13 deadlines, got %d", len(All))
	}
}

func TestAllDeadlines_IRSPaymentsByAccount(t *testing.T) {
	wantMonths := map[int]bool{7: false, 9: false, 12: false}
	found := 0

	for _, d := range All {
		if !d.HasTag("pagamento-por-conta") {
			continue
		}

		found++
		if d.Day != 20 {
			t.Errorf("expected statutory day 20 for %s, got %d", d.Name, d.Day)
		}
		if !d.AdjustToNextBusinessDay {
			t.Errorf("expected business-day adjustment for %s", d.Name)
		}
		if _, ok := wantMonths[d.Month]; !ok {
			t.Errorf("unexpected payment month %d", d.Month)
		} else {
			wantMonths[d.Month] = true
		}
	}

	if found != 3 {
		t.Fatalf("expected 3 IRS payments by account, got %d", found)
	}
	for month, seen := range wantMonths {
		if !seen {
			t.Errorf("missing IRS payment for month %d", month)
		}
	}
}

func TestAllDeadlines_NoMonthlySocialSecurityPayment(t *testing.T) {
	for _, d := range All {
		if d.IsMonthly() && d.HasTag("seguranca-social") && d.HasTag("pagamento") {
			t.Errorf("monthly Social Security payment should not be configured: %s", d.Name)
		}
	}
}

func TestDeadline_DateForYearAdjustsIRS2026Weekends(t *testing.T) {
	location := time.FixedZone("Europe/Lisbon", 0)
	tests := []struct {
		name  string
		month int
		want  time.Time
	}{
		{"July stays on Monday 20", 7, time.Date(2026, time.July, 20, 23, 59, 59, 0, location)},
		{"September moves from Sunday 20 to Monday 21", 9, time.Date(2026, time.September, 21, 23, 59, 59, 0, location)},
		{"December moves from Sunday 20 to Monday 21", 12, time.Date(2026, time.December, 21, 23, 59, 59, 0, location)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Deadline{Month: tt.month, Day: 20, AdjustToNextBusinessDay: true}
			if got := d.DateForYear(2026, location); !got.Equal(tt.want) {
				t.Errorf("expected %s, got %s", tt.want, got)
			}
		})
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
