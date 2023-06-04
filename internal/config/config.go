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
	var conf types.Config

	// Read configuration from disk.
	err := cleanenv.ReadConfig("/etc/observer/observer.yaml", &conf)
	if err != nil {
		log.Fatalf("Failed to parse configuration: %e", err)
	}

	// Process configuration.
	updateConfigurationNames(&conf)

	// Validate post-processed configuration.
	err = validateConfiguration(conf)
	if err != nil {
		log.Fatalf("Invalid configuration: %e", err)
	}

	return &conf
}

// Function to update names of mapped structures in the configuration.
func updateConfigurationNames(conf *types.Config) {
	// Iterate over stream configuration and update mapped names.
	for streamName, stream := range conf.Streams {
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

	// Iterate over node configuration an update mapped names.
	for nodeName, node := range conf.Nodes {
		node.Name = nodeName
	}
}

func validateConfiguration(conf types.Config) error {
	return nil
}
