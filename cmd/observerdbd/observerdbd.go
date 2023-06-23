package main

import (
	"github.com/levintp/observer/internal/common"
	"github.com/levintp/observer/internal/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	common.InitLogger()
	configuration := config.Get()
	if err := common.ConfigureLogger(configuration.DatabaseSpec.LogFile, configuration.DatabaseSpec.LogLevel); err != nil {
		log.Fatalf("Failed to configure logger: %v", err)
	}
}

func main() {
	log.Info("Observer Database daemon started")

	configuration := config.Get()
	log.Infof("Loaded configuration:\n%v", configuration)
}