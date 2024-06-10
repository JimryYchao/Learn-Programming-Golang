package helper_test

import (
	"fmt"
	. "gostd/reflect/helper"
	"reflect"
	"testing"
)

type mInt int
type MInt int

var log = fmt.Println

func logf(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}

func TestGetType(tt *testing.T) {
	var tp = reflect.TypeFor[[]int]()

	if st := new(SliceType); TryTypeTo(tp, st) {
		fmt.Println((*st).String())
	}

	if n := TypeOf(nil); IsType[Nil](n) {
		fmt.Println(n.String())
	}

	if t := TypeFor[[]int](); IsType[SliceType](t) {
		fmt.Println(t.String())
	}

	if t := TypeWrap(tp); IsType[SliceType](t) {
		fmt.Println(t.String())
	}
}

func TestTypeOf(t *testing.T) {
	type IntInline int
	var st SliceType
	if TryTypeTo(reflect.TypeOf(any([]MInt{})), &st) {
		testTypeCommon(st)
	}
	testTypeCommon(TypeOf(mInt(0)))

	testTypeCommon(TypeOf([]MInt{}).(SliceType))
	testTypeCommon(TypeWrap(reflect.TypeOf(int(0))))
	testTypeCommon(TypeWrap(reflect.TypeOf(IntInline(0))))
	testTypeCommon(TypeFor[struct{ anom int }]())
	testTypeCommon(TypeFor[[]struct{ anom int }]().(SliceType))
	testTypeCommon(TypeOf(nil))
}

func testTypeCommon(t Type) {
	if t == nil {
		log("t is nil")
	}
	logf("\n>>>>>  t : %s  <<<<<", t.Name())
	logf("Type : %s", t.Type())
	logf("String : %s", t.String())
	logf("Kind : %s", t.Kind())

	if !IsNilType(t) {
		t, _ := TypeCom(t)
		logf("Size : %d", t.Size())
		logf("Align : %d", t.Align())
		logf("PkgPath : %s", t.PkgPath())
		logf("Implements any : %t", t.Implements(reflect.TypeFor[any]()))
		logf("AssignableTo any : %t", t.AssignableTo(reflect.TypeFor[any]()))
		logf("ConvertibleTo any : %t", t.ConvertibleTo(reflect.TypeFor[any]()))
		logf("Comparable : %t", t.Comparable())
		logf("FieldAlign : %d", t.FieldAlign())

		// Type Property
		logf("IsDefined : %t", PropFor(t).IsDefined())
		logf("IsBuildIn : %t", PropFor(t).IsBuildIn())
		logf("IsAnonymous : %t", PropFor(t).IsAnonymous())
	}
}
