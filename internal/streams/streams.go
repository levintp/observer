package streams

import (
	"github.com/levintp/observer/internal/config"
	"github.com/levintp/observer/internal/types"
)

// The `Stream` type is a structure that holds the specification of a stream
// from the configuration and a target queue for the metrics' samples.
type Stream struct {
	spec    *types.StreamSpec
	samples types.Queue[types.Sample]
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
func (steram Stream) Start() error {
	// TODO:
	//  - for each metric in the stream:
	// 	  - get the configured module plugin for the metric.
	//	  - get the configured function from the module plugin.
	// 	  - start running the function with the configured arguments in a
	//		separate go-routine.
	//	- if any step of the way fails, return an error.
	return nil
}

// Function to collect a sample from the stream.
//
// Alias for `stream.Samples.Pop()`
func (stream Stream) Collect() (types.Sample, error) {
	return stream.samples.Pop()
}
