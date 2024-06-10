package helper

import (
	r "reflect"
)

//! >>>>>>>>>>>>>> Map <<<<<<<<<<<<<<

type Map = *vMap

type vMap struct {
	*valueBase
}

func (v Map) valueof(rv r.Value) Value {
	v = &vMap{newValue(rv)}
	return v
}
func (v Map) IsValid() bool { return true }
func (v Map) Kind() r.Kind  { return r.Map }
func (v Map) To() toValue   { return tovalue{v} }

func (v Map) MapType() MapType { return v.Type().To().MapType() }
