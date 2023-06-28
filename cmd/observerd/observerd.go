package main

import (
	"github.com/levintp/observer/internal/config"
	"github.com/levintp/observer/internal/log"
)

func init() {
	log.Info("Observer Agent daemon initializing")

	cfg := config.Get()

	if err := log.ConfigureLogger(cfg.AgentSpec.LogFile, cfg.AgentSpec.LogLevel); err != nil {
		log.Fatalf("Failed to configure logger: %v", err)
	}
}

func main() {
	log.Info("Observer Agent daemon started")
}
