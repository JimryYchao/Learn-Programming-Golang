package helper

import (
	"fmt"
	r "reflect"
)

type MapType = *mapType

type mapType struct {
	*typeBase
}

func (t MapType) typeof(tp r.Type) Type {
	t = &mapType{newType(tp)}
	return t
}

func (MapType) Kind() r.Kind         { return r.Map }
func (t MapType) Common() TypeCommon { return toTypeCom(t) }

func (t MapType) Elem() Type { return typeWrap(t.t.Elem()) }

func (t MapType) Key() Type { return typeWrap(t.t.Key()) }

// MapOf
func MapOf(key r.Type, elem r.Type) (MapType, error) {
	if key == nil || elem == nil {
		return nil, ErrTypeNil
	}
	if !key.Comparable() {
		return nil, newErr(fmt.Sprintf("invalid key type: %s", key))
	}

	mtp := r.MapOf(key, elem)
	return &mapType{newType(mtp)}, nil
}

func MapFor[K comparable, V any]() MapType {
	m, _ := MapOf(r.TypeFor[K](), r.TypeFor[V]())
	return m
}

// new
