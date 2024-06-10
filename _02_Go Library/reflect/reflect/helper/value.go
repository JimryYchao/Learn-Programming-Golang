package helper

import (
	r "reflect"
	"unsafe"
)

//! >>>>>>>>>>>> Value <<<<<<<<<<<<

type Value interface {
	valueof(r.Value) Value
	IsValid() bool
	Kind() r.Kind
	Type() Type
	String() string
	To() toValue
}

type valueBase struct {
	v r.Value
}

func (b *valueBase) valueof(r.Value) Value { return b }
func (v *valueBase) IsValid() bool         { return true }
func (v *valueBase) Kind() r.Kind          { return v.v.Kind() }
func (v *valueBase) Type() Type            { return typeWrap(v.v.Type()) }
func (b *valueBase) String() string        { return b.v.String() }
func (v *valueBase) To() toValue           { return tovalue{valueWrap(v.v)} }

// todo
func newValue(v r.Value) *valueBase {
	return &valueBase{v}
}

func getValue[V Value](v r.Value) Value {
	return V.valueof(*(*V)(unsafe.Pointer(new(V))), v)
}

func valueWrap(v r.Value) Value {
	if v.Kind() == r.Invalid {
		return NilValue{}
	}
	switch v.Kind() {
	case r.Slice:
		return getValue[Slice](v)
	case r.Array:
		return getValue[Array](v)
	default:
		return &valueBase{v}
	}
}

// 构造 Type 零值
func Zero(t Type) Value {
	if IsNilType(t) {
		return NilValue{}
	}
	return valueWrap(r.Zero(t.Type()))
}

func ValueOf(i any) Value {
	return valueWrap(r.ValueOf(i))
}

func ValueWrap(v r.Value) Value {
	return valueWrap(v)
}

func IsValuep[V Value](v Value) bool {
	if v == nil {
		return false
	}
	return (*(*V)(unsafe.Pointer(new(V)))).Kind() == v.Kind()
}

//! >>>>>>>>>>>> toValue <<<<<<<<<<<<

type toValue interface {
	// Array() Array
	// Chan() Chan
	// Pointer() Pointer
	// Func() Func
	// Map() Map
	// Struct() Struct
	// Slice() Slice
	// Float() Float
	// Int() Int
	// String() String
	// Bytes() Bytes
	// Bool() Bool
	// Complex() Complex
	// Uint() Uint
	//TODO Interface()
}

type tovalue struct {
	Value
}

func toV[V Value](v Value) V {
	if v, ok := v.(V); ok {
		return v
	} else {
		return *(*V)(unsafe.Pointer(new(V)))
	}
}

func ValueFrom[V Value](i any) Value {
	v := ValueOf(i)
	return toV[V](v)
}

//! >>>>>>>>>>>> ValueCommon <<<<<<<<<<<<

type ValueCommon interface {
	Comparable() bool
	// Equal() bool
	IsZero()
}
type valueCom struct {
	*valueBase
}

func (c *valueCom) Comparable() bool { return c.v.Comparable() }
func (c *valueCom) IsZero() bool     { return c.v.IsZero() }

func ValueCom(v Value) (ValueCommon, error) {
	// if !IsNilType()
	return nil, nil
}

// func FromValue[V Value](v r.Value) V {
// 	// TODO
// 	return *(*V)(nil)
// }

//! >>>>>>>>>>>>>> Nil Value <<<<<<<<<<<<<<<<

type NilValue struct{}

func (n NilValue) valueof(r.Value) Value { return nil }
func (n NilValue) IsValid() bool         { return false }
func (n NilValue) Kind() r.Kind          { return r.Invalid }
func (n NilValue) Type() Type            { return Nil{} }
func (n NilValue) String() string        { return "<nil value>" }
func (n NilValue) To() toValue           { return nil }
