// The `config` package implements configuration parsing into fixed configuration
// and system structs adn types.
package config

import (
	"log"

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

	// Generate minimal default configuration.
	log.Println("Generating default minimal configuration")
	err := setDefaults(&conf)
	if err != nil {
		log.Fatalf("Failed to generate default configuration: %e", err)
	}

	// Read the configuration from commandline interface.
	log.Println("Reading configuration from the commandline interface")
	err = getConfigurationCli(&conf)
	if err != nil {
		log.Fatalf("Failed to read configuration flags: %e", err)
	}

	// Read the configuration from environment.
	log.Println("Reading configuration from the environment")
	err = getConfigurationEnv(&conf)
	if err != nil {
		log.Fatalf("Failed to read configuration environment: %e", err)
	}

	// Read configuration from file.
	log.Println("Reading configuration from the configuration file")
	err = getConfigurationFile(conf.ConfigFile, &conf)
	if err != nil {
		log.Fatalf("Failed to read configuration file: %e", err)
	}

	// Fill empty fields with default values after configuration expansion.
	log.Println("Filling empty configuration fields with default values")
	err = setDefaults(&conf)
	if err != nil {
		log.Fatalf("Failed to fill defaults in configuration: %e", err)
	}

	// Process configuration.
	updateNames(&conf)

	// Validate post-processed configuration.
	err = validateConfiguration(conf)
	if err != nil {
		log.Fatalf("Invalid configuration: %e", err)
	}

	return &conf
}

// Function to update names of mapped structures in the configuration.
func updateNames(conf *types.Config) {
	// Iterate over stream configuration and update mapped names.
	for streamName, stream := range conf.Streams {
		stream.Name = streamName
		for metricName, metric := range stream.Metrics {
			metric.Name = metricName
			for thresholdName, threshold := range metric.Thresholds {
				threshold.Name = thresholdName
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
