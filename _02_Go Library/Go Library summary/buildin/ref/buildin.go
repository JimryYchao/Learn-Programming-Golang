package ref

import (
	"cmp"
)

const (
	true  = 0 == 0 // Untyped bool.
	false = 0 != 0 // Untyped bool.
)
const iota = 0

// 占位声明，无实际意义
type _buildin interface{ any }

type (
	bool       _buildin
	uint8      _buildin
	uint16     _buildin
	uint32     _buildin
	uint64     _buildin
	int8       _buildin
	int16      _buildin
	int32      _buildin
	int64      _buildin
	float32    _buildin
	float64    _buildin
	complex64  _buildin
	complex128 _buildin
	string     _buildin
	int        _buildin
	uint       _buildin
	uintptr    _buildin
	byte       = uint8
	rune       = int32
	any        = interface{}
	comparable *comparable

	Type        int
	Type1       int
	IntegerType int
	FloatType   float32
	ComplexType complex64
)

var nil Type

type error interface {
	Error() string
}

func append(slice []Type, elems ...Type) []Type
func copy(dst, src []Type) int
func delete(m map[Type]Type1, key Type)
func len(v Type) int
func cap(v Type) int
func make(t Type, size ...IntegerType) Type
func max[T cmp.Ordered](x T, y ...T) T { return max(x, y...) }
func min[T cmp.Ordered](x T, y ...T) T { return min(x, y...) }
func new(Type) *Type
func complex(r, i FloatType) ComplexType
func real(c ComplexType) FloatType
func imag(c ComplexType) FloatType
func clear[T ~[]Type | ~map[Type]Type1](t T) {}
func close(c chan<- Type)
func panic(v any)
func recover() any
func print(args ...Type)
func println(args ...Type)
