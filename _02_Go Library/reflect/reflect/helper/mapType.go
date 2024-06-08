package helper

import (
	r "reflect"
)

var mapSet map[string]*MapType = make(map[string]*MapType)

type MapType struct {
	*typeBase
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
	mapSet[mtp.Name()] = mt
	return mt
}

func MapFor[K comparable, V any]() *MapType {
	return MapOf(r.TypeFor[K](), r.TypeFor[V]())
}

// new
