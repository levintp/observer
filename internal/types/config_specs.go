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

// The `ControllerSpec` type is a structure that defines the configuration
// specification of the controller component.
type ControllerSpec struct {
	Host     string `yaml:"host" env:"CONTROLLER_HOST" flag:"controller-host"`                                               // Hostname of the controller.
	Port     int    `yaml:"port" env:"CONTROLLER_PORT" flag:"controller-port" default:"1139"`                                // Connection port of the controller.
	LogFile  string `yaml:"log_file" env:"CONTROLLER_LOG_FILE" flag:"log-file" default:"/var/log/observer/observerctld.log"` // Output file of the controller log.
	LogLevel string `yaml:"log_level" env:"CONTROLLER_LOG_LEVEL" flag:"log-level" default:"info"`                            // Level of the controller log.
}

// The `DatabaseSpec` type is a structure that defines the configuration
// specification of the database component.
type DatabaseSpec struct {
	Host     string `yaml:"host" env:"DATABASE_HOST" flag:"database-host"`                                                // Hostname of the database.
	Port     int    `yaml:"port" env:"DATABASE_PORT" flag:"database-port" default:"9200"`                                 // Connection port of the database.
	User     string `yaml:"user" env:"DATABASE_USER" flag:"database-user"`                                                // Username used to authenticate with the database.
	Pass     string `yaml:"pass" env:"DATABASE_PASS" flag:"database-pass"`                                                // Password used to authenticate with the database.
	Index    string `yaml:"index" env:"DATABASE_INEDX" flag:"database-index" default:"observer-streams"`                  // The elasticsearch index to write metrics to.
	LogFile  string `yaml:"log_file" env:"DATABASE_LOG_FILE" flag:"log-file" default:"/var/log/observer/observerdbd.log"` // Output file of the database log.
	LogLevel string `yaml:"log_level" env:"DATABASE_LOG_LEVEL" flag:"log-level" default:"info"`                           // Level of the database log.
}

// The `AgentSpec` type is a structure that defines the configuration
// specification of the agent component.
type AgentSpec struct {
	Port     int    `yaml:"port" env:"AGENT_PORT" flag:"agent-port" default:"1016"`                                       // Connection port to the agent.
	LogFile  string `yaml:"log_file" env:"DATABASE_AGENT_FILE" flag:"log-file" default:"/var/log/observer/observerd.log"` // Output file of the agent log.
	LogLevel string `yaml:"log_level" env:"DATABASE_AGENT_LEVEL" flag:"log-level" default:"info"`                         // Level of the agent log.
}

// The `APISpec` type is a structure that defines the configuration
// specification of the REST API component.
type APISpec struct {
	Host     string `yaml:"host" env:"API_HOST" flag:"api-host"`                                                           // Hostname of the API.
	Port     int    `yaml:"port" env:"API_PORT" flag:"api-port" default:"1086"`                                            // RPC port of the API.
	RestPort int    `yaml:"rest_port" env:"API_REST_PORT" flag:"api-rest-port" default:"1291"`                             // REST port of the API.
	LogFile  string `yaml:"log_file" env:"DATABASE_API_FILE" flag:"log-file" default:"/var/log/observer/observerapid.log"` // Output file of the API log.
	LogLevel string `yaml:"log_level" env:"DATABASE_API_LEVEL" flag:"log-level" default:"info"`                            // Level of the API log.
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
	Plugin    string                 `yaml:"plugin"`                  // Name of the module plugin.
	Timeout   int                    `yaml:"timeout" default:"1"`     // Maximum execution time.
	Function  string                 `yaml:"function" default:"exec"` // Symbol name of the function to run within the module.
	Arguments map[string]interface{} `yaml:"arguments,omitempty"`     // Arguments to pass to the module.
}

// The `NodeSpec` type is a structure that defines a monitored node in the
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
