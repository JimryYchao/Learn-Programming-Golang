## Go 组成元素






### Unsafe 

内置包 `unsafe` 为低级编程提供了便利，包括违反类型系统的操作。使用 `package unsafe` 必须手动检查类型安全性，并且可能不可移植。该包提供以下接口：

```go
package unsafe

type ArbitraryType int  // shorthand for an arbitrary Go type; it is not a real type
type Pointer *ArbitraryType

func Alignof(variable ArbitraryType) uintptr
func Offsetof(selector ArbitraryType) uintptr
func Sizeof(variable ArbitraryType) uintptr

type IntegerType int  // shorthand for an integer type; it is not a real type
func Add(ptr Pointer, len IntegerType) Pointer
func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType
func SliceData(slice []ArbitraryType) *ArbitraryType
func String(ptr *byte, len IntegerType) string
func StringData(str string) *byte
```






















---


---


---


>---
#### 


>---
#### 
>---
#### 
>---
#### 