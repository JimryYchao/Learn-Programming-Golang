<div id="top" style="z-index:99999999;position:fixed;bottom:35px;right:50px;float:right">
	<a href="./code/bufio_test.go" target="_blank"><img id="img-code" src="../_rsc/to-code.drawio.png" ></img></a>
	<a href="#TOP" ><img id="img-top" src="../_rsc/to-top.drawio.png" ></img></a>
	<a href="..\README.md"><img id="img-back" src="../_rsc/back.drawio.png"></img></a>
</div>
<a id="TOP"></a>

## Package [bufio](https://pkg.go.dev/bufio)

`bufio` 包实现缓冲 I/O。它包装了一个 `io.Reader` 或 `io.Writer` 对象，创建了另一个对象（`Reader` 或 `Writer`），该对象也实现了该接口，但提供了缓冲和一些文本 I/O 帮助。

---
### const MaxScanTokenSize

```go
const MaxScanTokenSize = 64 * 1024
```

`MaxScanTokenSize` 是用于缓冲 *tokens* 的最大大小，除非用户使用 `Scanner.Buffer` 提供显式缓冲区。实际的 `MaxScanTokenSize` 可能更小，因为缓冲区可能需要包括，例如换行符。

>---
### var bufio.Err

```go
var (
    ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
    ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
    ErrBufferFull        = errors.New("bufio: buffer full")
    ErrNegativeCount     = errors.New("bufio: negative count")
)
```

>---
### var Scanner.Err

由 `Scanner` 返回的错误：

```go
var (
    ErrTooLong         = errors.New("bufio.Scanner: token too long")
    ErrNegativeAdvance = errors.New("bufio.Scanner: SplitFunc returns negative advance count")
    ErrAdvanceTooFar   = errors.New("bufio.Scanner: SplitFunc returns advance count beyond input")
    ErrBadReadCount    = errors.New("bufio.Scanner: Read returned impossible count")
)
```

>---
### var ErrFinalToken

```go
var ErrFinalToken = errors.New("final token")
```

`ErrFinalToken` 是一个特殊的 *哨兵* 错误值。它应由 `Scanner.Split` 函数返回，以指示扫描应该无错误地停止。如果传递的错误 *token* 不为 `nil`，则该 *token* 为最后一个标记。

该值对于提前停止处理或需要交付最终空 *token*（不同于 `nil` 标记）时很有用。可以通过自定义错误值实现相同的行为，但在这里提供一个更整洁的值。

>---
### type Scanner

```go
type Scanner struct {
    // contains filtered or unexported fields
}
```

`Scanner` 为读取数据（如由换行符分隔的文本行组成的文件）提供了方便的接口。连续调用 `Scanner.Scan` 方法将遍历文件的 “*token*”，跳过 *token* 之间的字节。*token* 的规范由 `SplitFunc` 类型的 `split` 函数定义；默认的 `split` 函数将输入分成行，并弃掉行终止符。`bufio` 包中定义了 `Scan` 拆分函数，用于将文件扫描为行、字节、UTF-8 编码的字符和空白符分隔的单词。客户端可以提供自定义拆分函数。

`Scanner` 在出现错误、第一个 I/O 错误或 *token* 太大而无法装入 `Scanner.Buffer` 提供的缓冲时（默认为 `MaxScanTokenSize` 大小）会不可恢复地停止。当扫描停止时，读取器可能已经任意地推进超过最后一个 *token*。需要对错误处理或大 *token* 进行更多控制的程序，或者必须在读取器上运行顺序扫描的程序，应该使用 `bufio.Reader`。

>---
#### func NewScanner

```go
func NewScanner(r io.Reader) *Scanner
```

`NewScanner` 返回一个新的从 `r` 读取的 `Scanner`。`Scanner` 的 `split` 函数默认为 `bufio.ScanLines`。

#### methods of Scanner

```go
func (s *Scanner) Buffer(buf []byte, max int)
func (s *Scanner) Bytes() []byte
func (s *Scanner) Err() error
func (s *Scanner) Scan() bool
func (s *Scanner) Split(split SplitFunc)
func (s *Scanner) Text() string
```

`Scanner.Buffer` 设置扫描时使用的初始缓冲区以及扫描期间可能分配的最大缓冲区大小。最大标记大小必须小于 `max` 和 `cap(buf)` 中的较大值。如果 `max <= cap(buf)`，`Scanner.Scan` 将只使用这个缓冲区，且不再做任何新的分配。默认情况下，`Scanner.Scan` 使用内部缓冲区并将最大标记大小设置为 `MaxScanTokenSize`。若在扫描开始后调用缓冲区，则会产生 *panic*。

`Scanner.Bytes` 返回通过调用 `Scanner.Scan` 生成的最新标记。底层数组可能指向将被后续对 `Scan` 调用的覆盖的数据。它不做分配。

`Scanner.Err` 返回扫描程序遇到的第一个非 `EOF` 错误。

`Scanner.Scan` 使 `Scanner` 前进到下一个标记，然后可以通过 `Scanner.token` 或 `Scanner.Text` 方法使用该标记。当没有更多的标记时，它返回 `false`，表示到达输入的结尾，或者是出现错误。`Scan` 返回 `false` 后，`Scanner.Err` 方法将返回扫描期间发生的任何错误，但如果是 `io.EOF`，`Scanner.Err` 将返回 `nil`。如果 `split` 函数返回太多的空标记而没有推进输入，则扫描将发生异常。这是 `Scanner` 的常见错误模式。

`Scanner.Split` 设置 `Scanner` 使用的 `split` 拆分函数。默认的 `split` 函数为 `bufio.ScanLines`。如果在扫描开始后调用 `Scanner.Split` ，则会发生 *panic*。

`Scanner.Text` 返回调用 `Scanner.Scan` 生成的最新 *token*，返回新分配的字符串并保存其字节。

```go
func main() {
	file, _ := os.Open("_02_Go Library/Go Library summary/bufio/ref/readline.file")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 512), 512)
	scanner.Split(bufio.ScanLines) // 行扫描，默认行为
	for scanner.Scan() {
		fmt.Printf("Read line : %q\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
```

>---
#### func ScanBytes

```go
func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
```

`ScanBytes` 是 `Scanner` 的 `split` 可以使用的一个拆分函数，它将每个字节作为 *token* 返回。

```go

```

>---
#### func ScanLines

```go
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
```

`ScanLines` 是 `Scanner` 的 `split` 可以使用的一个拆分函数，它返回每一行文本，去掉任何尾随的行尾标记。返回的行可能为空 `""`。行尾标记是一个可选的回车符，后跟一个强制性的换行符。在正则表达式表示法中，它是 `\r?\n`。输入的最后一个非空行将被返回，即使它没有换行符。

```go
// 读取标准输入
func ScanLines() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines) // 设置行拆分
	fmt.Println(`Start ScanLines in STDIN, input "^Z" to stop scanning.`)
	for scanner.Scan() {
		fmt.Printf("read a line : %q\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
```

>---
#### func ScanRunes

```go
func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
```

`ScanRunes` 是 `Scanner` 的 `split` 可以使用的一个拆分函数，它将每个 UTF-8 编码的字符码位作为 *token* 返回。返回的码位序列等价于对输入字符串的范围循环，这意味着错误的 UTF-8 编码转换为 U+FFFD = `"\xef\xbf\xbd"`。由于 *Scan* 接口，这使得客户端无法区分正确编码的替换字符和编码错误。

```go
// 读取一段输入中的英文字符
func ScanRunes() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)
	fmt.Println(`Start ScanRunes in STDIN, input "^Z" to stop scanning.`)
	for scanner.Scan() {
		rangeTable := []*unicode.RangeTable{unicode.Lu, unicode.Ll}
		if l := scanner.Bytes(); unicode.IsOneOf(rangeTable, bytes.Runes(l)[0]) {
			fmt.Printf("read a letter : %q\n", string(l))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
```

>---
#### func ScanWords
```go
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
```

`ScanWords` 是 `Scanner` 的 `split` 可以使用的一个拆分函数，它返回每个空白分隔的文本单词，删除周围的空格。它永远不会返回空字符串。空白的定义由 `unicode.IsSpace` 设置。

```go
// 读取标准输入中的数字
func ScanWords() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	fmt.Println(`Start ScanWords in STDIN, input "^Z" to stop scanning.`)
	for scanner.Scan() {
		s := scanner.Text()
		if num, ok := func(s string) (string, bool) {
			_, err := strconv.ParseFloat(s, 64)
			return s, err == nil
		}(s); ok {
			fmt.Printf("read a number : %q\n", num)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
```

>---
#### funcType SplitFunc

```go
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
```

`SplitFunc` 是用于对输入的 *token* 进行拆分的 `split` 函数的签名。参数是剩余未处理数据的初始子字符串 `data` 和一个标志 `atEOF`，它报告读取器是否没有更多的数据要可提供。返回值是用于推进输入的字节数 `advance` 和返回给用户的下一个  `token`（如果有），以及一个 `err`（如果有）。

如果函数返回错误，则扫描停止，在这种情况下，某些输入可能会被丢弃。如果该错误是 `ErrFinalToken`，则扫描无错误地停止。使用 `ErrFinalToken` 传递的非 `nil` *token* 将是最后一个 *token*，使用 `ErrFinalToken` 的 `nil` *token* 将立即停止扫描。

否则，`Scanner` 将推进输入。如果 `token` 不为 `nil`，`Scanner` 会将其返回给用户。如果 `token` 为 `nil`，`Scanner` 读取更多数据并继续扫描；如果没有更多数据，且如果 `atEOF` 为 `true`，则 `Scanner` 返回。如果数据还没有包含完整的 *token*，例如，如果它在扫描行时没有换行符，`SplitFunc` 可以返回 `(0，nil，nil)` 来通知 `Scanner` 将更多数据读入切片，并从输入中的同一点开始重新尝试一个更长的切片。

该函数永远不会在空数据切片上调用，除非 `atEOF` 为 `true`。如果 `atEOF` 为 `true`， `data` 也可能是非空的，并且始终包含未处理的文本。

```go

```


>---
### type Reader

>---
#### func NewReader

>---
#### func NewReaderSize

