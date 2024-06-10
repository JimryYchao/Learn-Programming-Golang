package helper_test

import (
	. "gostd/reflect/helper"
	"testing"
)

func TestStruct(t *testing.T) {
	type s1 struct {
		v1 int `tag:"????"`
		v2 int `tag:"????" tag2:"tag2"`
		E1 int `mess tag:"????" tag2:"tag2"`
		E2 int `tag:'????'`
		int
	}

	s := TypeOf([]int{}).To().SliceType()
	log(s)
	fields := VisibleFields(TypeOf(s1{}).To().StructType())

	for _, f := range fields {
		logf("name:%s, tag:%s, tag2:%s", f.Name, f.Tag.Get("tag"), f.Tag.Get("tag2"))
		logf(f.Get("tag"))
	}

	type s2 struct {
		s1
		E1 int
		v1 string
	}
	type s3 struct {
		E1 string
		s2
	}

	testStruct(TypeTo[StructType](s1{}))
	testStruct(TypeTo[StructType](s2{}))
	testStruct(TypeTo[StructType](s3{}))

	// sf, ok := TypeTo[StructType](s3{}).Type().FieldByName("")
	// log(sf, ok)

	f, ok := TypeTo[StructType](s3{}).FieldByIndex([]int{1, 2})
	log(f, ok)
}

func testStruct(s StructType) {
	if s == nil {
		log("struct is a nil")
		return
	}
	// testTypeCommon(s)
	logf("num: %d", s.NumField())
	fs := VisibleFields(s)
	for _, f := range fs {
		logf("field: %s, Kind:%s, Index:%v", f.Name, f.Type().Kind(), f.Index)
	}
	iterMethods(MethodOf(s))
}
