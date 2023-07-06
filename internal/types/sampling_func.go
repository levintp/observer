package types

import (
	"reflect"
)

type SamplingFunc func(ModuleArguments) ([]MetricData, error)

var SamplingFuncType reflect.Type

func init() {
	var funcInterface SamplingFunc
	SamplingFuncType = reflect.TypeOf(funcInterface)
}
