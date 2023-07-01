package main

import (
	"github.com/levintp/observer/internal/config"
	"github.com/levintp/observer/internal/log"
)

func init() {
	log.Info("Observer Agent daemon initializing")

	cfg := config.Get()

	if err := log.ConfigureLogger(cfg.Agent.LogFile, cfg.Agent.LogLevel); err != nil {
		log.Fatalf("Failed to configure logger: %v", err)
	}
}

func main() {
	log.Info("Observer Agent daemon started")

	for _, streamSpec := range config.Get().Streams {
		log.Infow("Registering stream", "stream", streamSpec.Name)
		for _, metricSpec := range streamSpec.Metrics {
			log.Infow("Registering metric",
				"stream", streamSpec.Name,
				"metric", metricSpec.Name)
		}
	}
}
