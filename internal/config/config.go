// The `config` package implements configuration parsing into fixed configuration
// and system structs adn types.
package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/levintp/observer/internal/types"
)

var globalConfiguration *types.Config // Global singleton configuration.

// Function to get the configuration.
func Get() *types.Config {
	if globalConfiguration == nil {
		log.Println("Building new configuration")
		globalConfiguration = buildConfiguration()
	}

	return globalConfiguration
}

// Function to build a new global configuration.
func buildConfiguration() *types.Config {
	var c types.Config
	err := cleanenv.ReadConfig("/etc/observer/observer.yaml", &c)
	if err != nil {
		log.Fatalf("Failed to parse configuration: %e", err)
	}

	// Update mapped structures with keys as names.
	for streamName, stream := range c.Streams {
		stream.Name = streamName
		for metricName, metric := range stream.Metrics {
			metric.Name = metricName
			if metric.Thresholds != nil {
				for thresholdName, threshold := range metric.Thresholds {
					threshold.Name = thresholdName
				}
			}
		}
	}
	for nodeName, node := range c.Nodes {
		node.Name = nodeName
	}

	return &c
}
