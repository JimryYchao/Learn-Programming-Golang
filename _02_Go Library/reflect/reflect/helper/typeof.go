package helper

import (
	"fmt"
	r "reflect"
)

type TypeHelper struct {
	t r.Type
}

func (t TypeHelper) Type() r.Type       { return t.t }
func (t TypeHelper) Kind() r.Kind       { return t.t.Kind() }
func (t TypeHelper) typeof(r.Type) Type { return t }

type Type interface {
	Type() r.Type
	Kind() r.Kind
	typeof(r.Type) Type
}

func getType[T Type](tp r.Type) Type {
	var t T
	return t.typeof(tp)
}
func getKind[T Type]() r.Kind {
	var t T
	return t.Kind()
}
func typeof(tp r.Type) Type {
	switch tp.Kind() {
	case r.Slice:
		return getType[*SliceType](tp)
	// case r.Map:
	// 	return getType[*MapType](tp)

	default:
		return TypeHelper{tp} // 返回常规 reflect.Type
	}
}

func TypeTo[T Type](tp r.Type) (T, error) {
	var t T
	if tp.Kind() != getKind[T]() {
		return t, newErr(fmt.Sprintf("tp.Kind() is %s but not %s", tp.Kind(), t.Kind()))
	}
	t.typeof(tp)
	return t, nil
}

func TypeFor[T any]() Type {
	return typeof(r.TypeOf((*T)(nil)).Elem())
}

func TypeOf(i any) Type {
	return typeof(r.TypeOf(i))
}

func TypeWrap(tp r.Type) Type {
	return typeof(tp)
}
