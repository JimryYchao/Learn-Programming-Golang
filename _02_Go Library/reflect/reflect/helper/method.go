package helper

import (
	"reflect"
)

// 类型与方法集

type MethodInfo struct {
	Name   string
	Type   *FuncType
	method *reflect.Method
}

func (m MethodInfo) Method() reflect.Method {
	return *m.method
}

// TODO
func (m MethodInfo) Func() *Func {
	return FromValue[*Func](m.method.Func)
}

type MethodSet struct {
	t   reflect.Type
	num int
	ms  []MethodInfo
}

func MethodOf(t Type) *MethodSet {
	if IsNilType(t) {
		return nil
	}
	mset := &MethodSet{t.Type(), 0, nil}
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
	if s.ms == nil {
		s.ms = make([]MethodInfo, s.num)
		for i := range s.num {
			m := s.Method(uint(i))
			s.ms[i] = MethodInfo{m.Name, TypeWrap(m.Type).(*FuncType), &m}
		}
	}
	return s.ms
}
