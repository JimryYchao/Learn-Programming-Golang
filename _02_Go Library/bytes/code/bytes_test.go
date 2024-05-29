package gostd

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"testing"
)

// bytes functions
// ? go test -v -run=^TestBytesFunctions$
func TestBytesFunctions(t *testing.T) {
	t.Run("generic functions", func(t *testing.T) {
		//! Some functions in `bytes` like `strings` see [strings_test.go]

		//! Rune 将 []byte 解释为 []rune
		rs := bytes.Runes([]byte("go gopher"))
		for _, r := range rs {
			logfln("%#U", r)
		}
	})
}

// ! bytes.Buffer 是一个可变大小的字节缓冲区，具有 Buffer.Read 和 Buffer.Write 方法
// ! bytes.Reader 实现 io.Reader,io.ReaderAt,io.WriterTo,io.Seeker,io.ByteScanner,io.RuneScanner 接口以读取字节片。与 Buffer 不同，Reader 是只读的
func TestBuffer(t *testing.T) {
	t.Run("new bytes.Buffer", func(t *testing.T) {
		var b bytes.Buffer // A Buffer needs no initialization. == bytes.NewBuffer(nil)
		b.Write([]byte("Hello "))
		fmt.Fprintf(&b, "world!\n")
		b.WriteTo(os.Stdout) // Hello World!

		// A Buffer can turn a string or a []byte into an io.Reader.
		buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
		io.Copy(os.Stdout, base64.NewDecoder(base64.StdEncoding, buf)) // Gophers rule!
	})

	//! Buffer.Bytes() 返回一个长度为 B.Len() 的缓冲区的未读部分的切片，尽在下一次调用 Buffer.Read、Buffer.Write、Buffer.Reset 或 Buffer.Truncate 等方法前有效
	t.Run("Buffer.Bytes", func(t *testing.T) {
		buf := bytes.Buffer{}
		buf.Write([]byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd', '\n'})
		os.Stdout.Write(buf.Bytes()) // hello world
		buf.Read(make([]byte, 6))
		os.Stdout.Write(buf.Bytes()) // world
	})

	//! Buffer.Next(n) 返回一个包含缓冲区中接下来的 n 个字节的切片, 并使缓冲区前进（如同 Buffer.Read）
	t.Run("Buffer.Next", func(t *testing.T) {
		var b bytes.Buffer
		b.Write([]byte("abcde"))
		logfln("%s", b.Next(2)) // ab
		logfln("%s", b.Next(2)) // cd
		logfln("%s", b.Next(2)) // e
	})

	//! Buffer.Truncate(n) 从缓冲区中丢弃除前 n 个未读字节以外的所有字节，但继续使用相同的已分配存储
	t.Run("Buffer.Truncate", func(t *testing.T) {
		b := bytes.NewBufferString("Hello World")
		b.Truncate(6)
		os.Stdout.Write(b.Bytes()) // Hello
		b.WriteString("WORLD")
		os.Stdout.Write(b.Bytes()) // Hello WORLD
	})
}
