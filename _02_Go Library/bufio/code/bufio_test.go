package gostd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

var content = `This is a content for bufio. 
Hello World!
你好，世界
abcdefg123456789
`
var testInput = []byte("012\n345\n678\n9ab\ncde\nfgh\nijk\nlmn\nopq\nrst\nuvw\nxy")
var testInputrn = []byte("012\r\n345\r\n678\r\n9ab\r\ncde\r\nfgh\r\nijk\r\nlmn\r\nopq\r\nrst\r\nuvw\r\nxy\r\n\n\r\n")

/* buffered input
! bufio.Reader 为 io.Reader 对象实现缓冲。
	Buffered 返回可从当前缓冲区读取的字节数
	Discard 跳过接下来的 n 个字节，返回丢弃的字节数
	Peek 返回接下来的 n 个字节但不推进 reader
	Read 将数据读入 p 并返回写入的字节数
	ReadByte 读取单个字节
	ReadBytes 连续读取字节直到首次遇到 `delim`
	ReadLine 尝试返回单行。行过长时则设置 isPrefix。line 的其余部分在后续调用中返回
	ReadRune 读取一个 UTF-8 编码的 Unicode 字符
	ReadSlice 连续读取字节直到首次遇到 `delim`，返回一个指向缓冲区中字节的切片
	ReadString 连续读取字节直到首次遇到 `delim`，返回一个字符串
	Reset 重置并丢弃任何缓冲数据。重置所有状态并将缓冲区切换从 r 读取
	Size 基础缓冲区的大小
	UnreadByte 取消最近读取操作读取的最后一个字节。Peek、Discard 和 WriteTo 不被视为读操作
	UnreadRune 取消上一次 `ReadRune` 操作读取的字符
	WriteTo 实现了 io.WriterTo。这可能会多次调用底层 Reader 的 Read 或 WriteTo（若实现）
! NewReaderSize & NewReader 构造缓冲 io Reader 对象。
*/
//? go test -v -run=^TestBufReader&
func TestBufReader(t *testing.T) {
	t.Run("read byte, rune", func(t *testing.T) {
		beforeTest(t)
		r := bufio.NewReaderSize(strings.NewReader(content), 64)

		if bytes, err := r.ReadBytes('\n'); err == nil {
			logfln("read %#v, \nto string: %[1]s", bytes[:len(bytes)-1])
		}

		for b, err := r.ReadByte(); err == nil && b != '\n'; b, err = r.ReadByte() {
			logfln("read byte: %#v, %[1]c", b)
		}

		for c, size, err := r.ReadRune(); err == nil && c != '\n'; c, size, err = r.ReadRune() {
			logfln("read rune: %c, size : %d", c, size)
		}

		if _, err := r.Discard(7); err != nil {
			checkErr(err)
		} else {
			bytes, err := r.Peek(10)
			logfln("peek %d bytes: %#v, \nto string: %[2]s", len(bytes), bytes)
			checkErr(err)
		}
	})

	t.Run("ReadLine", func(t *testing.T) {
		beforeTest(t)
		var err error
		var line []byte
		testReadLine := func(input []byte) {
			r := bufio.NewReader(bytes.NewReader(input))
			for err == nil {
				line, _, err = r.ReadLine()
				if len(line) > 0 {
					logfln("read line: %s", line)
				}
			}
			checkErr(err)
			err = nil
		}
		testReadLine(testInput)
		testReadLine(testInputrn)
	})

	t.Run("read slice, string", func(t *testing.T) {
		beforeTest(t)
		r := bufio.NewReader(strings.NewReader(content))
		logfln("buffer size : %d", r.Size())

		s, _ := r.ReadString('\n')
		log(s)

		slice, _ := r.ReadSlice('\n')
		log(strings.ToUpper(string(slice))[:len(slice)-1])

		for r.Buffered() > 0 {
			slice, _ = r.ReadSlice('\n')
			logfln("read slice: %s", slice[0:len(slice)-1])
		}
	})

	t.Run("WriteTo", func(t *testing.T) {
		beforeTest(t)
		r := bufio.NewReader(strings.NewReader("This is a content for WriteTo test\n"))
		r.WriteTo(os.Stdout)
	})
}

/* buffered output
! bufio.Writer 为 io.Writer 对象实现缓冲。
	Available 返回缓冲区中未使用的字节数
	AvailableBuffer 返回一个空的缓冲区。此缓冲区旨在追加并传递给紧接其后的 Writer.Write 调用
	Buffered 返回以写入缓冲区的字节数
	Flush 将所有缓冲数据写入底层 io.Writer
	ReadFrom 实现了 io.ReaderFrom。则此方法可调用底层 writer 的 ReadFrom（若支持）。如果存在缓冲数据和 ReadFrom，则在调用 ReadFrom 之前填充缓冲区并写入缓冲区
	Reset 丢弃所有未 flush 的缓冲数据，清除所有错误并重置所有状态并将缓冲区切换至写入 w
	Size 基础缓冲区的大小
	Write 将 p 的内容写入缓冲区
	WriteByte 写入单个字节
	WriteRune 写入单个 Unicode 代码点
	WriteString 写入字符串
! NewWriterSize & NewWriter 构造缓冲 io Writer 对象。
*/
// ? go test -v -run=^TestBufferedReader&
func TestBufWriter(t *testing.T) {
	w := bufio.NewWriterSize(os.Stdout, 10)
	r := bufio.NewReader(strings.NewReader(content))
	getBytes := func() []byte {
		if bytes, err := r.ReadBytes('\n'); err == nil {
			return bytes
		}
		return []byte{}
	}
	t.Run("WriteByte", func(t *testing.T) {
		beforeTest(t)
		for _, b := range getBytes() {
			if w.Available() < 1 {
				w.Flush()
			}
			w.WriteByte(b)
		}
		w.Flush()
	})

	t.Run("Write & AvailableBuffer", func(t *testing.T) {
		beforeTest(t)
		var b []byte
		if bytes := getBytes(); len(bytes) > w.Size() {
			b = w.AvailableBuffer()
			b = append(b, bytes...)
		} else {
			b = bytes
		}
		w.Write(b)
	})

	t.Run("WriteRune", func(t *testing.T) {
		beforeTest(t)
		runeReader := bytes.NewBuffer(getBytes())
		rn := 0
		for c, _, err := runeReader.ReadRune(); err == nil && c != '\n'; c, _, err = runeReader.ReadRune() {
			w.WriteRune(c)
			rn++
		}
		w.WriteRune('\n')
		w.Flush()
		logfln("total write %d runes", rn)
	})

	t.Run("ReadFrom", func(t *testing.T) {
		beforeTest(t)
		if n, err := w.ReadFrom(r); err == nil { // not need flush
			logfln("total read %d bytes from r", n)
		}
		logfln("Buffered=%d, Size=%d, Available=%d", w.Buffered(), w.Size(), w.Available())
	})

	w.WriteString("test ending\n")
	w.Flush()
}

// ! ReadWriter 保存一组 Reader 和 Writer 的指针
// ? go test -v -run=^TestReadWriter$
func TestReadWriter(t *testing.T) {
	beforeTest(t)
	wr := bufio.NewReadWriter(bufio.NewReader(strings.NewReader(content)), bufio.NewWriter(os.Stdout))

	for {
		line, err := wr.ReadBytes('\n')
		if len(line) > 0 {
			wr.Write(line)
			wr.Flush()
		}
		if err != nil {
			if err == io.EOF {
				return
			} else {
				t.Fatal(err)
			}
		}

	}
}

/* Scanner
! bufio.Scanner 根据给定的拆分函数 split 读取数据并依次返回拆分后的令牌（数据）
	Buffer 设置扫描时使用的初始缓冲区以及扫描期间可能分配的最大缓冲区大小
	Bytes 返回通过调用 Scanner.Scan 生成的最新令牌；非分配，会被后续的 Scan 调用覆盖
	Err 返回扫描程序遇到的第一个非 EOF 错误
	Scan 使 Scanner 前进到下一个令牌，然后可以通过 Scanner.Bytes 或 Scanner.Text 方法获取该令牌。如果 split 函数返回太多的空标记而没有推进输入则 panic
	Split 设置扫描仪的 split 函数。默认 split 函数为 ScanLines。如果在扫描开始后调用 Split，则会发生 panic
	Text 返回调用 Scanner.Scan 生成的最新令牌，并作为新分配的字符串保存其字节
! NewScanner 从 io.Reader 构造一个 Scanner
! ScanBytes 以 byte 为令牌单元进行拆分
! ScanLines 以 line 为令牌单元进行拆分，并丢弃尾随的行尾标记
! ScanRunes 以 rune 为令牌单元进行拆分
! ScanWords 以空白分隔的文本为令牌单元进行拆分
! bufio.SplitFunc 表示 split 函数的签名
*/
// ? go test -v -run=^TestScanner$
func TestScanner(t *testing.T) {
	beforeTest(t)
	scan := func(fn bufio.SplitFunc, kind string) {
		scr := bufio.NewScanner(strings.NewReader(content))
		scr.Split(fn)
		scr.Buffer(make([]byte, 512), 512)

		for scr.Scan() {
			logfln("scan %s : %s", kind, scr.Text())
		}
		checkErr(scr.Err())
		log("")
	}

	scan(bufio.ScanLines, "line")
	scan(bufio.ScanRunes, "rune")
	scan(bufio.ScanWords, "word")
	scan(bufio.ScanBytes, "byte")
}

// ? go test -v -run=^TestCustomSpiltFunc$
func TestCustomSpiltFunc(t *testing.T) {
	beforeTest(t)
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	// Set the split before start scanning
	scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}

// ?  go test -v -run=^TestScannerFile$
func TestScannerFile(t *testing.T) {
	beforeTest(t)
	file, err := os.Open("readline.file")
	if err != nil {
		t.Fatal(err)
	}
	scr := bufio.NewScanner(file)

	scr.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		var c = 0
		for len(data) > 0 { // 跳过空行
			if data[0] == '\r' || data[0] == '\n' {
				c++
				data = data[1:]
			} else {
				advance, token, err = bufio.ScanLines(data, atEOF)
				return advance + c, token, err
			}
		}
		return 0, nil, io.EOF
	})

	for scr.Scan() {
		logfln("scan line: %s", scr.Text())
	}
	checkErr(scr.Err())
}
