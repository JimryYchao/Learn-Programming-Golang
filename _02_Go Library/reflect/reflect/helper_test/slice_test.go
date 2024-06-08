package helper_test

import (
	. "gostd/reflect/helper"
	"testing"
)

func TestSliceType(t *testing.T) {
	testSliceType[int]()
	testSliceType[struct{ V int }]()
	testSliceType[[5]string]()

	if SliceOf(nil) != nil {
		t.Fatal("SliceOf(nil) is not return nil")
	}

	log(TypeTo[*SliceType]([]int{}))       // []int
	log(TypeTo[*SliceType](nil))           // <nil>
	log(TypeTo[*SliceType](10))            // <nil>
	log(TypeTo[*SliceType]([][][][]int{})) // [][][][]int
}

func testSliceType[T any]() {
	st := SliceFor[T]()
	testTypeCommon(st)
	logf("Elem: %s, Kind: %s", st.Elem().String(), st.Elem().Kind())
}

func TestNewSlice(t *testing.T) {
	// TODO
}
