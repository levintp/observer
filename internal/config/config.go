// The `config` package implements configuration parsing into fixed configuration
// and system structs adn types.
package config

import (
	"log"
	"os"

	"github.com/pelletier/go-toml"

	"github.com/levintp/observer/internal/types"
)

// Struct to contain the configuration of the observer system.
type Config struct {
	ControllerInfo controllerInfo `toml:"controller"` // Information about the controller.
	DatabaseInfo   databaseInfo   `toml:"database"`   // Information about the database.
	AgentInfo      agentInfo      `toml:"agent"`      // Information about the agent.
	ApiInfo        apiInfo        `toml:"api"`        // Information about the API.
	Metrics        []types.Metric `toml:"metric"`     // List of metrics.
	Nodes          []types.Node   `toml:"node"`       // List of nodes.
}

type controllerInfo struct {
	Host string `toml:"host" default:"localhost"` // Hostname of the controller.
	Port int    `toml:"port" default:"6710"`      // Connection port of the controller.
}

type databaseInfo struct {
	Host string `toml:"host" default:"localhost"` // Hostname of the database.
	Port int    `toml:"port" default:"9200"`      // Connection port of the database.
	User string `toml:"user"`                     // Username used to authenticate with the database.
	Pass string `toml:"pass"`                     // Password used to authenticate with the database.
}

type agentInfo struct {
	Port int `toml:"port" default:"6711"` // Connection port to the agent
}

type apiInfo struct {
	Host string `toml:"host" default:"localhost"` // Hostname of the REST API.
	Port int    `toml:"port" default:"6712"`      // Connection port of the REST API.
}

func Get(configPath string) Config {
	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := toml.Unmarshal(content, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
