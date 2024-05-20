<div id="top" style="z-index:99999999;position:fixed;bottom:35px;right:50px;float:right">
	<a href="./code/io_test.go" target="_blank"><img id="img-code" src="../_rsc/to-code.drawio.png" ></img></a>
	<a href="#TOP" ><img id="img-top" src="../_rsc/to-top.drawio.png" ></img></a>
	<a href="..\README.md"><img id="img-back" src="../_rsc/back.drawio.png"></img></a>
</div>
<a id="TOP"></a>

## Package [io](https://pkg.go.dev/io)

包 `io` 提供 I/O 原语的基本接口。它的主要工作是将这些原语的现有实现（如包操作系统中的原语）包装到共享的公共接口中，这些公共接口抽象了功能，并加上一些其他相关的原语。

---
### Readers

#### io.Reader 





### IO 错误定义

```go
var (
	// 当没有更多输入时，Read 返回的错误
	EOF = errors.New("EOF")		
	// 对封闭管道进行读或写操作的错误
	ErrClosedPipe = errors.New("io: read/write on closed pipe")
	// 当许多对 Read 的调用都未能返回任何数据或错误时，代表失败的标志
    ErrNoProgress = errors.New("multiple Read calls return no data or error")
	// 意味着读取使用的缓冲区过小
	ErrShortBuffer = errors.New("short buffer")
	// 表示写入接受的字节数少于请求的字节数
	ErrShortWrite = errors.New("short write")
	// 意味着在读取固定大小的块或数据结构时遇到错误
	ErrUnexpectedEOF = errors.New("unexpected EOF")
)
```

>---
### 基础接口定义

#### Writers

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
var Discard Writer = discard{}  // 无作用的 Writer，调用不起作用
```

`Writer` 是包装基本 `Write` 方法的接口。`Write` 将 `len(p)` 个字节写入 `p` 的底层数据流。并返回成功写入的字节数（`n` 小于 `len(p)`，实现应返回一个非空错误）。`p` 不可修改且实现不能保留。

```go
type WriterAt interface {
	WriteAt(p []byte, off int64) (n int, err error)
}
```

`WriterAt` 是包装基本 `WriteAt` 方法的接口。`WriteAt` 将 `len(p)` 个字节写入 `p` 写入偏移量为 `off` 的底层数据流。它返回从 `p`（`0 <= n <= len(p)`）写入的字节数以及导致写入提前停止的任何错误。如果 `WriteAt` 返回 `n < len(p)`，则它必须返回一个非 `nil` 错误。

如果 `WriteAt` 正在写入具有 *Seek* 偏移的目标，则 `WriteAt` 不应影响底层 *Seek* 偏移受到其影响。

如果范围不重叠，`WriteAt` 的客户端可以在同一目标上执行并行 `WriteAt` 调用。实现不能保留 `p`。







```go
type WriterTo interface {
	WriteTo(w Writer) (n int64, err error)
}
```


```go
type ByteWriter interface {
	WriteByte(c byte) error
}
```


#### Writers


#### Seeker


#### Closer

```go
type Closer interface{

}
```


### const SeekXXX

```go
const (
	SeekStart   = 0 // 查找相对于文件开头的位置
	SeekCurrent = 1 // 查找相对于当前偏移量的位置
	SeekEnd     = 2 // 查找相对于文件末尾的位置
)
```


>---


>---
### interface Reader

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

`Reader` 是包装基本 `Read` 方法的接口。

`Read` 将 `len(p)` 字节读入 `p`。它返回读取的字节数（`0 <= n <= len(p)`）和遇到的任何错误。即使 `Read` 返回 `n < len(p)`，它也可能在调用期间使用所有 `p` 作为暂存空间。如果某些数据可用，但不是 `len(p)` 字节，`Read` 通常会返回可用的数据，而不是等待更多的数据。

当 `Read` 在成功读取 `n > 0` 个字节后遇到错误或文件结束条件时，它将返回读取的字节数。它可以从同一个调用返回（非空）错误，或者从后续调用返回错误（并且 `n == 0`）。这种一般情况的一个例子是，在输入流的末尾返回非零字节数的 `Reader` 可能返回 `err == EOF` 或 `err == nil`。下一次读取应该返回 0，`EOF`。

调用方应该在考虑错误 `err` 之前处理返回的 `n > 0` 字节。这样做可以正确处理在读取某些字节后发生的 I/O 错误以及允许的 `EOF` 行为。

如果 `len(p) == 0`，`Read` 应总是返回 `n == 0`。如果已知某些错误条件，例如 `EOF`，则它可能返回非空错误。

`Read` 的实现不鼓励返回零字节计数和 `nil` 错误，除非 `len(p) == 0`。调用方应该将 0 和 `nil` 的返回值视为表示没有发生任何事情；特别是它并不表示 `EOF`。

`Read` 的实现实现不能保留 `P`。



>---
### interface ReaderAt

```go
type ReaderAt interface {
    ReadAt(p []byte, off int64) (n int, err error)
}
```

`ReaderAt` 是包装基本 `ReadAt` 方法的接口。

`ReadAt` 从底层输入源中的偏移 `off` 开始将 `len(p)` 字节读入 `p`。它返回读取的字节数（`0 <= n <= len(p)`）和遇到的任何错误。

当 `ReadAt` 返回 `n < len(p)` 时，它返回一个非空的错误，并解释为什么没有返回更多字节。在这方面，`ReadAt` 比 `Read` 更严格。

即使 `ReadAt` 返回 `n < len(p)`，它也可能在调用期间使用所有 `p` 作为暂存空间。如果某些数据可用，但 `len(p)` 字节不可用，`ReadAt` 将阻塞，直到所有数据可用或发生错误。在这方面，`ReadAt` 不同于 `Read`。

如果 `ReadAt` 返回的 `n = len(p)` 字节位于输入源的末尾，则 `ReadAt` 可能返回 `err == nil` 或 `err == EOF`。

如果 `ReadAt` 从具有 *seek* 偏移的输入源进行读取，则 `ReadAt` 不应影响底层 *seek* 偏移或受到其影响。

`ReadAt` 的客户端可以在同一输入源上执行并行 `ReadAt` 调用。实现不能保留 `p`。

>---
### interface WriterTo

```go
type WriterTo interface {
    WriteTo(w Writer) (n int64, err error)
}
```

`WriterTo` 是包装 `WriteTo` 方法的接口。

`WriteTo` 将数据写入 `w`，直到没有更多的数据可写或发生错误。返回值 `n` 是写入的字节数。在写入过程中遇到的任何错误也会返回。`Copy` 函数使用 `WriterTo`（如果可用）。

>---
### interface ReaderFrom

```go
type ReaderFrom interface {
    ReadFrom(r Reader) (n int64, err error)
}
```

`ReaderFrom` 是包装 `ReadFrom` 方法的接口。`ReadFrom` 从 `r` 读取数据，直到出现 `EOF` 或错误。返回值 `n` 是读取的字节数。在读取期间遇到的除 `EOF` 之外的任何错误都将返回。`Copy` 函数使用 `ReaderFrom`（如果可用）。

>---
### interface RuneReader

```go
type RuneReader interface {
    ReadRune() (r rune, size int, err error)
}
```

`RuneReader` 是包装 `ReadRune` 方法的接口。`ReadRune` 读取单个编码的 Unicode 字符，并返回该字符及其字节大小。如果没有可用的字符，则会设置 err。

>---
### interface RuneScanner 

```go
type RuneScanner interface {
    RuneReader
    UnreadRune() error
}
```

`UnreadRune` 导致下一次调用 `ReadRune` 返回最后一次读取的字符。如果最后一个操作不是对 `ReadRune` 的成功调用，`UnreadRune` 可能会返回一个错误，取消最后一个读取的字符（或最后一个未读取的字符之前的字符），或（在支持 `Seeker` 接口的实现中）查找当前偏移之前的字符的开头。

>---
### interface Closer

```go
type Closer interface {
    Close() error
}
```

`Close` 操作之后对 `Closer` 的调用的行为是未定义的。特定的实现可以记录它们自己的行为。

>---
### interface Seeker

```go
type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}
```

`Seeker` 是封装基本 `Seek` 方法的接口。`Seek` 将下一次读取或写入的偏移量设置为 `offset`。`SeekStart` 表示相对于文件的开始，`SeekCurrent` 表示相对于当前偏移量，`SeekEnd` 表示相对于结束（例如，`offset = -2` 指定文件的倒数第二个字节）。`Seek` 返回相对于文件开头的新偏移量或错误（如果有）。

在文件开始之前查找偏移量是错误的。可以允许查找任何正偏移量，但如果新偏移量超过底层对象的大小，则后续 I/O 操作的行为与实现相关。

>---
### interface ByteWriter

```go
type ByteWriter interface {
    WriteByte(c byte) error
}
```

`ByteWriter` 是包装 `WriteByte` 方法的接口。

>---
### interface ByteReader

```go
type ByteReader interface {
    ReadByte() (byte, error)
}
```

`ReadByte` 读取并返回输入中的下一个字节或遇到的任何错误。


>---
### interface ByteScanner

```go
type ByteScanner interface {
    ByteReader
    UnreadByte() error
}
```

`UnreadByte` 导致下一次调用 `ReadByte` 返回最后读取的字节。如果最后一个操作不是对 `ReadByte` 的成功调用，则 `UnreadByte` 可能返回错误，未读最后一个字节读取（或最后未读字节之前的字节），或（在支持 `Seeker` 接口的实现中）查找当前偏移量之前的一个字节。

>---
### interface StringWriter

```go
type StringWriter interface {
    WriteString(s string) (n int, err error)
}
```

>---
### interface ReadCloser

```go
type ReadCloser interface {
    Reader
    Closer
}
```

>---

#### func NopCloser

```go
func NopCloser(r Reader) ReadCloser
```

`NopCloser` 返回一个 `ReadCloser`，它带有一个无操作的 `Close` 方法，包装了提供的 `Reader r`。如果 `r` 实现了 `WriterTo`，则返回的 `ReadCloser` 将通过转发对 `r` 的调用来实现 `WriterTo`。

>---
### interface ReadSeekCloser

```go
type ReadSeekCloser interface {
    Reader
    Seeker
    Closer
}
```

>---
### interface ReadSeeker

```go
type ReadSeeker interface {
    Reader
    Seeker
}
```

>---
### interface ReadWriteCloser 

```go
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

>---
### interface ReadWriteSeeker

```go
type ReadWriteSeeker interface {
    Reader
    Writer
    Seeker
}
```

>---
### interface ReadWriter 

```go
type ReadWriter interface {
    Reader
    Writer
}
```

>---
### interface WriteCloser

```go
type WriteCloser interface {
    Writer
    Closer
}
```

>---
### interface WriteSeeker 

```go
type WriteSeeker interface {
    Writer
    Seeker
}
```

>---
### func MultiWriter

```go
func MultiWriter(writers ...Writer) Writer
```

`MultiWriter` 创建一个 `Writer`，将其写入复制到所有提供的 `writers`，类似于 Unix `tee(1)` 命令。每次写操作都写入到列出的写入器 `writers`，每次一个。如果列出的写入器返回错误，则整个写入操作停止并返回错误，并且它不会再沿着列表继续。

```go
func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	var buf1, buf2 strings.Builder
	w := io.MultiWriter(&buf1, &buf2)

	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}

	fmt.Print(buf1.String())
	fmt.Print(buf2.String())
}
```

>---
### func MultiReader

```go
func MultiReader(readers ...Reader) Reader
```

`MultiReader` 返回一个 `Reader`，它是所提供的输入读取器的逻辑串联。它们是按顺序读取的。一旦所有的输入都返回 `EOF`，`Read` 将返回 `EOF`。如果任何读取器返回一个非 `nil`、非 `EOF` 的错误，`Read` 将返回该错误。

```go
func main() {
	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")
	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
	// first reader second reader third reader
}
```

>---
### func TeeReader

```go
func TeeReader(r Reader, w Writer) Reader
```

`TeeReader` 返回一个 `Reader`，它将从 `r` 读取的内容写入 `w`。通过它执行的所有从 `r` 的读取都与对 `w` 的相应写入相匹配。无内部缓冲；写操作必须在读操作完成之前完成。写入时遇到的任何错误都报告为读取错误。

```go
func main() {
	var r io.Reader = strings.NewReader("some io.Reader stream to be read\n")
	r = io.TeeReader(r, os.Stdout)

	// Everything read from r will be copied to stdout.
	if _, err := io.ReadAll(r); err != nil {
		log.Fatal(err)
	}
}
```



>---
### func Copy

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

`Copy` 将副本从 `src` 复制到 `dst`，直到 `src` 上出现错误。它返回复制的字节数和复制时遇到的第一个错误（如果有）。

如果 `src` 实现 `WriterTo`，则通过调用 `src.WriteTo(dst)` 实现复制。否则，如果 `dst` 实现 `ReaderFrom`，则通过调用 `dst.ReadFrom(src)` 实现复制。

```go
func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	if count, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Copy Successful, And read %v bytes\n", count)
	}
}
```

>---
### func CopyBuffer

```go
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
```

`CopyBuffer` 与 `Copy` 相同，区别在于它使用用户提供的缓冲区而不是分配一个临时的缓冲区，若 `buf` 为 `nil`，则临时分配一个，若它的长度为零，则引发 *Panic*。

如果 `src` 实现了 `WriterTo` 或 `dst` 实现了 `ReaderFrom`，则 `buf` 将不会用于执行复制。

```go
func main() {
	r1 := strings.NewReader("first buf reader\n")
	r2 := strings.NewReader("second buf reader\n")
	r3 := strings.NewReader("nil reader\n")
	buf := make([]byte, 8)

	// buf is used here...
	if c, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Copy Successful, And read %v bytes\n", c)
	}

	// reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}

	// buf == nil
	if _, err := io.CopyBuffer(os.Stdout, r3, nil); err != nil {
		log.Fatal(err)
	}
}
```

>---
### func CopyN

```go
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
```

`CopyN` 将 `n` 个字节从 `src` 复制到 `dst`（或直到出现错误）。它返回复制的字节数和复制时遇到的最早错误。持续读取将在上一次 *Seek* 位置继续，读取结束返回 `EOF`。如果 `dst` 实现 `ReaderFrom`，则使用它实现复制。

```go
func main() {
	r := strings.NewReader("some io.Reader stream to be read")
	var err error
	for err == nil {
		if _, err = io.CopyN(os.Stdout, r, 4); err != nil {
			println()
			if err == io.EOF {
				log.Fatal("READ to EOF")
			}
			log.Fatal(err) // EOF
		}
	}
}
```

>---
### func Pipe

```go
func Pipe() (*PipeReader, *PipeWriter)
```

`Pipe` 创建同步内存管道。它可用于连接期望 `io.Reader` 的代码和期望 `io. Writer` 的代码。

管道上的读取和写入是一对一匹配的，除非需要多个读取来消耗单个写入。也就是说，每次写入 `PipeWriter` 都会阻塞，直到它满足了来自 `PipeReader` 的一个或多个完全消耗写入数据的读取。数据直接从 `Write` 复制到相应的 `Read`（或 `Reads`）；没有内部缓冲。

相互并行调用 `Read` 和 `Write` 或使用 `Close` 调用 `Read` 和 `Write` 是安全的。对 `Read` 的并行调用和对 `Write` 的并行调用也是安全的：各个调用将按顺序进行门控。

```go
func main() {
	r, w := io.Pipe()
	defer r.Close()

	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n") // 写入并复制到 r
		w.Close()                                           // 关闭写入
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}
```

>---

#### type PipeReader

```go
type PipeReader struct {
    // contains filtered or unexported fields
}

func (r *PipeReader) Close() error
func (r *PipeReader) CloseWithError(err error) error
func (r *PipeReader) Read(data []byte) (n int, err error)
```

`PipeReader` 是 `Pipe` 返回的读取端。

`Close` 关闭读取器；管道写入端的后续写操作将返回错误 `ErrClosedPipe`。

`CloseWithError` 关闭读取器；管道写入端的后续写操作将返回错误 `err`。`CloseWithError` 永远不会覆盖前一个错误，如果它存在，并且总是返回 `nil`。

`Read` 实现了标准的 `Read` 接口：它从管道中读取数据，阻塞直到写入器到达或写入端关闭。如果写入端因错误关闭，则 `Read` 返回写入端发生的错误；否则，`err` 为 `EOF`。

>---

#### type PipeWriter

```go
type PipeWriter struct {
    // contains filtered or unexported fields
}

func (w *PipeWriter) Close() error
func (w *PipeWriter) CloseWithError(err error) error
func (w *PipeWriter) Write(data []byte) (n int, err error)
```

`PipeWriter` 是 `Pipe` 的写入端。

`Close` 关闭写入器；管道读取端的后续读操作将不返回字节，而是返回 `EOF`。

`CloseWithError` 关闭写入器；管道读取端的后续读操作将不返回字节，并且返回错误 `err`，或者如果 `err` 为 `nil` 则返回 `EOF`。`CloseWithError` 永远不会覆盖前一个错误，如果它存在，则总是返回 `nil`。

`Write` 实现了标准的 `Write` 接口：它将数据写入管道，阻塞直到一个或多个读取器消耗了所有数据或读取端关闭。如果读取端因错误关闭，则 `Write` 返回读取端发生的错误；否则，`err` 为 `ErrClosedPipe`。


>---
### func ReadAll

```go
func ReadAll(r Reader) ([]byte, error)
```

`ReadAll` 从 `r` 开始读取，直到出现错误或中断，并返回读取的数据。因为 `ReadAll` 被定义为从 `src` 读取直到 `EOF`，所以它不会将从 `Read` 的 `EOF` 视为要报告的错误。

```go
func main() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	if b, err := io.ReadAll(r); err != nil {
		log.Fatal(err) // EOF 不被视为错误
	} else {
		fmt.Printf("%s", b)
	}
}
```

>---
### func ReadAtLeast

```go
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
```

`ReadAtLeast` 从 `r` 读取到 `buf`，至少读取了 `min` 字节，并返回复制的字节数。如果读取的字节较少，则返回错误。仅在无字节读取时返回 `EOF`，若读取字节少于 `min` 时返回 `ErrUnexpectedError`。`min` 大于 `buf` 的长度或 `buf == nil` 时，返回 `ErrShortBuffer`。

```go
func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 14)
	if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf) // some io.Reader

	// buffer smaller than minimal read size.
	if _, err := io.ReadAtLeast(r, make([]byte, 3), 4); err != nil {
		fmt.Println("error:", err) // error: short buffer
	}

	// minimal read size bigger than io.Reader stream
	if _, err := io.ReadAtLeast(r, make([]byte, 64), 64); err != nil {
		fmt.Println("error:", err) // error: unexpected EOF
	}

	// EOF read
	if _, err := io.ReadAtLeast(strings.NewReader(""), buf, 4); err != nil {
		fmt.Println("error:", err) // error: EOF
	}
}
```

>---
### func ReadFull

```go
func ReadFull(r Reader, buf []byte) (n int, err error)
```

`ReadFull` 将 `len(buf)` 字节从 `r` 精确读取到 `buf`。它返回复制的字节数，若读取的字节较少，则返回错误。仅当无字节读取时返回 `EOF`。读取部分字节发生 `EOF`，则返回 `ErrUnexpectedError`。

```go
func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// minimal read size bigger than io.Reader stream
	if _, err := io.ReadFull(r, make([]byte, 64)); err != nil {
		fmt.Println("error:", err)
	}

	// Read EOF
	if _, err := io.ReadFull(strings.NewReader(""), make([]byte, 64)); err != nil {
		fmt.Println("error:", err)
	}
}
```

>---
### func WriteString

```go
func WriteString(w Writer, s string) (n int, err error)
```

`WriteString` 将字符串 `s` 的内容写入 `w`，后者可以接受一个 `[]byte`。如果 `w` 实现了 `StringWriter`，则直接调用 `StringWriter.WriteString`。否则，`Writer.Write` 只被调用一次。

```go
func main() {
	if _, err := io.WriteString(os.Stdout, "Hello World"[0:5]); err != nil {
		log.Fatal(err)
	}
}
```

>---
### func LimitReader

```go
func LimitReader(r Reader, n int64) Reader
```

`LimitReader` 返回一个 `Reader`，它从 `r` 开始读取，但在 `n` 个字节后停止。底层实现是 `*LimitedReader`。

```go
func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	var lr io.Reader = io.LimitReader(r, 4) // 限制只读取 4 字节

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}
```

>---

#### type LimitedReader 

```go
type LimitedReader struct {
    R Reader // underlying reader
    N int64  // max bytes remaining
}

// 接口实现
func (l *LimitedReader) Read(p []byte) (n int, err error)
```

`LimitedReader` 从 `R` 读取数据，但将返回的数据量限制为 `N` 字节。每次调用 `Read` 都会更新 `N` 以反映新的剩余量。当 `N <= 0` 或底层 `R` 返回时，`Read` 返回。


>---
### type OffsetWriter

```go
type OffsetWriter struct {
    // contains filtered or unexported fields
}

// 实现接口
func (o *OffsetWriter) Seek(offset int64, whence int) (int64, error)
func (o *OffsetWriter) Write(p []byte) (n int, err error)
func (o *OffsetWriter) WriteAt(p []byte, off int64) (n int, err error)
```

`OffsetWriter` 将偏移量 `base` 写入映射到底层写入器中的偏移量 `base+off`。

>---

#### func NewOffsetWriter 

```go
func NewOffsetWriter(w WriterAt, off int64) *OffsetWriter
```

`NewOffsetWriter` 返回一个从偏移 `off` 开始写入 `w` 的 `OffsetWriter`。

>---
### type SectionReader

```go
type SectionReader struct {
    // contains filtered or unexported fields
}

func (s *SectionReader) Outer() (r ReaderAt, off int64, n int64)
func (s *SectionReader) Read(p []byte) (n int, err error)
func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)
func (s *SectionReader) Seek(offset int64, whence int) (int64, error)
func (s *SectionReader) Size() int64
```

`SectionReader` 在底层 `ReaderAt` 的节 *section* 上实现 `Read`、`Seek` 和 `ReadAt`。

`Outer` 返回与创建 `SectionReader` 时传递给 `NewSectionReader` 的值。

`Read` 读取字节到给定的字节数组 `p`。

`ReadAt` 从指定 `off` 位置读取字节到 `p`。

在当前的 *section* 上，`Seek` 相对于 `whence` 的位置设置下一次读取的偏移 `offset`。

`Size` 返回实际节 *section* 的大小（以字节为单位）。

```go
func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)

	// SectionReader.Read
	buf := make([]byte, 10)
	if _, err := s.Read(buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf) // io.Reader

	// SectionReader.ReadAt
	buf2 := make([]byte, 6)
	if _, err := s.ReadAt(buf2, 10); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf2) // stream

	// SectionReader.Seek
	_, err := s.Seek(3, io.SeekStart) // 定位到 SectionReader.off + 3
	if err == nil {
		_, err = io.Copy(os.Stdout, s) // Reader stream
	}
	if err != nil {
		log.Fatal(err)
	}

	// SectionReader.Size
	s1 := io.NewSectionReader(r, 5, 100)
	fmt.Println(s.Size())  // 17
	fmt.Println(s1.Size()) // 100
}
```

>---

#### func NewSectionReader

```go
func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader
```

`NewSectionReader` 返回一个 `SectionReader`，它从 `r` 开始读取，从偏移 `off` 开始，在读取 `n` 个字节后或提前到达 `EOF` 时停止。`n < 0` 时表示表示完全读取。

```go
func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 10)

	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}
	// n < 0, 将从 off 处完整读取
	s = io.NewSectionReader(r, 5, -1) // or n = 100 
	println()
	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}
}
// io.Reader
// io.Reader stream to be read
```

---