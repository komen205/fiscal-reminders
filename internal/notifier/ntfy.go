package notifier

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/komen205/fiscal-reminders/internal/config"
	"github.com/komen205/fiscal-reminders/internal/deadline"
)

// Ntfy handles sending notifications via ntfy.sh
type Ntfy struct {
	config *config.Config
	client *http.Client
}

// New creates a new Ntfy notifier
func New(cfg *config.Config) *Ntfy {
	return &Ntfy{
		config: cfg,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

// Send sends a notification for a deadline
func (n *Ntfy) Send(d deadline.Deadline, daysUntil int, deadlineDate time.Time) {
	url := fmt.Sprintf("%s/%s", n.config.NtfyServer, n.config.NtfyTopic)

	message := n.buildMessage(d, daysUntil, deadlineDate)

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(message))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}

	n.setHeaders(req, d)

	resp, err := n.client.Do(req)
	if err != nil {
		log.Printf("Error sending notification: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Printf("✅ Sent: %s (%d days until %s)", d.Name, daysUntil, deadlineDate.Format("02/01"))
	} else {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("❌ Failed to send notification: %s - %s", resp.Status, string(body))
	}
}

func (n *Ntfy) buildMessage(d deadline.Deadline, daysUntil int, deadlineDate time.Time) string {
	var urgency string
	switch {
	case daysUntil == 0:
		urgency = "🚨 HOJE!"
	case daysUntil == 1:
		urgency = "⚠️ AMANHÃ!"
	default:
		urgency = fmt.Sprintf("📅 Faltam %d dias", daysUntil)
	}

	return fmt.Sprintf("%s\n%s\nPrazo: %s", urgency, d.Description, deadlineDate.Format("02/01/2006"))
}

func (n *Ntfy) setHeaders(req *http.Request, d deadline.Deadline) {
	req.Header.Set("Title", d.Name)
	req.Header.Set("Priority", d.Priority)
	req.Header.Set("Tags", strings.Join(d.Tags, ","))

	// Add HTTP Basic Auth if configured
	if n.config.NtfyUser != "" && n.config.NtfyPass != "" {
		req.SetBasicAuth(n.config.NtfyUser, n.config.NtfyPass)
	}

	// Add click action to open relevant portal
	if d.HasTag("seguranca-social") {
		req.Header.Set("Click", "https://app.seg-social.pt")
	} else if d.HasTag("iva") || d.HasTag("irs") {
		req.Header.Set("Click", "https://www.portaldasfinancas.gov.pt")
	}
}

