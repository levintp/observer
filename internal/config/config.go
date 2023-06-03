// The `config` package implements configuration parsing into fixed configuration
// and system structs adn types.
package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"

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

// Function to read the configuration from the configuration file.
func readFile(configPath string) types.Config {
	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read configuration file: %e", err)
	}

	var config types.Config
	if err := yaml.Unmarshal(content, &config); err != nil {
		log.Fatalf("Failed to parse configuration file: %e", err)
	}

	return config
}

// Function to build a new global configuration.
func buildConfiguration() *types.Config {
	var c types.Config
	err := cleanenv.ReadConfig("/etc/observer/observer.yaml", &c)
	if err != nil {
		log.Fatalf("Failed to parse configuration: %e", err)
	}

	// c = readFile("/etc/observer/observer.yaml")
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
	return &c
}
