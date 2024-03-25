## Package buildin

内置包 `buildin` 提供了 Go 预先声明的标识符的文档。这里记录的项实际上并不是内置的，但它们的描述允许 godoc 为语言的特殊标识符提供文档。

>---
### const bool

```go
const (
    true  = 0 == 0 // Untyped bool.
    false = 0 != 0 // Untyped bool.
)
```

`true` 和 `false` 是两个无类型的布尔值。

>---
### const iota

```go
const iota = 0 // Untyped int.
```

`iota` 是一个预先声明的标识符，表示（通常带括号的）`const` 声明中当前常量规范的非类型化整数序号。它从零索引开始。

```go
const (
	Sunday    1 << iota // 1
	Monday              // 2
	Tuesday             // 4
	Wednesday           // 8
	Thursday            // 16
	Friday              // 32
	Saturday            // 64
)
```

>---
### builtin types

```go
// 预定义类型声明
type bool bool
type uint8 uint8
type uint16 uint16
type uint32 uint32
type uint64 uint64
type int8 int8
type int16 int16
type int32 int32
type int64 int64
type float32 float32
type float64 float64
type complex64 complex64
type complex128 complex128
type string string
type int int
type uint uint
type uintptr uintptr
type comparable interface{ comparable }

// 以下是用于文档目的声明
type Type int
type Type1 int
type IntegerType int
type FloatType float32
type ComplexType complex64
```

`comparable` 是一个由所有可比较类型（布尔值，数字，字符串，指针，通道，可比较类型的数组，字段都是可比较类型的结构）实现的接口。可比较接口只能用作类型参数约束，而不能用作变量的类型。

`complex128` 是具有 `float64` 实部和虚部的所有复数的集合。

`complex64` 是具有 `float32` 实部和虚部的所有复数的集合。

`uintptr` 是一个整数类型，它是足够容纳保存任何指针的位模式。


>---
### type alias

```go
type byte = uint8
type rune = int32
type any = interface{}
```

`byte` 是 `uint8` 的别名，在所有方面都等同于 `uint8`。按照惯例，它用于区分字节值和 8 位无符号整数值。

`rune` 是 `int32` 的别名，在所有方面都等同于 `int32`。按照惯例，它用于区分字符值和整数值。

`any` 是 `interface{}` 的别名，在所有方面都等效于 `interface{}`。


>---
### type error

```go
type error interface {
    Error() string
}
```

`error` 内置接口类型是表示错误条件的常规接口，`nil` 值表示没有错误。


>---
### var nil

```go
var nil Type 
```

`nil` 是一个预先声明的标识符，`Type` 必须是指针、通道、函数、接口、映射或切片类型，`nil` 表示这些类型的零值。

```go
var (
	ch      chan<- string
	slice   []byte
	m       map[string]bool
	f       func(int)
	p       *struct { i int; s []byte }
	fsys    fs.FS
)
```

>---
### func append

```go
func append(slice []Type, elems ...Type) []Type
```

`append` 内置函数将元素追加到切片的末尾。如果它有足够的容量，则对目标 `slice` 重新切片以容纳新元素；否则，将分配一个二倍容量长度的底层数组。`append` 返回更新后的切片。因此，有必要存储 `append` 的结果，通常在保存切片本身的变量中：

```go
slice = append(slice, elem1, elem2)
slice = append(slice, anotherSlice...)
```

作为一种特殊情况，在字节片上附加一个字符串是合法的，如下所示：

```go
slice = append([]byte("hello "), "world!"...)
```

>---
### func cap

```go
func cap(v Type) int
```

`cap` 内置函数根据类型返回 `v` 的容量：
- 数组：返回数组的元素数目，等效于 `len(v)`。
- 指向数组的指针：返回指向的数组的元素数目，等效于 `len(v)`。
- 切片：返回切片能够到达的最大长度。`cap(nil) == 0`。
- 通道：返回通道缓冲区容量，以元素的类型为单位。`cap(nil) == 0`。

对于某些参数（如简单数组表达式），结果可以是常量。

>---
### func clear

```go
func clear[T ~[]Type | ~map[Type]Type1](t T)
```

`clear` 内置函数用于清空映射和切片。对于映射，`clear` 将删除所有键值对，从而生成空映射。对于切片，`clear` 将切片长度以内的所有元素设置为相应元素类型的零值。如果参数类型是类型参数，则类型参数的类型集必须仅包含映射或切片类型，并且 `clear` 执行类型参数所隐含的操作。

>---
### func close

```go
func close(c chan<- Type)
```

`close` 内置函数关闭一个通道，该通道必须是双向的或只发送的。它应该只由发送方执行，而不是由接收方执行，并且在接收到最后一个发送的值后关闭通道。在从关闭的通道 `c` 接收到最后一个值之后，任何来自 `c` 的接收都将成功而不会阻塞，并为通道元素返回零值。形式 `x, ok := <- c` 也会为关闭的空通道设置 `ok` 为 `false`。

>---
### complex

```go
func complex(r, i FloatType) ComplexType
```

`complex` 内置函数从两个浮点值构造一个复数值。实部和虚部必须具有相同的大小，即均为 `float32` 或 `float64`（或可分配给它们），并且返回值将是相应的复数类型（`float32` 为 `complex64`，`float64` 为 `complex128`）。

>---
### func copy

```go
func copy(dst, src []Type) int
```

`copy` 内置函数将元素从源切片复制到目标切片。在特殊情况下，它也可以将字节从字符串复制到字节片。`src` 和 `dst` 可能重叠。`copy` 返回复制的元素数，它将是 `len(src)` 和 `len(dst)` 中的最小值。

>---
### func delete

```go
func delete(m map[Type]Type1, key Type)
```

`delete` 内置函数从 `m` 中删除具有指定键（`m[key]`）的元素。如果 `m` 为 `nil` 或者没有这样的元素，则 `delete` 不进行任何操作。

>---
### func imag

```go
func imag(c ComplexType) FloatType1
```

`imag` 内置函数返回复数 `c` 的虚部。返回值将是与 `c` 类型对应的浮点类型。

>---
### func len

```go
func len(v Type) int
```

`len` 内置函数根据类型返回 `v` 的长度：

- 数组：返回数组的元素个数。
- 指向数组的指针：返回指向的数组的元素个数。
- 切片或映射：返回 `v` 的元素个数。若 `v == nil` 则返回 0。
- 字符串：返回字符串的字节数目。
- 通道：返回通道缓冲区中排队（未读）的元素数目。若 `v == nil` 则返回 0。

对于某些参数（如字符串文字或简单数组表达式），结果可以是常量。

>---
### func make

```go
func make(t Type, size ...IntegerType) Type
```

`make` 内置函数分配和复制 slice、map 或 chan 类型的对象（仅限）。和 `new` 一样，第一个参数是类型，而不是值。与 `new` 不同的是，`make` 的返回类型与其参数的类型相同，而不是返回指向它的指针。结果的指定取决于类型：

- 切片：返回 `size` 长度和容量的切片；可以提供第二个参数指定不同的容量 `make(slice, size, capacity)`，容量至少和长度一样大。
- 映射：返回可选容量 `size` 起始大小的映射，添加元素将扩容。
- 通道：返回可选缓冲区大小的通道，若忽略大小，通道为无缓冲。

>---
### func max

```go
func max[T cmp.Ordered](x T, y ...T) T
```

`max` 内置函数返回 `cmp.Ordered` 类型的固定数量参数中的最大值。必须至少有一个参数。如果 `T` 是浮点类型，并且任何参数都是 *NaN*，则 `max` 将返回 *NaN*。


>---
### func min

```go
func min[T cmp.Ordered](x T, y ...T) T
```

`min` 内置函数返回 `cmp.Ordered` 类型的固定数量参数中的最小值。必须至少有一个参数。如果 `T` 是浮点类型，并且任何参数都是 *NaN*，则 `min` 将返回 *NaN*。

>---
### func new

```go
func new(Type) *Type
```

`new` 内置函数分配内存。第一个参数是一个类型，而不是一个值，返回的值是一个指向该类型的新分配的零值的指针。

>---
### func panic

```go
func panic(v any)
```

内置函数 `panic` 停止当前 *goroutine* 的正常执行。当函数 `F` 调用 `panic` 时，`F` 的正常执行立即停止。任何被 `F` 延迟执行的函数都以常规方式运行，然后 `F` 返回给它的调用方。对于调用方 `G`，调用 `F` 的行为就像调用 `panic`，即终止 `G` 的执行并运行任何延迟的函数。这过程将一致持续到 *goroutine* 中的所有函数以相反的顺序执行终止。此时，程序将以非零退出代码终止。这种终止序列称为 *panicking*，可以通过内置函数 *recover* 控制。

从 Go 1.21 开始，使用 `nil` 接口值或非类型化 `nil` 调用 `panic` 会导致 *run-time error*（另一种 `panic`）。GODEBUG 设置 `panicnil=1` 禁用 *run-time error*。

>---
### func print

```go
func print(args ...Type)
```

`print` 内置函数以特定于实现的方式格式化其参数，并将结果写入标准错误。`print` 对于引导和调试很有用；无法确保它能保留在语言中。

>---
### func println

```go
func println(args ...Type)
```

`println` 内置函数以特定于实现的方式格式化其参数，并将结果写入标准错误。总是在参数之间添加空格，并追加一个换行符。`println` 对于引导和调试很有用；无法确保它能保留在语言中。

>---
### func real
 
```go
func real(c ComplexType) FloatType
```

`real` 内置函数返回复数 `c` 的实部。返回值将是与 `c` 类型对应的浮点类型。


>---
### func recover

```go
func recover() any
```

`recover` 内置函数允许程序管理 *panicking* *goroutine* 的行为。在延迟函数（但不是它调用的任何函数）内调用执行 `recover`，以通过恢复正常执行来停止 *panic* 序列，并检索传递给调用 `panic` 的 `error` 值。如果在延迟函数之外调用 `recover`，它将不会停止 *panic* 序列。在这种情况下，或者当 *goroutine* 没有 *panicking* 时，`recover` 返回 `nil`。

---