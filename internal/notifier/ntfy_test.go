package notifier

import (
	"strings"
	"testing"
	"time"

	"github.com/komen205/fiscal-reminders/internal/config"
)

func TestNtfy_buildMessage_Today(t *testing.T) {
	n := New(&config.Config{})
	notif := Notification{
		Description: "Test deadline",
		DaysUntil:   0,
		Deadline:    time.Now(),
	}

	msg := n.buildMessage(notif)

	if !strings.Contains(msg, "HOJE") {
		t.Error("message should contain 'HOJE' for 0 days")
	}
}

func TestNtfy_buildMessage_Tomorrow(t *testing.T) {
	n := New(&config.Config{})
	notif := Notification{
		Description: "Test deadline",
		DaysUntil:   1,
		Deadline:    time.Now().AddDate(0, 0, 1),
	}

	msg := n.buildMessage(notif)

	if !strings.Contains(msg, "AMANHÃ") {
		t.Error("message should contain 'AMANHÃ' for 1 day")
	}
}

func TestNtfy_buildMessage_Days(t *testing.T) {
	n := New(&config.Config{})
	notif := Notification{
		Description: "Test deadline",
		DaysUntil:   7,
		Deadline:    time.Now().AddDate(0, 0, 7),
	}

	msg := n.buildMessage(notif)

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

func TestHasTag(t *testing.T) {
	tags := []string{"iva", "trimestral"}

	if !hasTag(tags, "iva") {
		t.Error("should find 'iva' tag")
	}

	if !hasTag(tags, "trimestral") {
		t.Error("should find 'trimestral' tag")
	}

	if hasTag(tags, "irs") {
		t.Error("should not find 'irs' tag")
	}

	if hasTag(nil, "anything") {
		t.Error("nil tags should return false")
	}

	if hasTag([]string{}, "anything") {
		t.Error("empty tags should return false")
	}
}

func TestNotification_Fields(t *testing.T) {
	notif := Notification{
		Name:        "Test",
		Description: "Test description",
		Priority:    "high",
		Tags:        []string{"test"},
		DaysUntil:   5,
		Deadline:    time.Now().AddDate(0, 0, 5),
	}

	if notif.Name != "Test" {
		t.Error("name not set")
	}

	if notif.Priority != "high" {
		t.Error("priority not set")
	}

	if len(notif.Tags) != 1 {
		t.Error("tags not set")
	}
}
