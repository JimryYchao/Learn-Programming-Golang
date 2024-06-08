package helper

import (
	r "reflect"
)

// SliceType
var arraySet map[string]*ArrayType = make(map[string]*ArrayType)

type ArrayType struct {
	*typeBase
	len int
}

func (t *ArrayType) typeof(tp r.Type) Type {
	if value, ok := arraySet[tp.String()]; ok {
		t = value
		return value
	}
	t = &ArrayType{newType(tp), tp.Len()}
	arraySet[tp.String()] = t
	return t
}

func (*ArrayType) Kind() r.Kind { return r.Array }

func (t *ArrayType) Elem() Type { return typeof(t.t.Elem()) }

func (t *ArrayType) Len() int { return t.len }

func (t *ArrayType) Common() TypeCommon { return TypeCom(t) }

// ArrayOf
func ArrayOf(length uint, tp r.Type) *ArrayType {
	if tp == nil {
		return nil
	}
	at := &ArrayType{newType(r.ArrayOf(int(length), tp)), int(length)}
	arraySet[tp.String()] = at
	return at
}

func ArrayFor[T any](length uint) *ArrayType {
	return ArrayOf(length, r.TypeFor[T]())
}

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
