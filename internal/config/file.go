package config

import (
	"os"

	"github.com/levintp/observer/internal/types"
	"gopkg.in/yaml.v3"
)

// Function to get the configuration from the configuration file.
func getConfigurationFile(filename string, conf *types.Config) error {
	// Read the configuration file.
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Parse the configuration file.
	if err := yaml.Unmarshal(content, conf); err != nil {
		return err
	}

	return nil
}
