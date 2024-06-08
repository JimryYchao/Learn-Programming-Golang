package helper

import "reflect"

type Value interface {
}

func FromValue[V Value](v reflect.Value) V {
	// TODO
	return *(*V)(nil)
}
