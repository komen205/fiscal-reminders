package main

import (
	"log"
	"time"

	"github.com/komen205/fiscal-reminders/internal/config"
	"github.com/komen205/fiscal-reminders/internal/deadline"
)

func main() {
	log.Println("🔔 Fiscal Reminders - Starting...")

	cfg := config.Load()
	log.Printf("Config: topic=%s, server=%s, check every %dh", cfg.NtfyTopic, cfg.NtfyServer, cfg.CheckInterval)

	checker := deadline.NewChecker(cfg)

	// Run immediately on start
	checker.CheckAll()

	// Then run periodically
	ticker := time.NewTicker(time.Duration(cfg.CheckInterval) * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		checker.CheckAll()
	}
}

