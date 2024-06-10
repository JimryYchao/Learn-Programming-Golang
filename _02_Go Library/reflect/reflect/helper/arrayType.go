package helper

import (
	r "reflect"
)

//! >>>>>>>>>>>> ArrayType <<<<<<<<<<<<

type ArrayType = *arrayType

type arrayType struct {
	*typeBase
	len int
}

func (t ArrayType) typeof(tp r.Type) Type {
	t = &arrayType{newType(tp), tp.Len()}
	return t
}
func (ArrayType) Kind() r.Kind         { return r.Array }
func (t ArrayType) Elem() Type         { return typeWrap(t.t.Elem()) }
func (t ArrayType) Len() int           { return t.len }
func (t ArrayType) Common() TypeCommon { return toTypeCom(t) }
func (t ArrayType) To() toType         { return totype{t} }

// ArrayOf
func ArrayOf(length int, tp r.Type) (ArrayType, error) {
	if tp == nil {
		return nil, ErrTypeNil
	}
	if length < 0 {
		return nil, ErrNegative
	}
	return &arrayType{newType(r.ArrayOf(int(length), tp)), int(length)}, nil
}

func ArrayFor[T any](length int) ArrayType {
	a, _ := ArrayOf(length, r.TypeFor[T]())
	return a
}
