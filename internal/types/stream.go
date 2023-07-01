package types

type Stream struct {
	Spec    *StreamSpec
	Samples Queue[Sample]
}
