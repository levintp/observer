package main

import (
	"github.com/levintp/observer/internal/config"
	"github.com/levintp/observer/internal/log"
)

func init() {
	configuration := config.Get()
	if err := log.ConfigureLogger(configuration.ControllerSpec.LogFile, configuration.ControllerSpec.LogLevel); err != nil {
		log.Fatalf("Failed to configure logger: %v", err)
	}
}

func main() {
	log.Info("Observer Controller daemon started")

	configuration := config.Get()
	log.Infof("Loaded configuration:\n%v", configuration)
}
