package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Config holds the application configuration
type Config struct {
	NtfyTopic       string `json:"ntfy_topic"`
	NtfyServer      string `json:"ntfy_server"`
	NtfyUser        string `json:"ntfy_user"`
	NtfyPass        string `json:"ntfy_pass"`
	CheckInterval   int    `json:"check_interval_hours"`
	DaysBeforeAlert []int  `json:"days_before_alert"`
}

// Deadline represents a fiscal deadline
type Deadline struct {
	Name        string
	Description string
	Month       int
	Day         int
	Priority    string // "urgent", "high", "default"
	Tags        []string
}

var deadlines = []Deadline{
	// Declaração Trimestral Segurança Social
	{Name: "📋 Declaração Trimestral SegSoc", Description: "Declarar rendimentos Out-Dez à Segurança Social", Month: 1, Day: 31, Priority: "high", Tags: []string{"seguranca-social", "trimestral"}},
	{Name: "📋 Declaração Trimestral SegSoc", Description: "Declarar rendimentos Jan-Mar à Segurança Social", Month: 4, Day: 30, Priority: "high", Tags: []string{"seguranca-social", "trimestral"}},
	{Name: "📋 Declaração Trimestral SegSoc", Description: "Declarar rendimentos Abr-Jun à Segurança Social", Month: 7, Day: 31, Priority: "high", Tags: []string{"seguranca-social", "trimestral"}},
	{Name: "📋 Declaração Trimestral SegSoc", Description: "Declarar rendimentos Jul-Set à Segurança Social", Month: 10, Day: 31, Priority: "high", Tags: []string{"seguranca-social", "trimestral"}},

	// IVA Trimestral
	{Name: "💶 Declaração IVA Trimestral", Description: "Entregar declaração IVA do 1º trimestre", Month: 5, Day: 20, Priority: "high", Tags: []string{"iva", "trimestral"}},
	{Name: "💶 Declaração IVA Trimestral", Description: "Entregar declaração IVA do 2º trimestre", Month: 8, Day: 20, Priority: "high", Tags: []string{"iva", "trimestral"}},
	{Name: "💶 Declaração IVA Trimestral", Description: "Entregar declaração IVA do 3º trimestre", Month: 11, Day: 20, Priority: "high", Tags: []string{"iva", "trimestral"}},
	{Name: "💶 Declaração IVA Trimestral", Description: "Entregar declaração IVA do 4º trimestre", Month: 2, Day: 20, Priority: "high", Tags: []string{"iva", "trimestral"}},

	// Pagamento contribuições SegSoc (mensal)
	{Name: "💳 Pagamento SegSoc", Description: "Pagar contribuições Segurança Social", Month: 0, Day: 20, Priority: "default", Tags: []string{"seguranca-social", "pagamento"}}, // Month 0 = every month

	// IRS Anual
	{Name: "📝 IRS Anual - Início", Description: "Período de entrega do IRS começa", Month: 4, Day: 1, Priority: "default", Tags: []string{"irs", "anual"}},
	{Name: "📝 IRS Anual - Fim", Description: "Último dia para entregar IRS", Month: 6, Day: 30, Priority: "urgent", Tags: []string{"irs", "anual"}},
}

func main() {
	log.Println("🔔 Fiscal Reminders - Starting...")

	config := loadConfig()
	log.Printf("Config: topic=%s, server=%s, check every %dh", config.NtfyTopic, config.NtfyServer, config.CheckInterval)

	// Run immediately on start
	checkDeadlines(config)

	// Then run periodically
	ticker := time.NewTicker(time.Duration(config.CheckInterval) * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		checkDeadlines(config)
	}
}

func loadConfig() Config {
	config := Config{
		NtfyTopic:       os.Getenv("NTFY_TOPIC"),
		NtfyServer:      os.Getenv("NTFY_SERVER"),
		CheckInterval:   24,                // Default: check every 24 hours
		DaysBeforeAlert: []int{7, 3, 1, 0}, // Alert 7, 3, 1, and 0 days before
	}

	if config.NtfyTopic == "" {
		config.NtfyTopic = "fiscal-reminders"
	}
	if config.NtfyServer == "" {
		config.NtfyServer = "https://ntfy.sh"
	}

	// Try to load from config file
	if data, err := os.ReadFile("config.json"); err == nil {
		json.Unmarshal(data, &config)
	}

	return config
}

func checkDeadlines(config Config) {
	now := time.Now()
	year := now.Year()

	log.Printf("Checking deadlines for %s", now.Format("2006-01-02"))

	for _, deadline := range deadlines {
		// Handle monthly deadlines (Month = 0)
		if deadline.Month == 0 {
			checkMonthlyDeadline(config, deadline, now)
			continue
		}

		// Calculate deadline date for current year
		deadlineDate := time.Date(year, time.Month(deadline.Month), deadline.Day, 23, 59, 59, 0, now.Location())

		// If deadline already passed this year, check next year
		if deadlineDate.Before(now) {
			deadlineDate = time.Date(year+1, time.Month(deadline.Month), deadline.Day, 23, 59, 59, 0, now.Location())
		}

		daysUntil := int(deadlineDate.Sub(now).Hours() / 24)

		for _, alertDay := range config.DaysBeforeAlert {
			if daysUntil == alertDay {
				sendNotification(config, deadline, daysUntil, deadlineDate)
				break
			}
		}
	}
}

func checkMonthlyDeadline(config Config, deadline Deadline, now time.Time) {
	// Monthly deadline - check for current month
	deadlineDate := time.Date(now.Year(), now.Month(), deadline.Day, 23, 59, 59, 0, now.Location())

	// If already passed this month, check next month
	if deadlineDate.Before(now) {
		deadlineDate = deadlineDate.AddDate(0, 1, 0)
	}

	daysUntil := int(deadlineDate.Sub(now).Hours() / 24)

	for _, alertDay := range config.DaysBeforeAlert {
		if daysUntil == alertDay {
			sendNotification(config, deadline, daysUntil, deadlineDate)
			break
		}
	}
}

func sendNotification(config Config, deadline Deadline, daysUntil int, deadlineDate time.Time) {
	url := fmt.Sprintf("%s/%s", config.NtfyServer, config.NtfyTopic)

	// Build message
	var urgency string
	switch {
	case daysUntil == 0:
		urgency = "🚨 HOJE!"
	case daysUntil == 1:
		urgency = "⚠️ AMANHÃ!"
	default:
		urgency = fmt.Sprintf("📅 Faltam %d dias", daysUntil)
	}

	message := fmt.Sprintf("%s\n%s\nPrazo: %s", urgency, deadline.Description, deadlineDate.Format("02/01/2006"))

	// Create request
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(message))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}

	req.Header.Set("Title", deadline.Name)
	req.Header.Set("Priority", deadline.Priority)
	req.Header.Set("Tags", joinTags(deadline.Tags))

	// Add HTTP Basic Auth if configured
	if config.NtfyUser != "" && config.NtfyPass != "" {
		req.SetBasicAuth(config.NtfyUser, config.NtfyPass)
	}

	// Add click action to open relevant portal
	if contains(deadline.Tags, "seguranca-social") {
		req.Header.Set("Click", "https://app.seg-social.pt")
	} else if contains(deadline.Tags, "iva") || contains(deadline.Tags, "irs") {
		req.Header.Set("Click", "https://www.portaldasfinancas.gov.pt")
	}

	// Send notification
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending notification: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Printf("✅ Sent: %s (%d days until %s)", deadline.Name, daysUntil, deadlineDate.Format("02/01"))
	} else {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("❌ Failed to send notification: %s - %s", resp.Status, string(body))
	}
}

func joinTags(tags []string) string {
	result := ""
	for i, tag := range tags {
		if i > 0 {
			result += ","
		}
		result += tag
	}
	return result
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
