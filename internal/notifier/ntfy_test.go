package notifier

import (
	"strings"
	"testing"
	"time"

	"github.com/komen205/fiscal-reminders/internal/config"
	"github.com/komen205/fiscal-reminders/internal/deadline"
)

func TestNtfy_buildMessage_Today(t *testing.T) {
	n := New(&config.Config{})
	d := deadline.Deadline{Description: "Test deadline"}
	deadlineDate := time.Now()

	msg := n.buildMessage(d, 0, deadlineDate)

	if !strings.Contains(msg, "HOJE") {
		t.Error("message should contain 'HOJE' for 0 days")
	}
}

func TestNtfy_buildMessage_Tomorrow(t *testing.T) {
	n := New(&config.Config{})
	d := deadline.Deadline{Description: "Test deadline"}
	deadlineDate := time.Now().AddDate(0, 0, 1)

	msg := n.buildMessage(d, 1, deadlineDate)

	if !strings.Contains(msg, "AMANHÃ") {
		t.Error("message should contain 'AMANHÃ' for 1 day")
	}
}

func TestNtfy_buildMessage_Days(t *testing.T) {
	n := New(&config.Config{})
	d := deadline.Deadline{Description: "Test deadline"}
	deadlineDate := time.Now().AddDate(0, 0, 7)

	msg := n.buildMessage(d, 7, deadlineDate)

	if !strings.Contains(msg, "7 dias") {
		t.Error("message should contain '7 dias'")
	}
}

func TestNew(t *testing.T) {
	cfg := &config.Config{
		NtfyServer: "https://ntfy.sh",
		NtfyTopic:  "test",
	}

	n := New(cfg)

	if n.config != cfg {
		t.Error("config not set")
	}

	if n.client == nil {
		t.Error("http client not initialized")
	}
}

