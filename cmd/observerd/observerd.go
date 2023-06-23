package main

import (
	"github.com/levintp/observer/internal/config"
	"github.com/levintp/observer/internal/logging"
)

func main() {
	logging.Init()
	configuration := config.Get()
	logging.Logger.Infof("Configuration:\n\n%v\n", configuration)
}
