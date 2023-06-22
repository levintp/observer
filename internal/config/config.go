// The `config` package implements configuration parsing into fixed configuration
// and system structs adn types.
package config

import (
	"fmt"
	"log"
	"reflect"
	"strconv"

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

	// Generate default configuration.
	err := getDefault(&conf)
	if err != nil {
		log.Fatalf("Failed to generate default configuration: %e", err)
	}

	// Read the configuration from commandline interface.
	// err = getConfigurationCli(&conf)
	// if err != nil {
	// 	log.Fatalf("Failed to read configuration flags: %e", err)
	// }

	// // Read the configuration from environment.
	// err = getConfigurationEnv(&conf)
	// if err != nil {
	// 	log.Fatalf("Failed to read configuration environment: %e", err)
	// }

	// // Read configuration from file.
	// err = getConfigurationFile(conf.ConfigFile, &conf)
	// if err != nil {
	// 	log.Fatalf("Failed to read configuration file: %e", err)
	// }

	// Process configuration.
	updateNames(&conf)

	// Validate post-processed configuration.
	err = validateConfiguration(conf)
	if err != nil {
		log.Fatalf("Invalid configuration: %e", err)
	}

	return &conf
}

// Function to generate a default configuration and populate it into the configuration structure.
func getDefault(conf *types.Config) error {
	*conf = types.Config{}

	return setStructDefaults(conf)
}

// Function to set default values of all fields of a struct, recursively.
func setStructDefaults(ptr any) error {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		return fmt.Errorf("Not a pointer")
	}

	structValue := reflect.ValueOf(ptr).Elem()
	structType := structValue.Type()

	for i := 0; i < structType.NumField(); i++ {
		fieldValue := structValue.Field(i)
		fieldType := structType.Field(i)
		if fieldType.Type.Kind() == reflect.Struct {
			if err := setStructDefaults(fieldValue.Addr().Interface()); err != nil {
				return err
			}
		}
		if defaultVal := fieldType.Tag.Get("default"); defaultVal != "-" {
			if err := setField(fieldValue, defaultVal); err != nil {
				return err
			}

		}
	}
	return nil
}

func setField(field reflect.Value, defaultVal string) error {

	if !field.CanSet() {
		return fmt.Errorf("Can't set value\n")
	}

	switch field.Kind() {

	case reflect.Int:
		if val, err := strconv.ParseInt(defaultVal, 10, 64); err == nil {
			field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
		}
	case reflect.String:
		field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
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

func validateConfiguration(conf types.Config) error {
	return nil
}
