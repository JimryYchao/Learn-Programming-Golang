package gostd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

type Buffer struct {
	bytes.Buffer
	io.ReaderFrom
	io.WriterTo
}

func beforeTest[TBF testing.TB](t TBF) {
	clear(buf)
	if !testing.Verbose() {
		return
	}
	fmt.Printf(EnterTest, t.Name())
	t.Cleanup(func() {
		fmt.Printf(EndTest, t.Name())
	})
}

var (
	buf         = make([]byte, 512)
	hello       = "Hello World"
	READ        = "Read : `%s`"
	READ_BYTES  = "Read %d Bytes: `%s`"
	WRITE_BYTES = "Write %d Bytes: `%s`"
	EnterTest   = ">>> Enter %s :\n"
	EndTest     = ">>> End   %s\n"
	content     = "some io.Reader stream to be read"
)

type _error *error

func readToStdout(r io.Reader) {
	_, err := io.Copy(os.Stdout, r)
	println()
	CheckErr(err)
}

func writeToStdout(r *bytes.Buffer) {
	logf("%s", r.Bytes())
}

func logf(format string, args ...any) {
	fmt.Printf(format+"\n", args...)
}
func CheckErr(err error) {
	if err == nil {
		return
	}
	logf("ERROR : %s", err)
}
func newReader(s string) io.Reader {
	return strings.NewReader(s)
}
func newWriter(n int) io.Writer {
	return newBytesBuffer(n)
}
func newBytesBuffer(n int) *bytes.Buffer {
	return bytes.NewBuffer(make([]byte, n))
}
func newStdoutWriter() io.Writer {
	return os.Stdout
}

/*
! Reader 包装基本的 `Read` 方法；`Read` 返回读取的字节数
	io.LimitReader		包装一个限制读取字节数的 Reader *LimitedReader
	io.MultiReader		串联一组 Readers，这些 Readers 在内部按顺序 Read
	io.TeeReader		返回一个关联 w 和 r 的 Reader，从 r 读取的内容会相应的写入 w
*/
//? go test -v -run=^TestReader$
func TestReader(t *testing.T) {
	t.Helper()

	t.Run("Reader.Read", func(t *testing.T) {
		beforeTest(t)
		r := newReader("some io.Reader stream to be read")
		if c, err := r.Read(buf); err == nil && c > 0 {
			logf(READ_BYTES, c, buf)
		} else {
			// handle error
		}
		// output: some io.Reader stream to be read
	})

	t.Run("LimitReader", func(t *testing.T) {
		beforeTest(t)
		lr := io.LimitReader(newReader("some io.Reader stream to be read"), 4) // 限制读取字节数
		readToStdout(lr)

		lr2 := io.LimitReader(newReader(hello), -1)
		readToStdout(lr2) // N < 0, lr2.Read return EOF
		// output: some
	})

	t.Run("MultiReader", func(t *testing.T) {
		beforeTest(t)
		r1 := newReader("first reader ")
		r2 := newReader("second reader ")
		r3 := newReader("third reader")
		// 按顺序调用 Reader.Read
		mr := io.MultiReader(r1, r2, r3)
		readToStdout(mr)
		// output: first reader second reader third reader
	})

	t.Run("TeeReader", func(t *testing.T) {
		beforeTest(t)
		r := newReader("some io.Reader stream to be read\n")
		tr := io.TeeReader(r, os.Stdout) // 关联 tr 到 Stdout

		// 从 tr 的任何读取都会复制到 stdout
		if _, err := io.ReadAll(tr); err != nil {
			t.Fatal(err)
		}
	})
}

/*
! Writer 包装基本的 `Write` 方法；`Write` 将最多 `len(p)` 字节写入到底层数据流。
	io.MultiWriter      串联一组 Writer，这些 Writers 在内部按顺序 Write
*/
//? go test -v -run=^TestWriter$
func TestWriter(t *testing.T) {
	t.Helper()

	t.Run("Writer.Write", func(t *testing.T) {
		beforeTest(t)

		//? os.Stdout
		var w io.Writer = newStdoutWriter()
		w.Write([]byte("Writing to os.Stdout\n"))

		//? bytes.Buffer
		var bw *bytes.Buffer = newBytesBuffer(128)
		c, err := bw.Write([]byte("Writing to bytes.Buffer"))
		CheckErr(err)
		logf(WRITE_BYTES, c, bw.Bytes())

		// output:
		// `Writing to os.Stdout`
		// `Writing to bytes.Buffer`
	})

	t.Run("MultiWriter", func(t *testing.T) {
		beforeTest(t)

		w1 := newBytesBuffer(5)
		w2 := &strings.Builder{}
		w3 := os.Stdout

		mw := io.MultiWriter(w1, w2, w3)
		mw.Write([]byte(hello))

		logf("\nbytes.Buffer : %s", w1.Bytes())
		logf("strings.Builder : %s", w2.String())

		// output:
		// Hello World
		// bytes.Buffer : Hello World
		// strings.Builder : Hello World
	})
}

// ! Closer 包装基本的 Close 方法。
/*
! Seeker 包装基本的 Seek 方法；Seek 将下一次读取或写入的偏移量依照 `whence` 设置为 `offset`; `whence` 解释为：
		SeekStart		相对于开始
		SeekEnd 		相对于末尾
		SeekCurrent 	相对于当前偏移量
*/

// ! Copy 将副本从 `src` 复制到 `dst`。它返回复制的字节数和第一个错误（如果有）。
// ! CopyBuffer 等效于 Copy。不能提供 0 长度的 `buf`，传递 `nil` 时将内部创建一个 `buf`。
// ! CopyN 将最多 `n` 个字节从 `src` 复制到 `dst`）。
// ? go test -v -run=^TestCopyFunctions$
func TestCopyFunctions(t *testing.T) {
	t.Helper()
	var wb *bytes.Buffer = newBytesBuffer(128)

	t.Run("Copy", func(t *testing.T) {
		beforeTest(t)

		//? Copy
		wb.Reset() // 从 Reader 到 Writer 复制
		if c, err := io.Copy(wb, newReader(hello)); err != nil {
			t.Fatal(err)
		} else if c > 0 {
			logf(READ_BYTES, c, wb.Bytes())
		}

		//? Copy with negative LimitedReader
		wb.Reset() // N < 0, 将返回 ""
		c, _ := io.Copy(wb, &io.LimitedReader{R: newReader(hello), N: -1})
		logf(READ_BYTES, c, wb.Bytes())

		// output:
		// `Hello World`
		// ``
	})

	t.Run("CopyBuffer", func(t *testing.T) {
		beforeTest(t)
		//? CopyBuffer
		wb.Reset() // `CopyBuffer(dst, src, nil)` same as `Copy(dst, src)`
		if c, err := io.CopyBuffer(wb, newReader(hello), buf); err != nil {
			t.Fatal(err)
		} else {
			logf(READ_BYTES, c, wb.Bytes())
		}

		//? CopyBuffer with empty buffer
		defer func() {
			if err := recover(); err != nil {
				logf("Panicking : %s", err)
			}
		}()
		wb.Reset() // panicking with empty buf
		io.CopyBuffer(wb, newReader(hello), []byte{})

		// output:
		// `Hello World`
		// Panicking : empty buffer in CopyBuffer
	})

	t.Run("CopyN", func(t *testing.T) {
		//? CopyN with small N
		wb.Reset() // len(hello) > 5, 返回 (5, nil)
		if c, err := io.CopyN(wb, newReader(hello), 5); err != nil && err != io.EOF {
			t.Fatal(err)
		} else {
			logf(READ_BYTES, c, wb.Bytes())
		}

		//? CopyN with negative N
		wb.Reset() // N < 0, 返回 (0, nil)
		c, err := io.CopyN(wb, newReader(hello), -1)
		if err != nil {
			logf(io.EOF.Error())
		}
		logf(READ_BYTES, c, wb.Bytes())

		//? CopyN with large N
		wb.Reset() // len(hello) < 100, (len(hello), io.EOF)
		if c, err = io.CopyN(wb, newReader(hello), 100); err != nil {
			if err == io.EOF {
				CheckErr(err)
			} else {
				t.Fatal(err)
			}
		}
		logf(READ_BYTES, c, wb.Bytes())

		// output:
		// `Hello`
		// ``
		// EOF
		// `Hello World`
	})
}

/* PipeReader & PipeWriter
! Pipe 创建一组同步内存管道。它用于连接一组 `io.Reader` 和 `io. Writer`。管道上的读取和写入是一（多）对一匹配的。
! PipeReader 是 Pipe() 的读取端。
	Close			关闭读取端；管道的后续写入将返回错误 `ErrClosedPipe`。
	CloseWithError	关闭读取端；管道的后续写入将返回错误 `err`。
	Read			从管道中读取数据，阻塞直到写入端末尾或写入端关闭。如果写入端因错误而关闭，则返回该错误。
! PipeWriter 是 Pipe() 的写入端。
	Close 			关闭写入端；管道的后续读取 Read 将返回 `(0, EOF)`。
	CloseWithError	关闭写入端；管道的后续读取将返回错误 `(0, err)` 或 `(0, EOF) // err == nil`。
	Write 			将数据写入管道，阻塞直到一个或多个读取端消耗了所有数据或读取端关闭（返回 `ErrClosedPipe`）。如果读取端因错误而关闭，则返回该错误。
*/
//? go test -v -run=^TestPipe$
func TestPipe(t *testing.T) {
	t.Helper()

	t.Run("Pipe", func(t *testing.T) {
		t.Run("CloseWriter", func(t *testing.T) {
			beforeTest(t)
			pr, pw := io.Pipe()
			// PipeWriter
			go func() {
				defer pw.Close() // 后续返回 EOF
				for range 3 {
					pw.Write([]byte(hello))
					time.Sleep(500 * time.Millisecond)
				}
			}()
			// PipeReader
			defer pr.Close()
			for {
				if n, err := pr.Read(buf); err == nil {
					if n != 0 {
						logf(READ_BYTES, n, buf)
					}
				} else {
					CheckErr(err) // EOF
					break
				}
			}
		})

		t.Run("CloseReader", func(t *testing.T) {
			beforeTest(t)
			pr, pw := io.Pipe()
			// PipeReader
			go func() {
				defer pr.Close()
				for range 3 {
					if n, _ := pr.Read(buf); n > 0 {
						logf(READ_BYTES, n, buf)
					}
					time.Sleep(500 * time.Millisecond)
				}
			}()
			// PipeWriter
			defer pw.Close()
			for {
				if _, err := pw.Write([]byte(hello)); err == nil {
					time.Sleep(500 * time.Millisecond)
				} else {
					CheckErr(err)
					break
				}
			}
		})

	})

	t.Run("Pipe CloseWithError", func(t *testing.T) {
		beforeTest(t)
		pr, pw := io.Pipe()
		var uerr = errors.New("user error")
		go func() {
			pw.CloseWithError(uerr)
		}()
		if _, err := pr.Read(buf); err != nil {
			CheckErr(err)
		}
	})
}

/*
! ReaderAt 包装 `ReadAt` 方法。`ReadAt` 从底层输入源中的偏移 `off` 开始将最多 `len(p)` 字节读入 `p`。
! SectionReader 在底层 `ReaderAt` 的片段上实现 `Read`、`Seek` 和 `ReadAt`。
	io.NewSectionReader 	包装一个 `ReaderAt` 并返回一个 `SectionReader`，它从 `off` 偏移开始读取，并在最多 `n` 个字节处停止。
	Outer 					返回底层 `(ReaderAt, off, n)`。是创建它的 `NewSectionReader` 的逆运算。
	Size 					返回片段的字节大小。
*/
//? go test -v -run=^TestReadAt$
func TestReadAt(t *testing.T) {
	t.Helper()
	t.Run("ReadAt", func(t *testing.T) {
		beforeTest(t)
		var ra io.ReaderAt = strings.NewReader(content)
		readAt := func(_case string, lenbuf int64, offset int64) {
			buf := make([]byte, lenbuf)
			n, err := ra.ReadAt(buf, offset)
			logf("case : %s", _case)
			logf(READ_BYTES, n, buf)
			CheckErr(err)
		}
		readAt("lenbuf(50) > len(content) - offset(10)", 50, 10)   // ok, EOF
		readAt("len(content) - offset(10) > len(content)", 15, 10) // ok
		readAt("offset(40) > len(content)", 10, 50)                // read 0, EOF
		readAt("negative offset(-1)", 10, -1)                      // read 0, Err : negative offset
	})

	t.Run("SectionReader", func(t *testing.T) {
		beforeTest(t)
		readSec := func(_case string, offset, n int64) {
			sr := io.NewSectionReader(strings.NewReader(content), offset, n)
			logf("case : %s", _case)
			readToStdout(sr)
		}

		readSec("offset + n < len(content)", 10, 10)                        // ok
		readSec("offset + n > len(content); offset < len(content)", 10, 50) // ok, full read
		readSec("offset > len(content)", 50, 10)                            // ""
		readSec("negative offset", -10, 20)                                 // "", Err : negative offset
		readSec("negative n", 10, -1)                                       // ok, full read
	})
}

/*
! WriterAt 包装 `WriteAt` 方法。
! OffsetWriter 将基础偏移量处的写入映射到基础写入器中的偏移量 base+off。
	io.NewOffsetWriter 		返回一个 `OffsetWriter`，它从 `off` 偏移开始 `WriterAt` 写入。
*/
// ? go test -v -run=^TestWriteAt$
func TestWriteAt(t *testing.T) {
	t.Helper()

	t.Run("WriteAt", func(t *testing.T) {
		// var wa io.WriterAt =
		// // readAt := func(_case string, lenbuf int64, offset int64) {
		// 	buf := make([]byte, lenbuf)
		// 	n, err := ra.ReadAt(buf, offset)
		// 	logf("case : %s", _case)
		// 	logf(READ_BYTES, n, buf)
		// 	CheckErr(err)
		// }
		// readAt("lenbuf(50) > len(content) - offset(10)", 50, 10)   // ok, EOF
		// readAt("len(content) - offset(10) > len(content)", 15, 10) // ok
		// readAt("offset(40) > len(content)", 10, 50)                // read 0, EOF
		// readAt("negative offset(-1)", 10, -1)                      // read 0, Err : negative offset
	})
}

/* ReaderFrom & WriterTo
! ReaderFrom 包装 `ReadFrom` 方法。`ReadFrom` 从 `r` 读取数据并返回值读取的字节数 `n`。
! WriterTo 包装 `WriteTo` 方法。`WriteTo` 将数据写入 `Writer`。
*/

/* Bytes Readers & Writers
! ByteReader 包装 `ReadByte` 方法。`ReadByte` 读取并返回输入中的下一个字节或错误。
! ByteScanner 将 `UnreadByte` 方法添加到 `ByteReader`。`UnreadByte` 导致下一次调用 `ReadByte` 返回最后读取的字节。
! ByteWriter 包装 `WriteByte` 方法。`WriteAt` 将 `len(p)` 个字节从 `p` 写入偏移量为 `off` 的底层数据流。
*/

/* Other Readers & Writers
! RuneReader 包装 `ReadRune` 方法。`ReadRune` 读取单个编码的 Unicode 字符，并返回该字符及其字节大小。
! RuneScanner 将 `UnreadRune` 方法添加到 `RuneReader`。`UnreadRune` 导致下一次调用 `ReadRune` 返回最后读取的字符。
*/

/*
! StringWriter 包装 `WriteString` 方法。
! WriteString 将字符串 `s` 的内容写入 `w`。如果 `w` 实现了 `StringWriter`，则直接调用 `StringWriter.WriteString`。
*/
//? go test -v -run=^TestStringWriter$
func TestStringWriter(t *testing.T) {
	beforeTest(t)
	var sw io.StringWriter = os.Stdout
	sw.WriteString(hello + "\n")                           // Hello World
	io.WriteString(os.Stdout, strings.ToLower(hello)+"\n") // hello world
}

// ! ReadAll 从 `Reader` 开始读取并返回读取的数据，直到出现错误或 EOF。
// ! ReadAtLeast 从 `Reader` 读取到 `buf`，直到它至少读取了 `min` 字节。
// ! ReadFull 将最多 `len(buf)` 字节从 `Reader` 精确读取到 `buf`。
// ? go test -v -run=^TestReadFunctions$
func TestReadFunctions(t *testing.T) {
	t.Helper()
	content := "some io.Reader stream to be read"
	logBuf := func(buf []byte, err error, n int) {
		if n < 0 {
			logf(READ, buf)
		} else {
			logf(READ_BYTES, n, buf)
		}
		CheckErr(err)
	}

	t.Run("ReadAll", func(t *testing.T) {
		beforeTest(t)
		buff, err := io.ReadAll(newReader(content))
		logBuf(buff, err, -1)
		// output :
		// `some io.Reader stream to be read`
	})

	t.Run("ReadAtLeast", func(t *testing.T) {
		beforeTest(t)
		readAtLeast := func(_case string, min, lenbuf int) {
			buff := make([]byte, lenbuf)
			logf("case : %s", _case)
			n, err := io.ReadAtLeast(newReader(content), buff, min)
			logBuf(buff, err, n)
		}
		readAtLeast("min(10) < len(content) < lenbuf(60))", 10, 60)
		readAtLeast("min(10) < lenbuf(15) < len(content)", 10, 15)
		readAtLeast("len(content) < min(50) < lenbuf(60)", 50, 60)
		readAtLeast("len(content) < lenbuf(50) < min(60)", 60, 50) // io.ErrUnexpectedEOF: unexpected EOF
		readAtLeast("lenbuf(10) < len(content) < min", 50, 10)     // io.ErrShortBuffer: short buffer
		readAtLeast("lenbuf(10) < min(15) < len(content)", 15, 10) // io.ErrShortBuffer: short buffer
	})

	t.Run("ReadFull", func(t *testing.T) {
		readFull := func(_case string, lenbuf int) {
			buff := make([]byte, lenbuf)
			logf("case : %s", _case)
			n, err := io.ReadFull(newReader(content), buff)
			logBuf(buff, err, n)
		}
		readFull("lenbuf < len(content)", 10)
		readFull("lenbuf > len(content)", 50) // ERROR : unexpected EOF
		readFull("lenbuf = len(content)", len(content))
	})
}

// var Discard Writer

// func MultiWriter(writers ...Writer) Writer
// func MultiReader(readers ...Reader) Reader
// func TeeReader(r Reader, w Writer) Reader
// func NopCloser(r Reader) ReadCloser

// func Copy(dst Writer, src Reader) (written int64, err error)
// func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
// func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
// func ReadAll(r Reader) ([]byte, error)
// func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
// func ReadFull(r Reader, buf []byte) (n int, err error)
// func WriteString(w Writer, s string) (n int, err error)

// func Pipe() (*PipeReader, *PipeWriter)

// type PipeReader struct{}

// func (r *PipeReader) Close() error
// func (r *PipeReader) CloseWithError(err error) error
// func (r *PipeReader) Read(data []byte) (n int, err error)

// type PipeWriter struct{}

// func (w *PipeWriter) Close() error
// func (w *PipeWriter) CloseWithError(err error) error
// func (w *PipeWriter) Write(data []byte) (n int, err error)

// type LimitedReader struct {
// 	R Reader
// 	N int64
// }

// func LimitReader(r Reader, n int64) Reader

// func (l *LimitedReader) Read(p []byte) (n int, err error)

// type OffsetWriter struct{}

// func NewOffsetWriter(w WriterAt, off int64) *OffsetWriter
// func (o *OffsetWriter) Seek(offset int64, whence int) (int64, error)
// func (o *OffsetWriter) Write(p []byte) (n int, err error)
// func (o *OffsetWriter) WriteAt(p []byte, off int64) (n int, err error)

// type SectionReader struct{}

// func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader
// func (s *SectionReader) Outer() (r ReaderAt, off int64, n int64)
// func (s *SectionReader) Read(p []byte) (n int, err error)
// func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)
// func (s *SectionReader) Seek(offset int64, whence int) (int64, error)
// func (s *SectionReader) Size() int64
