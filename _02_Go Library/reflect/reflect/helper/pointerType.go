package helper

import (
	r "reflect"
)

type PointerType = *pointerType

type pointerType struct {
	*typeBase
}

func (t PointerType) typeof(tp r.Type) Type {
	t = &pointerType{newType(tp)}
	return t
}

func (PointerType) Kind() r.Kind         { return r.Pointer }
func (t PointerType) Common() TypeCommon { return TypeCom(t) }

func (t PointerType) Elem() Type { return typeof(t.t.Elem()) }

// PointerTo

func PointerTo(tp r.Type) (PointerType, error) {
	if tp == nil {
		return nil, ErrTypeNil
	}
	return &pointerType{newType(r.PointerTo(tp))}, nil
}

func PointerFor[T any]() PointerType {
	s, _ := PointerTo(r.TypeFor[T]())
	return s
}
