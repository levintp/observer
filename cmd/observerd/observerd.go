package main

import (
	"time"

	"github.com/levintp/observer/internal/config"
	"github.com/levintp/observer/internal/data_stream"
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

	log.Info("Building streams")
	streams := make([]data_stream.DataStream, 0)
	for _, streamSpec := range config.Get().Streams {
		stream := data_stream.New(streamSpec)
		if stream.ShouldRun() {
			log.Debugw("Registering stream", "stream", streamSpec.Name)
			streams = append(streams, stream)
		} else {
			log.Debugw("Skipping stream", "stream", streamSpec.Name)
		}
	}

	log.Info("Starting all registered streams")
	for _, stream := range streams {
		stream.Start()
	}

	time.Sleep(1000000)
}
