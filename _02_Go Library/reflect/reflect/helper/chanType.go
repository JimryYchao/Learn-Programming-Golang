package helper

import (
	r "reflect"
)

var chanSet map[string]*ChanType = make(map[string]*ChanType)

type ChanType struct {
	*typeBase
}

func (t *ChanType) typeof(tp r.Type) Type {
	if value, ok := chanSet[tp.String()]; ok {
		t = value
		return value
	}
	t = &ChanType{newType(tp)}
	chanSet[tp.String()] = t
	return t
}

func (*ChanType) Kind() r.Kind         { return r.Chan }
func (t *ChanType) Common() TypeCommon { return TypeCom(t) }

func (t *ChanType) Elem() Type         { return typeof(t.t.Elem()) }
func (t *ChanType) ChanDir() r.ChanDir { return t.t.ChanDir() }

// ChanOf

func ChanOf(dir r.ChanDir, t r.Type) *ChanType {
	if t == nil || t.Size() > 1<<16-1 { // 65535
		return nil
	}
	ctp := r.ChanOf(dir, t)
	ct := &ChanType{newType(ctp)}
	chanSet[ctp.String()] = ct
	return ct
}

func ChanFor[E any](dir r.ChanDir) *ChanType {
	return ChanOf(dir, r.TypeFor[E]())
}
