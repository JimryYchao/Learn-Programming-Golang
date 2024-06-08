package helper_test

import (
	. "gostd/reflect/helper"
	"testing"
)

func TestArrayType(t *testing.T) {
	testArrayType[int](10)
	testArrayType[struct{ V int }](0)
	testArrayType[[5]string](15)

	if ArrayOf(10, nil) != nil {
		t.Fatal("ArrayOf(nil) is not return nil")
	}

	log(TypeTo[*ArrayType]([20]int{}))       // []int
	log(TypeTo[*ArrayType](nil))             // <nil>
	log(TypeTo[*ArrayType](10))              // <nil>
	log(TypeTo[*ArrayType]([10][][][]int{})) // [10][][][]int
}

func testArrayType[T any](len uint) {
	at := ArrayFor[T](len)
	testTypeCommon(at)
	logf("Len: %d, Elem: %s", at.Len(), at.Elem().String())
}
