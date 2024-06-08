package helper

import (
	"reflect"
)

// 类型与方法集

type MethodInfo struct {
	Name string
}

type MethodSet struct {
	t   reflect.Type
	num int
}

func MethodOf(t Type) *MethodSet {
	if IsNilType(t) {
		return nil
	}
	mset := &MethodSet{t.Type(), 0}
	mset.num = mset.t.NumMethod()
	return mset
}

func (s *MethodSet) NumMethod() int {
	return s.num
}

func (s *MethodSet) Method(i uint) reflect.Method {
	var m reflect.Method = s.t.Method(int(i))
	return m
}

func (s *MethodSet) Methods() []MethodInfo {
	return nil
}
