package helper

import (
	r "reflect"
)

//! >>>>>>>>>>>>>> Chan <<<<<<<<<<<<<<

type Chan = *vChan

type vChan struct {
	*valueBase
}

func (v Chan) valueof(rv r.Value) Value {
	v = &vChan{newValue(rv)}
	return v
}
func (v Chan) IsValid() bool { return true }
func (v Chan) Kind() r.Kind  { return r.Chan }
func (v Chan) To() toValue   { return tovalue{v} }

func (v Chan) ChanType() ChanType { return v.Type().To().ChanType() }

func (v Chan) Cap() int { return v.v.Cap() }
func (v Chan) Len() int { return v.v.Len() }
