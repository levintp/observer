// The `types` package defines different types, methods and interfaces used to
// represent data and objects throughout the system.
package types

// The `Metric` type is a structure that defines a collectable metric.
type Metric struct {
	Name       string      `toml:"name"`                   // Name of the metric.
	Interval   string      `toml:"interval" default:"60s"` // Collection interval.
	Timeout    string      `toml:"timeout" default:"59s"`  // Maximum collection time.
	Categories []string    `toml:"categories"`             // Categories to collect the metric on.
	Module     string      `toml:"module"`                 // Collection module.
	Args       ModuleArgs  `toml:"arguments"`              // Arguments to the module.
	Thresholds []Threshold `toml:"threshold"`              // Thresholds applied to the metric.
}

// The `Threshold` type is a structure that defines a threshold on the
// collected value of a metric.
//
// This threshold represents a barrier to watch for and when its condition is
// met, it will trigger the execution of a response module.
type Threshold struct {
	Expr   string     `toml:"expression"` // Threshold as a JSONPath expression.
	Module string     `toml:"module"`     // Module used when expression is met.
	Args   ModuleArgs `toml:"arguments"`  // Arguments to the module.
}

// The `Node` type is a structure that definse a monitored node in the
// cluster.
type Node struct {
	Name       string   `toml:"name"`       // Hostname of the node.
	Categories []string `toml:"categories"` // Categories the node is a part of.
}

// The `ModuleArgs` types is a list of arguments passed to a module when
// executed.
type ModuleArgs []any
