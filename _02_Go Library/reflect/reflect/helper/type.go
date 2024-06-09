package helper

import (
	r "reflect"
	"unsafe"
)

func f() {
	// TypeCommon()
}

type Type interface {
	typeof(r.Type) Type
	Kind() r.Kind
	Type() r.Type
	String() string
	Name() string
	To() toType
}

type toType interface {
	ArrayType() ArrayType
	ChanType() ChanType
	PointerType() PointerType
	FuncType() FuncType
	MapType() MapType
	StructType() StructType
	SliceType() SliceType
}

type to struct {
	Type
}

func (t to) ArrayType() ArrayType     { return To[ArrayType](t.Type) }
func (t to) ChanType() ChanType       { return To[ChanType](t.Type) }
func (t to) PointerType() PointerType { return To[PointerType](t.Type) }
func (t to) FuncType() FuncType       { return To[FuncType](t.Type) }
func (t to) MapType() MapType         { return To[MapType](t.Type) }
func (t to) StructType() StructType   { return To[StructType](t.Type) }
func (t to) SliceType() SliceType     { return To[SliceType](t.Type) }

// func (t to) FuncType() *FuncType       { return To[*FuncType](t.Type) }

// func (t to) Pointer() *

// Type 接口转发的底层实现
type typeBase struct {
	t r.Type
}

func (c *typeBase) typeof(r.Type) Type { return c }
func (c *typeBase) String() string     { return c.t.String() }
func (c *typeBase) Kind() r.Kind       { return c.t.Kind() }
func (c *typeBase) Type() r.Type       { return c.t }
func (c *typeBase) Name() string       { return c.t.Name() }
func (c *typeBase) To() toType         { return to{typeof(c.t)} }

// Type 的通用方法
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

func TypeCom(c Type) TypeCommon {
	if !IsNilType(c) {
		return &typeCom{&typeBase{c.Type()}}
	}
	return nil
}

func (c *typeCom) Size() uintptr               { return c.t.Size() }
func (c *typeCom) Align() int                  { return c.t.Align() }
func (c *typeCom) PkgPath() string             { return c.t.PkgPath() }
func (c *typeCom) Implements(u r.Type) bool    { return c.t.Implements(u) }
func (c *typeCom) AssignableTo(u r.Type) bool  { return c.t.AssignableTo(u) }
func (c *typeCom) ConvertibleTo(u r.Type) bool { return c.t.ConvertibleTo(u) }
func (c *typeCom) Comparable() bool            { return c.t.Comparable() }
func (c *typeCom) FieldAlign() int             { return c.t.FieldAlign() }

func newType(tp r.Type) *typeBase {
	return &typeBase{tp}
}

func getType[T Type](tp r.Type) Type {
	return T.typeof((*new(T)), tp)
}

func typeof(tp r.Type) Type {
	if tp == nil {
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

// 尝试包装为特定的 Type
func TryTypeTo[T Type](tp r.Type, out *T) bool {
	if tp.Kind() != (*out).Kind() {
		return false
	}
	*out = (*out).typeof(tp).(T)
	return true
}

func TypeTo[T Type](i any) T {
	t := TypeOf(i)
	if t.Kind() != (*new(T)).Kind() {
		return *(new(T))
	}
	return (t).(T)
}

// 从类型构造一个 Type
func TypeFor[T any]() Type {
	return typeof(r.TypeOf((*T)(nil)).Elem())
}

// 从 v 提取一个 Type
func TypeOf(i any) Type {
	return typeof(r.TypeOf(i))
}

// 包装一个 reflect.Type
func TypeWrap(tp r.Type) Type {
	return typeof(tp)
}

// 检查包装类型
func Is[T Type](t Type) bool {
	return (*new(T)).Kind() == t.Kind()
}

func To[T Type](t Type) T {
	if Is[T](t) {
		return t.(T)
	}
	return *(*T)(unsafe.Pointer(new(T)))
}

// 附加属性
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

// Nil for nil value
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

type TypeErr struct {
	err string
}

func (e *TypeErr) Error() string {
	return e.err
}

func newErr(s string) error {
	return &TypeErr{s}
}

var ErrTypeNil = newErr("reflect.Type is nil")
var ErrOutOfRange = newErr("index out of range")
var ErrChanElemSize = newErr("element size too large")
var ErrNegative = newErr("negative argument passed to parameter")
var ErrTooManyArgus = newErr("too many arguments")
var ErrVaNotSlice = newErr("last arg of variadic func must be slice")

// var ErrTypeNil = newErr("reflect.Type is nil")
