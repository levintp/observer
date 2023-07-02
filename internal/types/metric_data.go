package types

// The `MetricData` type is a structure that defines a sample returned by a metric.
type MetricData struct {
	Value  float64
	Labels map[string]any
}
