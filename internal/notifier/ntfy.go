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
)

// Notification contains all data needed to send a notification
type Notification struct {
	Name        string
	Description string
	Priority    string
	Tags        []string
	DaysUntil   int
	Deadline    time.Time
}

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

// Send sends a notification
func (n *Ntfy) Send(notif Notification) {
	url := fmt.Sprintf("%s/%s", n.config.NtfyServer, n.config.NtfyTopic)

	message := n.buildMessage(notif)

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(message))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}

	n.setHeaders(req, notif)

	resp, err := n.client.Do(req)
	if err != nil {
		log.Printf("Error sending notification: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Printf("✅ Sent: %s (%d days until %s)", notif.Name, notif.DaysUntil, notif.Deadline.Format("02/01"))
	} else {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("❌ Failed to send notification: %s - %s", resp.Status, string(body))
	}
}

func (n *Ntfy) buildMessage(notif Notification) string {
	var urgency string
	switch {
	case notif.DaysUntil == 0:
		urgency = "🚨 HOJE!"
	case notif.DaysUntil == 1:
		urgency = "⚠️ AMANHÃ!"
	default:
		urgency = fmt.Sprintf("📅 Faltam %d dias", notif.DaysUntil)
	}

	return fmt.Sprintf("%s\n%s\nPrazo: %s", urgency, notif.Description, notif.Deadline.Format("02/01/2006"))
}

func (n *Ntfy) setHeaders(req *http.Request, notif Notification) {
	req.Header.Set("Title", notif.Name)
	req.Header.Set("Priority", notif.Priority)
	req.Header.Set("Tags", strings.Join(notif.Tags, ","))

	// Add HTTP Basic Auth if configured
	if n.config.NtfyUser != "" && n.config.NtfyPass != "" {
		req.SetBasicAuth(n.config.NtfyUser, n.config.NtfyPass)
	}

	// Add click action to open relevant portal
	if hasTag(notif.Tags, "seguranca-social") {
		req.Header.Set("Click", "https://app.seg-social.pt")
	} else if hasTag(notif.Tags, "iva") || hasTag(notif.Tags, "irs") {
		req.Header.Set("Click", "https://www.portaldasfinancas.gov.pt")
	}
}

func hasTag(tags []string, tag string) bool {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}
