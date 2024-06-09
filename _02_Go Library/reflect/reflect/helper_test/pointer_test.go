package helper_test

import (
	. "gostd/reflect/helper"
	"testing"
)

func TestPointerType(t *testing.T) {
	testPointerType[int]()
	testSliceType[struct{ V int }]()
	testPointerType[[5]string]()

	if s, _ := PointerTo(nil); s != nil {
		t.Fatal("SliceOf(nil) is not return nil")
	}

	log(TypeTo[PointerType]([]int{}))          // []int
	log(TypeTo[PointerType](nil))              // <nil>
	log(TypeTo[PointerType](new(PointerType))) // <nil>
	log(TypeTo[PointerType](new(int)))         //
}

func testPointerType[T any]() {
	st := PointerFor[T]()
	testTypeCommon(st)
	logf("Elem: %s, Kind: %s", st.Elem().String(), st.Elem().Kind())
}
