package helper_test

import (
	. "gostd/reflect/helper"
	"reflect"
	"testing"
)

func TestMapType(t *testing.T) {

	testMapType(MapFor[string, int]())
	testMapType(MapFor[string, []int]())
	testMapType(MapFor[[5]int, []int]())
	testMapType(MapFor[any, []int]())

	testMapType(TypeTo[MapType](map[any]int{}))
	testMapType(TypeTo[MapType](10086))
	log(MapOf(reflect.TypeFor[any](), reflect.TypeFor[any]()))
	log(MapOf(reflect.TypeFor[[]int](), reflect.TypeFor[any]()))
	log(MapOf(reflect.FuncOf(nil, nil, false), reflect.TypeFor[any]()))
	// m := make(map[any]int)
}

func testMapType(m MapType) {
	if m == nil {
		log("m is nil")
		return
	}
	testTypeCommon(m)
	logf("Key:%s, Elem:%s", m.Key(), m.Elem())
}
