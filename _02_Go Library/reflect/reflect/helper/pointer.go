package helper

import (
	r "reflect"
)

//! >>>>>>>>>>>>>> Pointer <<<<<<<<<<<<<<

type Pointer = *vPointer

type vPointer struct {
	*valueBase
}

func (v Pointer) valueof(rv r.Value) Value {
	v = &vPointer{newValue(rv)}
	return v
}
func (v Pointer) IsValid() bool { return true }
func (v Pointer) Kind() r.Kind  { return r.Pointer }
func (v Pointer) To() toValue   { return tovalue{v} }

func (v Pointer) PointerType() PointerType { return v.Type().To().PointerType() }

// 返回一个 CanSet  pValue
func (v Pointer) Elem() Value {
	// switch kind := v.v.Elem().Kind() {
	// 	case Int,
	// }
	return nil
}

// TODO : MUST *Array
func (v Pointer) Cap() int { return v.v.Cap() }
func (v Pointer) Len() int { return v.v.Len() }

//! >>>>>>>>>>>>>> Pointer Elems <<<<<<<<<<<<<<

type ArrayPtr struct {
	Pointer
}
