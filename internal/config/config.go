// The `config` package implements configuration parsing into fixed configuration
// and system structs adn types.
package config

import (
	"fmt"

	"github.com/levintp/observer/internal/types"
	log "github.com/sirupsen/logrus"
)

var globalConfiguration *types.Config // Global singleton configuration.

// Function to get the configuration.
func Get() *types.Config {
	if globalConfiguration == nil {
		log.Info("Loading configuration")
		globalConfiguration = &types.Config{}
		if err := buildConfiguration(globalConfiguration); err != nil {
			log.Fatalf("Failed to load configuration: %v", err)
		}
	}

	return globalConfiguration
}

// Function to build a new global configuration.
func buildConfiguration(conf *types.Config) error {

	// Generate minimal default configuration.
	if err := setDefaults(conf); err != nil {
		return fmt.Errorf("default: %v", err)
	}

	// Read the configuration from commandline interface.
	if err := getConfigurationCli(conf); err != nil {
		return fmt.Errorf("commandline: %v", err)
	}

	// Read the configuration from environment.
	if err := getConfigurationEnv(conf); err != nil {
		return fmt.Errorf("environment: %v", err)
	}

	// Read configuration from file.
	err := getConfigurationFile(conf.ConfigFile, conf)
	if err != nil {
		return fmt.Errorf("file: %v", err)
	}

	// Fill empty fields with default values after configuration expansion.
	if err := setDefaults(conf); err != nil {
		return fmt.Errorf("post-process: %v", err)
	}

	// Process configuration.
	updateNames(conf)

	// Validate post-processed configuration.
	if err := validateConfiguration(conf); err != nil {
		return fmt.Errorf("invalid configuration: %v", err)
	}

	return nil
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

// Function to validate the parsed configuration.
func validateConfiguration(conf *types.Config) error {
	return nil
}
