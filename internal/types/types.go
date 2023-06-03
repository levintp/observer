// The `types` package defines different types, methods and interfaces used to
// represent data and objects throughout the system.
package types

import "gopkg.in/yaml.v3"

// Struct to contain the configuration of the observer system.
type Config struct {
	// Configuration regarding the controller.
	ControllerSpec struct {
		Host string `yaml:"host" env:"OBSERVER_CONTROLLER_HOST"`                    // Hostname of the controller.
		Port int    `yaml:"port" env:"OBSERVER_CONTROLLER_PORT" env-default:"1139"` // Connection port of the controller.
	} `yaml:"controller"`
	// Configuration regarding the database.
	DatabaseSpec struct {
		Host  string `yaml:"host" env:"OBSERVER_DATABASE_HOST"`                                  // Hostname of the database.
		Port  int    `yaml:"port" env:"OBSERVER_DATABASE_PORT" env-default:"9200"`               // Connection port of the database.
		User  string `yaml:"user" env:"OBSERVER_DATABASE_USER" env-required:""`                  // Username used to authenticate with the database.
		Pass  string `yaml:"pass" env:"OBSERVER_DATABASE_PASS" env-required:""`                  // Password used to authenticate with the database.
		Index string `yaml:"inedx" env:"OBSERVER_DATABASE_INEDX" env-default:"observer-streams"` // The elasticsearch index to write metrics to.
	} `yaml:"database"`
	// Configuration regarding the local agent.
	AgentSpec struct {
		Port int `yaml:"port" env:"OBSERVER_AGENT_PORT" env-default:"1016"` // Connection port to the agent.
	} `yaml:"agent"`
	// Configuration regarding the API.
	APISpec struct {
		Host string `yaml:"host" env:"OBSERVER_API_HOST" env-required:""`    // Hostname of the API.
		Port int    `yaml:"port" env:"OBSERVER_API_PORT" env-default:"1086"` // Connection port of the API.
	} `yaml:"api"`
	Streams map[string]*StreamSpec `yaml:"streams"` // List of streams.
	Nodes   map[string]*NodeSpec   `yaml:"nodes"`   // List of nodes.
}

// The `StreamSpec` type is a structure that defines the specification of a
// stream of metrics.
type StreamSpec struct {
	Name    string                 `yaml:"name"`                    // Name of the stream.
	Metrics map[string]*MetricSpec `yaml:"metrics" env-required:""` // Metrics in the stream.
}

// The `MetricSpec` type is a structure that defines the specification of a
// collectable metric.
type MetricSpec struct {
	Name       string                    `yaml:"name"`                      // Name of the metric.
	Interval   int                       `yaml:"interval" env-default:"60"` // Collection interval.
	Categories []string                  `yaml:"categories,omitempty"`      // Categories to collect the metric on.
	Module     ModuleSpec                `yaml:"module" env-required:""`    // Collection module in seconds.
	Thresholds map[string]*ThresholdSpec `yaml:"thresholds,omitempty"`      // Thresholds applied to the metric.
}

// The `ThresholdSpec` type is a structure that defines the specification of a
// threshold on the collected value of a metric.
//
// This threshold represents a barrier to watch for and when its condition is
// met, it will trigger the execution of a respective module.
type ThresholdSpec struct {
	Name     string     `yaml:"name"`                      // Name of the threshold.
	Expr     string     `yaml:"expression"`                // Threshold as a JSONPath expression.
	Interval int        `yaml:"interval" env-default:"60"` // Check interval in seconds.
	Module   ModuleSpec `yaml:"module"`                    // Module used when expression is met.
}

// The `ModuleSpec` type is a structure that defines the specification of a
// module.
type ModuleSpec struct {
	Name      string                 `yaml:"name"`                     // Name of the module.
	Timeout   int                    `yaml:"timeout" env-default:"59"` // Maximum execution time.
	Arguments map[string]interface{} `yaml:"arguments,omitempty"`      // Arguments to pass to the module.
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
