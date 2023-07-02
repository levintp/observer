package types

// The `StreamSpec` type is a structure that defines the specification of a
// stream of metrics.
type StreamSpec struct {
	Name       string                 `yaml:"name"`                 // Name of the stream.
	Categories []Category             `yaml:"categories,omitempty"` // Categories to run the stream on.
	Metrics    map[string]*MetricSpec `yaml:"metrics"`              // Metrics in the stream.
}
