package deadline

import (
	"log"
	"time"

	"github.com/komen205/fiscal-reminders/internal/config"
	"github.com/komen205/fiscal-reminders/internal/notifier"
)

// Checker handles deadline checking logic
type Checker struct {
	config   *config.Config
	notifier *notifier.Ntfy
}

// NewChecker creates a new deadline checker
func NewChecker(cfg *config.Config) *Checker {
	return &Checker{
		config:   cfg,
		notifier: notifier.New(cfg),
	}
}

// CheckAll checks all deadlines and sends notifications as needed
func (c *Checker) CheckAll() {
	now := time.Now()
	log.Printf("Checking deadlines for %s", now.Format("2006-01-02"))

	for _, d := range All {
		if d.IsMonthly() {
			c.checkMonthly(d, now)
		} else {
			c.checkYearly(d, now)
		}
	}
}

func (c *Checker) checkYearly(d Deadline, now time.Time) {
	year := now.Year()

	// Calculate deadline date for current year
	deadlineDate := time.Date(year, time.Month(d.Month), d.Day, 23, 59, 59, 0, now.Location())

	// If deadline already passed this year, check next year
	if deadlineDate.Before(now) {
		deadlineDate = time.Date(year+1, time.Month(d.Month), d.Day, 23, 59, 59, 0, now.Location())
	}

	c.checkAndNotify(d, now, deadlineDate)
}

func (c *Checker) checkMonthly(d Deadline, now time.Time) {
	// Monthly deadline - check for current month
	deadlineDate := time.Date(now.Year(), now.Month(), d.Day, 23, 59, 59, 0, now.Location())

	// If already passed this month, check next month
	if deadlineDate.Before(now) {
		deadlineDate = deadlineDate.AddDate(0, 1, 0)
	}

	c.checkAndNotify(d, now, deadlineDate)
}

func (c *Checker) checkAndNotify(d Deadline, now time.Time, deadlineDate time.Time) {
	daysUntil := int(deadlineDate.Sub(now).Hours() / 24)

	for _, alertDay := range c.config.DaysBeforeAlert {
		if daysUntil == alertDay {
			c.notifier.Send(d, daysUntil, deadlineDate)
			break
		}
	}
}

// DaysUntil calculates days until a given deadline
func DaysUntil(deadline time.Time) int {
	return int(time.Until(deadline).Hours() / 24)
}

