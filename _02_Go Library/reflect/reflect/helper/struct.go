package helper

import (
	r "reflect"
)

//! >>>>>>>>>>>>>> Struct <<<<<<<<<<<<<<

type Struct = *vStruct

type vStruct struct {
	*valueBase
}

func (v Struct) valueof(rv r.Value) Value {
	v = &vStruct{newValue(rv)}
	return v
}
func (v Struct) IsValid() bool { return true }
func (v Struct) Kind() r.Kind  { return r.Struct }
func (v Struct) To() toValue   { return tovalue{v} }

func (v Struct) StructType() StructType { return v.Type().To().StructType() }
