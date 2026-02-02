package config

import (
	"os"
	"testing"
)

func TestLoad_Defaults(t *testing.T) {
	// Clear env vars
	os.Unsetenv("NTFY_TOPIC")
	os.Unsetenv("NTFY_SERVER")

	cfg := Load()

	if cfg.NtfyTopic != "fiscal-reminders" {
		t.Errorf("expected default topic 'fiscal-reminders', got %s", cfg.NtfyTopic)
	}

	if cfg.NtfyServer != "https://ntfy.sh" {
		t.Errorf("expected default server 'https://ntfy.sh', got %s", cfg.NtfyServer)
	}

	if cfg.CheckInterval != 24 {
		t.Errorf("expected default interval 24, got %d", cfg.CheckInterval)
	}

	if len(cfg.DaysBeforeAlert) != 4 {
		t.Errorf("expected 4 alert days, got %d", len(cfg.DaysBeforeAlert))
	}
}

func TestLoad_EnvOverride(t *testing.T) {
	os.Setenv("NTFY_TOPIC", "test-topic")
	os.Setenv("NTFY_SERVER", "https://custom.ntfy.sh")
	defer os.Unsetenv("NTFY_TOPIC")
	defer os.Unsetenv("NTFY_SERVER")

	cfg := Load()

	if cfg.NtfyTopic != "test-topic" {
		t.Errorf("expected env topic 'test-topic', got %s", cfg.NtfyTopic)
	}

	if cfg.NtfyServer != "https://custom.ntfy.sh" {
		t.Errorf("expected env server, got %s", cfg.NtfyServer)
	}
}

func TestGetEnvOrDefault(t *testing.T) {
	os.Setenv("TEST_VAR", "test-value")
	defer os.Unsetenv("TEST_VAR")

	if got := getEnvOrDefault("TEST_VAR", "default"); got != "test-value" {
		t.Errorf("expected 'test-value', got %s", got)
	}

	if got := getEnvOrDefault("NONEXISTENT_VAR", "default"); got != "default" {
		t.Errorf("expected 'default', got %s", got)
	}
}

