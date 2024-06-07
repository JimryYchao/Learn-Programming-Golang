package helper

import (
	"fmt"
	r "reflect"
)

// SliceType
var sliceSet map[string]*SliceType = make(map[string]*SliceType)

func SliceOf(tp r.Type) *SliceType {
	st := &SliceType{r.SliceOf(tp)}
	sliceSet[tp.Name()] = st
	return st
}

func SliceFor[T any]() *SliceType {
	return SliceOf(r.TypeFor[T]())
}

type SliceType TypeHelper

func (*SliceType) Kind() r.Kind { return r.Slice }

func (t *SliceType) Type() r.Type { return t.t }

func (t *SliceType) typeof(tp r.Type) Type {
	if value, ok := sliceSet[tp.String()]; ok {
		t = value
		return value
	}
	t = &SliceType{tp}
	sliceSet[tp.String()] = t
	return t
}

func (t *SliceType) ElemString() string {
	if t != nil {
		return t.t.Elem().String()
	}
	return "<!t=nil>"
}

func TypeOfSlice(v any) (*SliceType, error) {
	return SliceOfType(r.TypeOf(v))
}

func SliceOfType(tp r.Type) (*SliceType, error) {
	if tp.Kind() != r.Slice {
		return nil, newErr("v is not a slice")
	}
	if value, ok := sliceSet[tp.Elem().Name()]; ok {
		return value, nil
	}
	st := &SliceType{tp.Elem()}
	sliceSet[tp.Elem().Name()] = st
	return st, nil
}

func (t *SliceType) new(len, cap int) (*Slice, error) {
	if len < 0 {
		return nil, fmt.Errorf("len is a negative number")
	}
	if cap < len {
		return nil, fmt.Errorf("cap is less than len")
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
