package helper

import (
	"reflect"
	r "reflect"
	"testing"
)

type Value interface {
	Type() Type
	Kind() r.Kind
	IsValid() bool
	To() toValue
}

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

type valueBase struct {
	v r.Value
}

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

func valueWrap(v r.Value) Value {
	if v.Kind() == r.Invalid {

	}
	return nil
}
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

// func FromValue[V Value](v r.Value) V {
// 	// TODO
// 	return *(*V)(nil)
// }

func Test(t *testing.T) {
	var r chan int
	r = nil
	v := reflect.ValueOf(r)
	log(v.Kind())
	log(v.IsZero())
	log(v.IsValid())
	v = reflect.Zero(reflect.TypeFor[[]int]())
	log(v.Kind())
	log(v.IsZero())
	log(v.IsValid())

}

type NilValue struct{}

func (n NilValue) Type() Type    { return Nil{} }
func (n NilValue) To() toValue   { return nil }
func (n NilValue) Kind() r.Kind  { return r.Invalid }
func (n NilValue) IsValid() bool { return false }
