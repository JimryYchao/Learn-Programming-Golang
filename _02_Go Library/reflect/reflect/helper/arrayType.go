package helper

import (
	r "reflect"
)

type ArrayType = *arrayType

type arrayType struct {
	*typeBase
	len int
}

func (t ArrayType) typeof(tp r.Type) Type {
	t = &arrayType{newType(tp), tp.Len()}
	return t
}

func (ArrayType) Kind() r.Kind { return r.Array }

func (t ArrayType) Elem() Type { return typeof(t.t.Elem()) }

func (t ArrayType) Len() int { return t.len }

func (t ArrayType) Common() TypeCommon { return TypeCom(t) }

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

// func ToArrayType(r)

// New Slice
// func (t *ArrayType) new(len, cap int) (*Slice, error) {
// 	if len < 0 {
// 		return nil, newErr("len is a negative number")
// 	}
// 	if cap < len {
// 		return nil, newErr("cap is less than len")
// 	}

// 	slice := r.MakeSlice(t.t, len, cap)
// 	return &Slice{&slice, nil}, nil
// }

// func (t *ArrayType) New(len int) (*Array, error) {

// 	return t.NewC(len, len)
// }
// func (t *ArrayType) NewC(len, cap int) (*Array, error) {
// 	if t != nil {
// 		return t.new(len, cap)
// 	}
// 	return nil, newErr("SliceType is invalid")
// }
