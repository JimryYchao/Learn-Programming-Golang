package helper

import (
	r "reflect"
)

var mapSet map[string]*MapType = make(map[string]*MapType)

type MapType TypeHelper

func (*MapType) Kind() r.Kind { return r.Map }

func MapOf(key r.Type, elem r.Type) *MapType {
	mtp := r.MapOf(key, elem)
	mt := &MapType{mtp}
	mapSet[mtp.Name()] = mt
	return mt
}
