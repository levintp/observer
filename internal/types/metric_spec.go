package types

// The `MetricSpec` type is a structure that defines the specification of a
// collectable metric.
type MetricSpec struct {
	Name       string                    `yaml:"name"`                  // Name of the metric.
	Interval   int                       `yaml:"interval" default:"60"` // Collection interval.
	Categories []Category                `yaml:"categories,omitempty"`  // Categories to collect the metric on.
	Module     ModuleSpec                `yaml:"module"`                // Collection module in seconds.
	Thresholds map[string]*ThresholdSpec `yaml:"thresholds,omitempty"`  // Thresholds applied to the metric.
}
