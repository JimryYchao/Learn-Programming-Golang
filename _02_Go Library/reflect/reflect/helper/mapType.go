package helper

import (
	r "reflect"
)

var mapSet map[string]*MapType = make(map[string]*MapType)

type MapType struct {
	*typeBase
}

func (t *MapType) typeof(tp r.Type) Type {
	if value, ok := mapSet[tp.String()]; ok {
		t = value
		return value
	}
	t = &MapType{newType(tp)}
	mapSet[tp.String()] = t
	return t
}

func (*MapType) Kind() r.Kind         { return r.Map }
func (t *MapType) Common() TypeCommon { return TypeCom(t) }

func (t *MapType) Elem() Type { return typeof(t.t.Elem()) }

func (t *MapType) Key() Type { return typeof(t.t.Key()) }

// MapOf
func MapOf(key r.Type, elem r.Type) *MapType {
	if key == nil || !key.Comparable() {
		return nil
	}

	mtp := r.MapOf(key, elem)
	mt := &MapType{newType(mtp)}
	mapSet[mtp.String()] = mt
	return mt
}

func MapFor[K comparable, V any]() *MapType {
	return MapOf(r.TypeFor[K](), r.TypeFor[V]())
}

// new
