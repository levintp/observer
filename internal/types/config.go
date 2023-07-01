// The `types` package defines different types, methods and interfaces used to
// represent data and objects throughout the system.
package types

import "gopkg.in/yaml.v3"

// The `Config` type is a structure that contains the configuration of the
// Observer system.
type Config struct {
	Controller ControllerSpec         `yaml:"controller"`                                                                        // Configuration regarding the controller.
	Database   DatabaseSpec           `yaml:"database"`                                                                          // Configuration regarding the database.
	Agent      AgentSpec              `yaml:"agent"`                                                                             // Configuration regarding the local agent.
	API        APISpec                `yaml:"api"`                                                                               // Configuration regarding the API.
	PluginDir  string                 `yaml:"plugin_dir" env:"PLUGIN_DIR" flag:"plugin-dir" default:"/usr/lib/observer/plugins"` // Directory of plugin objects.
	Streams    map[string]*StreamSpec `yaml:"streams"`                                                                           // List of streams.
	Nodes      map[string]*NodeSpec   `yaml:"nodes"`                                                                             // List of nodes.
}

// Function to convert the `Config` type into a printable, human-readable
// string by marshaling it back to the configuration file format.
func (config Config) String() string {
	str, _ := yaml.Marshal(config)
	return string(str)
}
