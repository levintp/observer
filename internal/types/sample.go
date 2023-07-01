package types

// The `MetricData` type is a structure that defines a sample returned by a metric.
type MetricData struct {
	Value  float64
	Labels map[string]any
}

// The `Sample` type is a map of generic labels and values that defines a
// sample of data in a stream.
type Sample map[string]any
