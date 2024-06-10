package helper

import (
	r "reflect"
)

//! >>>>>>>>>>>>>> Func <<<<<<<<<<<<<<

type Func = *vFunc

type vFunc struct {
	*valueBase
}

func (v Func) valueof(rv r.Value) Value {
	v = &vFunc{newValue(rv)}
	return v
}
func (v Func) IsValid() bool { return true }
func (v Func) Kind() r.Kind  { return r.Chan }
func (v Func) To() toValue   { return tovalue{v} }

func (v Func) FuncType() FuncType { return v.Type().To().FuncType() }
