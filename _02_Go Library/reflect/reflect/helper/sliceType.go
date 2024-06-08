package helper

import (
	r "reflect"
)

// SliceType
var sliceSet map[string]*SliceType = make(map[string]*SliceType)

type SliceType struct {
	*typeBase
}

func (t *SliceType) typeof(tp r.Type) Type {
	if value, ok := sliceSet[tp.String()]; ok {
		t = value
		return value
	}
	t = &SliceType{newType(tp)}
	sliceSet[tp.String()] = t
	return t
}

func (*SliceType) Kind() r.Kind { return r.Slice }

func (t *SliceType) Elem() Type { return typeof(t.t.Elem()) }

func (t *SliceType) Common() TypeCommon { return TypeCom(t) }

// SliceOf
func SliceOf(tp r.Type) *SliceType {
	if tp == nil {
		return nil
	}
	st := &SliceType{newType(r.SliceOf(tp))}
	sliceSet[tp.String()] = st
	return st
}

func SliceFor[T any]() *SliceType {
	return SliceOf(r.TypeFor[T]())
}

// New Slice
func (t *SliceType) new(len, cap int) (*Slice, error) {
	if len < 0 {
		return nil, newErr("len is a negative number")
	}
	if cap < len {
		return nil, newErr("cap is less than len")
	}

	slice := r.MakeSlice(t.t, len, cap)
	return &Slice{&slice, nil}, nil
}

func (t *SliceType) New(len int) (*Slice, error) {
	return t.NewC(len, len)
}
func (t *SliceType) NewC(len, cap int) (*Slice, error) {
	if t != nil {
		return t.new(len, cap)
	}
	return nil, newErr("SliceType is invalid")
}
