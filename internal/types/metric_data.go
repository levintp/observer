package types

// The `MetricData` type is a structure that defines a sample returned by a metric.
type MetricData struct {
	Value  float64
	Labels map[string]any
}

// Function to convert the metric data into a sample.
func (data MetricData) ToSample() Sample {
	sample := make(Sample)

	// Copy the labels map into the sample.
	for label, value := range data.Labels {
		sample[label] = value
	}

	// Set the value of the sample.
	sample[SampleValueKey] = data.Value

	return sample
}
