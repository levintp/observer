package streams

import (
	"github.com/levintp/observer/internal/config"
	"github.com/levintp/observer/internal/log"
	"github.com/levintp/observer/internal/plugins"
	"github.com/levintp/observer/internal/types"
)

// The `Stream` type is a structure that holds the specification of a stream
// from the configuration and a target queue for the metrics' samples.
type Stream struct {
	spec    *types.StreamSpec
	samples chan types.Sample
}

// Function to initialize a new `Stream` structure with a given specification.
func New(spec *types.StreamSpec) Stream {
	return Stream{spec, make(chan types.Sample, config.Get().Agent.BufferSize)}
}

// Function to check if the stream should run on the local node.
func (stream Stream) ShouldRun() bool {
	node, err := config.Get().GetLocalNode()
	if err != nil {
		return false
	}

	return node.InCatrgories(stream.spec.Categories)
}

// Function to start running the stream.
func (stream Stream) Start() {
	for _, metric := range stream.spec.Metrics {
		go stream.metricHandler(metric)
	}
}

// Function to handle execution of a metric in the stream.
func (stream Stream) metricHandler(metric *types.MetricSpec) {
	// Get the sampling function as configured in the metric specification.
	samplingFunc, err := plugins.GetSamplingFunc(metric.Module.Plugin, metric.Module.Function)
	if err != nil {
		log.Errorw("Failed to obtain sampling function",
			"metric", metric.Name,
			"module", metric.Module.Plugin,
			"symbol", metric.Module.Function)
	}

	// Run the sampling function.
	result, err := samplingFunc(metric.Module.Arguments)
	if err != nil {
		log.Errorw("Failed to run sampling function",
			"metric", metric.Name,
			"module", metric.Module.Plugin,
			"symbol", metric.Module.Function)
	}

	// Publish the result of the sampling function execution down-stream.
	for _, data := range result {
		stream.publishData(data)
	}

	// TODO: time interval mechanism.
}

// Function to publish metric data into the stream.
func (stream Stream) publishData(data types.MetricData) {
	stream.samples <- data.ToSample()
}

// Function to collect a sample from the stream.
func (stream Stream) Collect() types.Sample {
	return <-stream.samples
}
