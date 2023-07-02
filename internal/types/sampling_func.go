package types

import (
	"reflect"
)

type SamplingFunc func(map[string]any) ([]MetricData, error)

var SamplingFuncType reflect.Type

func init() {
	var anyInterface interface{}
	var errorInterface error

	// Declare the types making the sampling function type.
	argumentsType := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(&anyInterface).Elem())
	dataType := reflect.SliceOf(reflect.TypeOf(MetricData{}))
	errorType := reflect.TypeOf(&errorInterface).Elem()

	SamplingFuncType = reflect.FuncOf([]reflect.Type{argumentsType}, []reflect.Type{dataType, errorType}, false)
}
