package main

import (
	"github.com/levintp/observer/internal/config"
	"github.com/levintp/observer/internal/log"
)

func init() {
	configuration := config.Get()
	if err := log.ConfigureLogger(configuration.AgentSpec.LogFile, configuration.AgentSpec.LogLevel); err != nil {
		log.Fatalf("Failed to configure logger: %v", err)
	}
}

func main() {
	log.Info("Observer Agent daemon started")

	configuration := config.Get()
	log.Infof("Loaded configuration:\n%v", configuration)
}
