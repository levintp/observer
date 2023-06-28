// The `config` package implements configuration parsing into fixed configuration
// and system structs adn types.
package config

import (
	"fmt"
	"os"

	"github.com/levintp/observer/internal/log"
	"github.com/levintp/observer/internal/meta"
	"github.com/levintp/observer/internal/types"
	"gopkg.in/yaml.v3"
)

var cfg *types.Config // Global singleton configuration.

const environmentPrefix = "OBSERVER_"
const configurationFileEnv = environmentPrefix + "CONFIG_FILE"
const configurationFileFlag = "config-file"
const defaultConfigurationFile = "/etc/observer/observer.yaml"

// Function to get the configuration.
func Get() *types.Config {
	if cfg == nil {
		log.Info("Loading configuration")
		cfg = &types.Config{}
		if err := buildConfiguration(cfg); err != nil {
			log.Fatalf("Failed to load configuration: %v", err)
		}
	}

	return cfg
}

// Function to build a new global configuration.
func buildConfiguration(cfg *types.Config) error {
	// Generate minimal default configuration.
	if err := meta.SetDefaults(cfg); err != nil {
		return fmt.Errorf("default: %v", err)
	}

	// Read configuration from file.
	err := readConfigurationFile(getConfigurationFile(), cfg)
	if err != nil {
		return fmt.Errorf("file: %v", err)
	}

	// Override configuration with higher priority values from enviroment.
	if err := meta.SetEnvironment(cfg, environmentPrefix); err != nil {
		return fmt.Errorf("environment: %v", err)
	}

	// Override configuration with highest priority values from flags.
	if err := meta.SetFlags(cfg); err != nil {
		return fmt.Errorf("commandline: %v", err)
	}

	// Fill empty fields with default values after configuration expansion.
	if err := meta.SetDefaults(cfg); err != nil {
		return fmt.Errorf("post-process: %v", err)
	}

	// Process configuration.
	updateNames(cfg)

	// Validate post-processed configuration.
	if err := validateConfiguration(cfg); err != nil {
		return fmt.Errorf("invalid configuration: %v", err)
	}

	return nil
}

// Function to get the configuration from the configuration file.
func readConfigurationFile(filename string, cfg *types.Config) error {
	log.Debugf("Reading configuration from %s", filename)

	// Read the configuration file.
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Parse the configuration file.
	err = yaml.Unmarshal(content, cfg)
	if err != nil {
		return err
	}

	return nil
}

// Function to get the configuration filename from either environment or flags, or default.
func getConfigurationFile() string {
	filename := defaultConfigurationFile

	if environFilename := os.Getenv(configurationFileEnv); environFilename != "" {
		filename = environFilename
	}

	if flagFilename := meta.GetFlag(configurationFileFlag); flagFilename != "" {
		filename = flagFilename
	}

	return filename
}

// Function to update names of mapped structures in the configuration.
func updateNames(cfg *types.Config) {
	// Iterate over stream configuration and update mapped names.
	for streamName, stream := range cfg.Streams {
		stream.Name = streamName
		for metricName, metric := range stream.Metrics {
			metric.Name = metricName
			for thresholdName, threshold := range metric.Thresholds {
				threshold.Name = thresholdName
			}
		}
	}

	// Iterate over node configuration an update mapped names.
	for nodeName, node := range cfg.Nodes {
		node.Name = nodeName
	}
}

// Function to validate the parsed configuration.
func validateConfiguration(conf *types.Config) error {
	return nil
}
