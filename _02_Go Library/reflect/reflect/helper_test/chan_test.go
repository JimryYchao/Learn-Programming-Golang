package helper_test

import (
	. "gostd/reflect/helper"
	"reflect"
	"testing"
)

func TestChanType(t *testing.T) {
	testChanType(TypeTo[*ChanType]([]int{}))                                  // ch is nil
	testChanType(TypeTo[*ChanType](nil))                                      // ch is nil
	testChanType(TypeTo[*ChanType](*new(chan<- *int)))                        // chan<- *int
	testChanType(ChanFor[int](reflect.RecvDir))                               // [][][][]int
	testChanType(ChanFor[any](reflect.BothDir))                               // chan interface {}
	testChanType(ChanFor[[88888]string](reflect.SendDir))                     // ch is nil
	testChanType(ChanOf(reflect.SendDir, reflect.TypeFor[bool]()))            // chan<- bool
	testChanType(ChanOf(reflect.SendDir, reflect.TypeFor[[1<<16 + 1]bool]())) // nil, too large
}

func testChanType(ch *ChanType) {
	if ch == nil {
		log("ch is nil")
		return
	}
	testTypeCommon(ch)
	logf("Dir: %s, Elem: %s", ch.ChanDir(), ch.Elem())

}
