package helper_test

import (
	. "gostd/reflect/helper"
	"reflect"
	"testing"
)

func TestFuncType(t *testing.T) {

	i := reflect.TypeFor[int]()
	s := reflect.TypeFor[string]()
	b := reflect.TypeFor[bool]()
	e := reflect.TypeFor[error]()
	f := reflect.TypeFor[float64]()
	slice := reflect.TypeFor[[]any]()

	testFuncType[int]([]reflect.Type{}, []reflect.Type{b}, true)
	testFuncType[int]([]reflect.Type{i, s}, []reflect.Type{}, false)
	testFuncType[int]([]reflect.Type{f, s, slice}, nil, true)
	testFuncType[int](nil, []reflect.Type{e}, true)
	testFuncType[int](nil, nil, false)
}

func testFuncType[T any](in []reflect.Type, out []reflect.Type, isVar bool) {
	var ft *FuncType
	if isVar {
		ft = FuncOfVar[T](in, out)
	} else {
		ft = FuncOf(in, out)
	}

	testTypeCommon(ft)
	logf("IsVariadic: %t", ft.IsVariadic())
	logf("NumIn: %d", ft.NumIn())
	logf("Ins: %s", ft.Ins())
	logf("NumOut: %d", ft.NumOut())
	logf("Outs: %s", ft.Outs())
}

// func fmtTypes(tps []Type) string {
// 	var sb strings.Builder
// 	for _, v := range tps {
// 		sb.WriteString(v.String() + ", ")
// 	}
// 	return
// }
