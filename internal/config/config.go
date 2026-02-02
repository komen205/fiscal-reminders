package config

import (
	"encoding/json"
	"os"
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

// Load reads configuration from environment variables and config file
func Load() *Config {
	cfg := &Config{
		NtfyTopic:       getEnvOrDefault("NTFY_TOPIC", "fiscal-reminders"),
		NtfyServer:      getEnvOrDefault("NTFY_SERVER", "https://ntfy.sh"),
		NtfyUser:        os.Getenv("NTFY_USER"),
		NtfyPass:        os.Getenv("NTFY_PASS"),
		CheckInterval:   24,
		DaysBeforeAlert: []int{7, 3, 1, 0},
	}

	// Try to load from config file (overrides defaults)
	cfg.loadFromFile("config.json")

	return cfg
}

func (c *Config) loadFromFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		return // File not found is OK, use defaults
	}
	json.Unmarshal(data, c)
}

func getEnvOrDefault(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

