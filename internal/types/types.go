// The `types` package defines different types, methods and interfaces used to
// represent data and objects throughout the system.
package types

import "gopkg.in/yaml.v3"

// Struct to contain the configuration of the observer system.
type Config struct {
	ConfigFile string `yaml:"config_file" env:"OBSERVER_CONFIG_FILE" default:"/etc/observer/observer.yaml"`
	// Configuration regarding the controller.
	ControllerSpec struct {
		Host     string `yaml:"host" env:"OBSERVER_CONTROLLER_HOST"`                                                      // Hostname of the controller.
		Port     int    `yaml:"port" env:"OBSERVER_CONTROLLER_PORT" default:"1139"`                                       // Connection port of the controller.
		LogFile  string `yaml:"log_file" env:"OBSERVER_CONTROLLER_LOG_FILE" default:"/var/log/observer/observerctld.log"` // Output file of the controller log.
		LogLevel string `yaml:"log_level" env:"OBSERVER_CONTROLLER_LOG_LEVEL" default:"info"`                             // Level of the controller log.
	} `yaml:"controller"`
	// Configuration regarding the database.
	DatabaseSpec struct {
		Host     string `yaml:"host" env:"OBSERVER_DATABASE_HOST"`                                                     // Hostname of the database.
		Port     int    `yaml:"port" env:"OBSERVER_DATABASE_PORT" default:"9200"`                                      // Connection port of the database.
		User     string `yaml:"user" env:"OBSERVER_DATABASE_USER"`                                                     // Username used to authenticate with the database.
		Pass     string `yaml:"pass" env:"OBSERVER_DATABASE_PASS"`                                                     // Password used to authenticate with the database.
		Index    string `yaml:"index" env:"OBSERVER_DATABASE_INEDX" default:"observer-streams"`                        // The elasticsearch index to write metrics to.
		LogFile  string `yaml:"log_file" env:"OBSERVER_DATABASE_LOG_FILE" default:"/var/log/observer/observerdbd.log"` // Output file of the database log.
		LogLevel string `yaml:"log_level" env:"OBSERVER_DATABASE_LOG_LEVEL" default:"info"`                            // Level of the database log.
	} `yaml:"database"`
	// Configuration regarding the local agent.
	AgentSpec struct {
		Port     int    `yaml:"port" env:"OBSERVER_AGENT_PORT" default:"1016"`                                         // Connection port to the agent.
		LogFile  string `yaml:"log_file" env:"OBSERVER_DATABASE_AGENT_FILE" default:"/var/log/observer/observerd.log"` // Output file of the agent log.
		LogLevel string `yaml:"log_level" env:"OBSERVER_DATABASE_AGENT_LEVEL" default:"info"`                          // Level of the agent log.
	} `yaml:"agent"`
	// Configuration regarding the API.
	APISpec struct {
		Host     string `yaml:"host" env:"OBSERVER_API_HOST"`                                                           // Hostname of the API.
		Port     int    `yaml:"port" env:"OBSERVER_API_PORT" default:"1086"`                                            // RPC port of the API.
		RestPort int    `yaml:"rest_port" env:"OBSERVER_API_REST_PORT" default:"1291"`                                  // REST port of the API.
		LogFile  string `yaml:"log_file" env:"OBSERVER_DATABASE_API_FILE" default:"/var/log/observer/observerapid.log"` // Output file of the API log.
		LogLevel string `yaml:"log_level" env:"OBSERVER_DATABASE_API_LEVEL" default:"info"`                             // Level of the API log.
	} `yaml:"api"`
	Streams map[string]*StreamSpec `yaml:"streams"` // List of streams.
	Nodes   map[string]*NodeSpec   `yaml:"nodes"`   // List of nodes.
}

// The `StreamSpec` type is a structure that defines the specification of a
// stream of metrics.
type StreamSpec struct {
	Name       string                 `yaml:"name"`                 // Name of the stream.
	Categories []string               `yaml:"categories,omitempty"` // Categories to run the stream on.
	Metrics    map[string]*MetricSpec `yaml:"metrics"`              // Metrics in the stream.
}

// The `MetricSpec` type is a structure that defines the specification of a
// collectable metric.
type MetricSpec struct {
	Name       string                    `yaml:"name"`                  // Name of the metric.
	Interval   int                       `yaml:"interval" default:"60"` // Collection interval.
	Categories []string                  `yaml:"categories,omitempty"`  // Categories to collect the metric on.
	Module     ModuleSpec                `yaml:"module"`                // Collection module in seconds.
	Thresholds map[string]*ThresholdSpec `yaml:"thresholds,omitempty"`  // Thresholds applied to the metric.
}

// The `ThresholdSpec` type is a structure that defines the specification of a
// threshold on the collected value of a metric.
//
// This threshold represents a barrier to watch for and when its condition is
// met, it will trigger the execution of a respective module.
type ThresholdSpec struct {
	Name       string     `yaml:"name"`                  // Name of the threshold.
	Expr       string     `yaml:"expression"`            // Threshold as a JSONPath expression.
	Interval   int        `yaml:"interval" default:"60"` // Check interval in seconds.
	Categories []string   `yaml:"categories,omitempty"`  // Categories to check the threshold on.
	Module     ModuleSpec `yaml:"module"`                // Module used when expression is met.
}

// The `ModuleSpec` type is a structure that defines the specification of a
// module.
type ModuleSpec struct {
	Name      string                 `yaml:"name"`                // Name of the module.
	Timeout   int                    `yaml:"timeout" default:"1"` // Maximum execution time.
	Arguments map[string]interface{} `yaml:"arguments,omitempty"` // Arguments to pass to the module.
}

// The `NodeSpec` type is a structure that definse a monitored node in the
// cluster.
type NodeSpec struct {
	Name       string   `yaml:"name"`                 // Hostname of the node.
	Categories []string `yaml:"categories,omitempty"` // Categories the node is a part of.
}

// Function to convert the `Config` type into a printable, human-readable
// string by marshaling it back to the configuration file format.
func (config Config) String() string {
	str, _ := yaml.Marshal(config)
	return string(str)
}
