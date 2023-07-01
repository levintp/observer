// The `config` package implements configuration parsing into fixed configuration
// and system structs adn types.
package config

import (
	"fmt"
	"os"

	"github.com/levintp/observer/internal/types"
	"gopkg.in/yaml.v3"
)

var cfg *Config // Global singleton configuration.

// The `Config` type is a structure that contains the configuration of the
// Observer system.
type Config struct {
	Controller types.ControllerSpec         `yaml:"controller"`                                                                        // Configuration regarding the controller.
	Database   types.DatabaseSpec           `yaml:"database"`                                                                          // Configuration regarding the database.
	Agent      types.AgentSpec              `yaml:"agent"`                                                                             // Configuration regarding the local agent.
	API        types.APISpec                `yaml:"api"`                                                                               // Configuration regarding the API.
	PluginDir  string                       `yaml:"plugin_dir" env:"PLUGIN_DIR" flag:"plugin-dir" default:"/usr/lib/observer/plugins"` // Directory of plugin objects.
	Streams    map[string]*types.StreamSpec `yaml:"streams"`                                                                           // List of streams.
	Nodes      map[string]*types.NodeSpec   `yaml:"nodes"`                                                                             // List of nodes.
}

// Function to convert the `Config` type into a printable, human-readable
// string by marshaling it back to the configuration file format.
func (cfg Config) String() string {
	str, _ := yaml.Marshal(cfg)
	return string(str)
}

// Function to get the specification of the local node.
func (cfg Config) GetLocalNode() (*types.NodeSpec, error) {
	// Get hostname of local node.
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	// Iterate over the nodes in the configuration.
	node, found := cfg.Nodes[hostname]
	if !found {
		return nil, fmt.Errorf("node %s not in configuration", hostname)
	}

	return node, nil
}
