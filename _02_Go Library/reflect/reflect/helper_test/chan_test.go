package helper_test

import (
	. "gostd/reflect/helper"
	"reflect"
	"testing"
)

func TestChanType(t *testing.T) {
	testChanType(TypeTo[ChanType]([]int{}))           // ch is nil
	testChanType(TypeTo[ChanType](nil))               // ch is nil
	testChanType(TypeTo[ChanType](*new(chan<- *int))) // chan<- *int
	testChanType(ChanFor[int](RecvDir))               // [][][][]int
	testChanType(ChanFor[any](BothDir))               // chan interface {}
	testChanType(ChanFor[[88888]string](SendDir))     // ch is nil
	ch, _ := ChanOf(SendDir, reflect.TypeFor[bool]()) // chan<- bool
	testChanType(ch)
	log(ChanOf(SendDir, reflect.TypeFor[[1<<16 + 1]bool]())) // nil, too large
}

func testChanType(ch ChanType) {
	if ch == nil {
		log("ch is nil")
		return
	}
	testTypeCommon(ch)
	logf("Dir: %s, Elem: %s", ch.ChanDir(), ch.Elem())

}
