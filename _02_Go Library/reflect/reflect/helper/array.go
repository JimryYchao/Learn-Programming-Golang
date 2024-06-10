package helper

import (
	r "reflect"
)

//! >>>>>>>>>>>>>> Array <<<<<<<<<<<<<<

type Array = *vArray

type vArray struct {
	*valueBase
}

func (v Array) valueof(rv r.Value) Value {
	v = &vArray{newValue(rv)}
	return v
}
func (v Array) IsValid() bool { return true }
func (v Array) Kind() r.Kind  { return r.Array }
func (v Array) To() toValue   { return tovalue{v} }

func (v Array) ArrayType() ArrayType { return v.Type().To().ArrayType() }

func (v Array) Cap() int { return v.v.Cap() }
