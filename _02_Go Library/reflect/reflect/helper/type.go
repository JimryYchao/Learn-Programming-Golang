package helper

import (
	r "reflect"
	"unsafe"
)

//! >>>>>>>>>>>>>> Type <<<<<<<<<<<<<<<<

type Type interface {
	typeof(r.Type) Type
	Kind() r.Kind
	Type() r.Type
	Name() string
	String() string
	To() toType
}
type typeBase struct {
	t r.Type
}

func (b *typeBase) typeof(r.Type) Type { return b }
func (b *typeBase) Kind() r.Kind       { return b.t.Kind() }
func (b *typeBase) Type() r.Type       { return b.t }
func (b *typeBase) Name() string       { return b.t.Name() }
func (b *typeBase) String() string     { return b.t.String() }
func (b *typeBase) To() toType         { return totype{typeWrap(b.t)} }

func newType(tp r.Type) *typeBase {
	return &typeBase{tp}
}

func getType[T Type](tp r.Type) Type {
	return T.typeof(*(*T)(unsafe.Pointer(new(T))), tp)
}

func typeWrap(tp r.Type) Type {
	if tp == nil || tp.Kind() == r.Invalid {
		return Nil{}
	}
	switch tp.Kind() {
	case r.Slice:
		return getType[SliceType](tp)
	case r.Map:
		return getType[MapType](tp)
	case r.Array:
		return getType[ArrayType](tp)
	case r.Func:
		return getType[FuncType](tp)
	case r.Chan:
		return getType[ChanType](tp)
	case r.Struct:
		return getType[StructType](tp)
	case r.Pointer:
		return getType[PointerType](tp)
	default:
		return &typeBase{tp} // 返回常规 reflect.Type
	}
}

// 从类型构造一个 Type
func TypeFor[T any]() Type {
	return typeWrap(r.TypeOf((*T)(nil)).Elem())
}

// 从 v 提取一个 Type
func TypeOf(i any) Type {
	return typeWrap(r.TypeOf(i))
}

// 包装一个 reflect.Type
func TypeWrap(tp r.Type) Type {
	return typeWrap(tp)
}

// 检查包装类型
func IsType[T Type](t Type) bool {
	if t == nil {
		return false
	}
	return (*(*T)(unsafe.Pointer(new(T)))).Kind() == t.Kind()
}

// 尝试包装为特定的 Type
func TryTypeTo[T Type](tp r.Type, out *T) bool {
	if tp.Kind() != (*out).Kind() {
		return false
	}
	*out = (*out).typeof(tp).(T)
	return true
}

//! >>>>>>>>>>>>>> totype <<<<<<<<<<<<<<<<

type toType interface {
	ArrayType() ArrayType
	ChanType() ChanType
	PointerType() PointerType
	FuncType() FuncType
	MapType() MapType
	StructType() StructType
	SliceType() SliceType
}

type totype struct {
	Type
}

func (t totype) ArrayType() ArrayType     { return toT[ArrayType](t.Type) }
func (t totype) ChanType() ChanType       { return toT[ChanType](t.Type) }
func (t totype) PointerType() PointerType { return toT[PointerType](t.Type) }
func (t totype) FuncType() FuncType       { return toT[FuncType](t.Type) }
func (t totype) MapType() MapType         { return toT[MapType](t.Type) }
func (t totype) StructType() StructType   { return toT[StructType](t.Type) }
func (t totype) SliceType() SliceType     { return toT[SliceType](t.Type) }

func toT[T Type](t Type) T {
	if t, ok := t.(T); ok {
		return t
	} else {
		return *(*T)(unsafe.Pointer(new(T)))
	}
}

func TypeTo[T Type](i any) T {
	t := TypeOf(i)
	return toT[T](t)
}

//! >>>>>>>>>>>>>> TypeCommon <<<<<<<<<<<<<<<<

type TypeCommon interface {
	Type
	Size() uintptr
	Align() int
	PkgPath() string
	Implements(r.Type) bool
	AssignableTo(r.Type) bool
	ConvertibleTo(r.Type) bool
	Comparable() bool
	FieldAlign() int
}
type typeCom struct {
	*typeBase
}

func (c *typeCom) Size() uintptr               { return c.t.Size() }
func (c *typeCom) Align() int                  { return c.t.Align() }
func (c *typeCom) PkgPath() string             { return c.t.PkgPath() }
func (c *typeCom) Implements(u r.Type) bool    { return c.t.Implements(u) }
func (c *typeCom) AssignableTo(u r.Type) bool  { return c.t.AssignableTo(u) }
func (c *typeCom) ConvertibleTo(u r.Type) bool { return c.t.ConvertibleTo(u) }
func (c *typeCom) Comparable() bool            { return c.t.Comparable() }
func (c *typeCom) FieldAlign() int             { return c.t.FieldAlign() }

func TypeCom(c Type) (TypeCommon, error) {
	if !IsNilType(c) {
		return &typeCom{&typeBase{c.Type()}}, nil
	}
	return nil, ErrTypeNil
}

func toTypeCom(c Type) TypeCommon {
	if !IsNilType(c) {
		return &typeCom{&typeBase{c.Type()}}
	}
	return nil
}

//! >>>>>>>>>>>>>> TypeProperty <<<<<<<<<<<<<<<<

type TypeProperty interface {
	IsDefined() bool
	IsBuildIn() bool
	IsAnonymous() bool
}

type typeProper struct {
	com TypeCommon
}

func PropFor(c TypeCommon) TypeProperty {
	return typeProper{c}
}

func (c typeProper) IsDefined() bool   { return (c.com).Name() != "" }
func (c typeProper) IsBuildIn() bool   { return (c.com).Name() != "" && (c.com).PkgPath() == "" }
func (c typeProper) IsAnonymous() bool { return (c.com).Name() == "" && (c.com).PkgPath() == "" }

//! >>>>>>>>>>>>>> Nil <<<<<<<<<<<<<<<<

type Nil struct{}

func (n Nil) typeof(r.Type) Type { return n }
func (n Nil) Type() r.Type       { return nil }
func (n Nil) Kind() r.Kind       { return r.Invalid }
func (n Nil) String() string     { return "<nil>" }
func (n Nil) Name() string       { return "nil" }
func (c Nil) To() toType         { return nil }
func IsNilType(t Type) bool {
	_, ok := t.(Nil)
	return ok
}

//! >>>>>>>>>>>>>> Err <<<<<<<<<<<<<<<<

type TypeErr struct {
	err string
}

func (e *TypeErr) Error() string {
	return e.err
}

func newErr(s string) error {
	return &TypeErr{s}
}

var ErrTypeNil = newErr("Type is nil")
var ErrOutOfRange = newErr("index out of range")
var ErrChanElemSize = newErr("element size too large")
var ErrNegative = newErr("negative argument passed to parameter")
var ErrTooManyArgus = newErr("too many arguments")
var ErrVaNotSlice = newErr("last arg of variadic func must be slice")

// var ErrTypeNil = newErr("reflect.Type is nil")
