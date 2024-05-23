<a id="TOP"></a>

## Package fmt

<div id="top" style="z-index:99999999;position:fixed;bottom:35px;right:50px;float:right">
	<a href="./code/fmt_test.go" target="_blank"><img id="img-code" src="../_rsc/to-code.drawio.png" ></img></a>
	<a href="#TOP" ><img id="img-top" src="../_rsc/to-top.drawio.png" ></img></a>	
	<a href="https://pkg.go.dev/fmt" ><img id="img-link" src="../_rsc/to-link.drawio.png" ></img></a>
	<a href="..\README.md"><img id="img-back" src="../_rsc/back.drawio.png"></img></a>
</div>

包 `fmt` 使用类似于 C 的 `printf` 和 `scanf` 的函数实现格式化 I/O。“*verbs*” 格式是从 C 派生的。

---
### verbs

```go
%v		// value 的默认格式; 对于 struct, %+v 添加字段名
%#v		// value 的 go 语法表示
%T		// type of value 的 go 语法表示
%% 		// %

// bool 
%t		// boolean 的字面值
// integer
%b		// 二进制整数
%c		// Unicode code 对应的字符
%d		// 十进制整数
%o		// 八进制整数
%O		// 0o 前缀的八进制整数
%q		// 安全转义的单引号字符，例如 65 转义为 'A'，无效整数码位改为 U+FFFD
%x		// 十六进制整数（小写）
%X		// 十六进制整数（大写）
%U		// 如 Unicode U+1234 形式 
// floating-point & complex
%b		// 指数为 2 的幂的无小数计数法, -123p-45
%e		// e 计数法 -1.23e+45
%E		// E 计数法 -1.23E+45
%f,%F	// 无指数的浮点表示 123.456
%g,%G	// 根据数值大小选择 %e/%E 还是 %f/%F
%x,%X	// 十六进制, 指数为 2 的幂, -0x1.23abc+20 / -0X1.23ABCP+20

// strings & []byte
%s		// 字符串或字节切片的无解释字节
%q		// 安全转义的双引号字符串
%x,%X	// 逆向转义为 \xhh，每个字节两个字符，例如 "我" > e68891

// pointers
%p		// 十六进制地址值，切片返回 &s[0] 的地址，指针用于整数 verbs 时被视为一般整数；例如 0xc0000dc4e0
```

对于 `%v` 的默认格式为：

```go
bool					%v >> %t
int,int8,...			%v >> %d
uint,uint8,...			%v >> %d, %#v >> %#x
float32,complex64,...	%v >> %g
string					%s
chan					%p
pointer					%p
struct					{field0 field1 ...}
array,slice				[elem0 elem1 ...]
map						map[k1:v1 k2:v2 ...]
compound pointers		&{}, &[], &map[] 
```

宽度和精度由 $width.precision$ 表示，宽度表示要输出的最小字符（`rune`）数，必要时空格填充。
- 对于 `string` 和 `bytestring`，精度限制了格式化输入的长度，以 `rune` 为单位；`%x, %X` 以字节为单位。
- 浮点的精度默认为 6。
- 对于复数，宽度和精度独立应用到它的两个分量。

```go
%f			// 默认宽度，精度
%9f			// 宽度 9，默认精度
%.2f		// 默认宽度，精度 2
%9.2f		// 宽度 9，精度 2	
%9.f		// 宽度 9，精度 0	
```

一些标志位：

```go
'+'			// 对于数值打印符号位；%+q 保证至输出 ASCII
'-'			// 空格填充在右侧
'#'			/* %#b,%#o,%#x,%#X		打印前导 0b,0,0x,0X
			   %#p					取消打印前导 0x		
			   %#q					打印原始字符串（若支持）
			   %#f,%#e,%#g			打印小数点，不删除 %g 的零 
			   %#U					U+0078 'x', 同时打印对应的可打印字符
			*/
' '			//  % d 整数预留符号位空格；% x 字节之间放置空格
'0'			// 0 代替空格作为宽度的填充，对于字符串类型将忽略
```

`[n]` 表示显示参数索引，`("%[2]d %[1]d\n", 11, 22)` 将产生 `"22 11"`。显式索引会影响后续的 *verbs*，后续的索引依次递增，`("%d %d %#[1]x %#x", 16, 17)` 将产生 `"16 17 0x10 0x11"`

如果操作数是接口值，则使用内部具体值。除了 `%T, %p`，实现某些接口的操作数应用一些特殊规则：
1. 如果操作数是一个 `reflect.Value`，则操作数将被它所保存的具体值替换，并且打印将继续执行下一个规则。 
2. 如果一个操作数实现了 `Formatter` 接口，它将被调用，*verbs* 和标志的解释由该实现控制。
3. 如果使用 `%#v`，并且操作数实现了 `GoStringer` 接口，则将调用该接口。如果格式（对于 `Println` 等，隐式为 `%v`）对于字符串（`%s %q %x %X`）有效，或者是 `%v` 但不是 `%#v`，则适用以下两个规则：
4. 如果一个操作数实现了 `error` 接口，则调用 `Error` 以将对象转换为字符串，然后根据 *verb*（如果有的话）的要求进行格式化。
5. 如果一个操作数实现了方法 `String() string`，则调用 `String()` 以将对象转换为字符串，然后根据 *verb*（如果有）的要求进行格式化。


---
