package types

// The `ThresholdSpec` type is a structure that defines the specification of a
// threshold on the collected value of a metric.
//
// This threshold represents a barrier to watch for and when its condition is
// met, it will trigger the execution of a respective module.
type ThresholdSpec struct {
	Name       string     `yaml:"name"`                  // Name of the threshold.
	Expr       string     `yaml:"expression"`            // Threshold as a JSONPath expression.
	Interval   int        `yaml:"interval" default:"60"` // Check interval in seconds.
	Categories []Category `yaml:"categories,omitempty"`  // Categories to check the threshold on.
	Module     ModuleSpec `yaml:"module"`                // Module used when expression is met.
}
