package helper

import (
	r "reflect"
)

//! >>>>>>>>>>>> ChanType <<<<<<<<<<<<

type ChanType = *chanType
type chanType struct {
	*typeBase
}

func (t ChanType) typeof(tp r.Type) Type {
	t = &chanType{newType(tp)}
	return t
}

func (ChanType) Kind() r.Kind         { return r.Chan }
func (t ChanType) Common() TypeCommon { return toTypeCom(t) }

func (t ChanType) Elem() Type         { return typeWrap(t.t.Elem()) }
func (t ChanType) ChanDir() r.ChanDir { return t.t.ChanDir() }

//! >>>>>>>>>>>> ChDir <<<<<<<<<<<<

type ChDir struct{ dir int } // 0 对应 BithDir
var (
	RecvDir ChDir = ChDir{1}
	SendDir ChDir = ChDir{2} // chan<-
	BothDir ChDir = ChDir{0} // chan
)

func (d ChDir) toChanDir() r.ChanDir {
	switch d.dir {
	case 1:
		return r.RecvDir
	case 2:
		return r.SendDir
	default:
		return r.BothDir
	}
}
func (d ChDir) String() string {
	return d.toChanDir().String()
}

// ChanOf

func ChanOf(dir ChDir, t r.Type) (ChanType, error) {
	if t == nil {
		return nil, ErrOutOfRange
	}
	if t.Size() > 1<<16-1 { // 65535
		return nil, ErrChanElemSize
	}
	ctp := r.ChanOf(dir.toChanDir(), t)
	return &chanType{newType(ctp)}, nil
}

func ChanFor[E any](dir ChDir) ChanType {
	ch, _ := ChanOf(dir, r.TypeFor[E]())
	return ch
}
