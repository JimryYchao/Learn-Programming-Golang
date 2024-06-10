package gostd

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	l "log"
	"os"
	"strings"
	"sync"
	"testing"
)

/*
decompress
! NewReader 从 r 构造一个 ReaderCloser 解压缩器，可能会从 r 中读取更多的数据；返回的 io.ReadCloser 也实现了 Resetter
! NewReaderDict 类似于 NewReader，但使用预设的 dict 作为支持
! flate.Reader 是 NewReader 实际需要的 reader 接口，但如果提供的 io.Reader 没有实现 ReadByte，则使用自己的 buffer
! flate.Resetter 使用 r 和 dict 重置 ReadCloser 解压缩器以重用
compress
! NewWriter 返回一个在给定级别压缩数据的新 Writer;
	在 zlib 之后，级别范围从 1（最佳速度）到 9（最佳压缩）
	0 表示无压缩, 但添加必要的 DEFLATE 帧; -1 表示默认压缩级别
	-2 表示仅使用霍夫曼压缩, 对所有类型的输入提供非常快速的压缩，但牺牲了相当大的压缩效率。
! NewWriterDict 类似于 NewWriter 但使用预设的 dict, 写入 w 的压缩数据只能由使用相同 dict 初始化的 reader 解压缩
! flate.Writer 写入数据的压缩形式到底层 writer
	Close 刷新并关闭写入器
	Flush 将所有挂起的数据刷新到底层写入器。在 zlib 库的术语中，Flush 相当于 Z_SYNC_FLUSH。
	Reset 丢弃 writer 的状态，并使其等效于使用 dst writer 和 w 的压缩级别与 dict 重置 Writer 压缩器以重用
*/

func TestWriter(t *testing.T) {

	// 字典是一个字节串。当压缩某些输入数据时，压缩器将尝试用字典中找到的匹配项替换子字符串。
	// 因此，字典应该只包含预期会在实际数据流中找到的子字符串。
	const dict = `<?xml version="1.0"?>` + `<book>` + `<data>` + `<meta name="` + `" content="`

	// 要压缩的数据应该(但不是必需)包含与字典中的子字符串匹配的频繁子字符串。
	const data = `<?xml version="1.0"?>
<book>
	<meta name="title" content="The Go Programming Language"/>
	<meta name="authors" content="Alan Donovan and Brian Kernighan"/>
	<meta name="published" content="2015-10-26"/>
	<meta name="isbn" content="978-0134190440"/>
	<data>...</data>
</book>
`
	var b bytes.Buffer
	// 使用指定的字典压缩数据。
	zw, _ := flate.NewWriterDict(&b, flate.DefaultCompression, []byte(dict))
	io.Copy(zw, strings.NewReader(data))
	zw.Close()

	print(len([]byte(data)), "   ", b.Len(), "\n")

	// 解压器必须与压缩器使用相同的字典。否则，输入可能显示为损坏。
	fmt.Println("Decompressed output using the dictionary:")
	zr := flate.NewReaderDict(bytes.NewReader(b.Bytes()), []byte(dict))
	io.Copy(os.Stdout, zr)
	zr.Close()
	fmt.Println()

	// 用 '#' 替换字典中的所有字节，以直观地演示使用预设字典的近似有效性。
	fmt.Println("Substrings matched by the dictionary are marked with #:")
	hashDict := []byte(dict)
	for i := range hashDict {
		hashDict[i] = '#'
	}
	zr = flate.NewReaderDict(&b, hashDict)
	io.Copy(os.Stdout, zr)
	zr.Close()
}

// DEFLATE 适用于通过网络传输压缩数据。
func TestNetConnect(t *testing.T) {
	var wg sync.WaitGroup
	defer wg.Wait()

	// 使用 io.Pipe 模拟一个 network 连接。真实的网络应用程序应该注意正确关闭底层连接。
	rp, wp := io.Pipe()

	// 启动一个程序作为发送端。
	wg.Add(1)
	go func() {
		defer wg.Done()
		zw, err := flate.NewWriter(wp, flate.BestSpeed)
		if err != nil {
			l.Fatal(err)
		}

		b := make([]byte, 256)
		for _, m := range strings.Fields("A long time ago in a galaxy far, far away...") {
			// 使用一个简单的帧格式，其中第一个字节是消息长度，后面跟着消息本身。
			b[0] = uint8(copy(b[1:], m))
			if _, err := zw.Write(b[:1+len(m)]); err != nil {
				l.Fatal(err)
			}
			// Flush 确保接收端可以读取到目前为止发送的所有数据。
			if err := zw.Flush(); err != nil {
				l.Fatal(err)
			}
		}
		if err := zw.Close(); err != nil {
			l.Fatal(err)
		}
	}()

	// 启动一个程序来充当接收端。
	wg.Add(1)
	go func() {
		defer wg.Done()
		zr := flate.NewReader(rp)
		b := make([]byte, 256)
		for {
			// 读取消息长度。这保证在发送端每次相应的 Flush 和 Close 都返回。
			if _, err := io.ReadFull(zr, b[:1]); err != nil {
				if err == io.EOF {
					break // 发送端关闭了信号流
				}
				l.Fatal(err)
			}

			// 读取消息内容
			n := int(b[0])
			if _, err := io.ReadFull(zr, b[:n]); err != nil {
				l.Fatal(err)
			}

			fmt.Printf("Received %d bytes: %s\n", n, b[:n])
		}
		if err := zr.Close(); err != nil {
			l.Fatal(err)
		}
	}()
}

func TestFlate(t *testing.T) {
	for _, f := range fileNames {
		flateFile(f)
	}
}

func flateFile(file string) error {
	f, err := os.Open("testdata/" + file)
	if err != nil {
		return err
	}
	defer f.Close()

	out, _ := os.Create("testdata/" + file + ".flate")
	defer func() {
		out.Close()
		if err != nil {
			os.Remove(out.Name())
		}
	}()

	zw, err := flate.NewWriter(out, flate.DefaultCompression)
	if err != nil {
		return err
	}
	defer zw.Close()

	if _, err = io.Copy(zw, f); err != nil {
		return err
	}
	return nil
}

func TestDeflate(t *testing.T) {
	deflateFile("e.txt.flate")
}

func deflateFile(file string) error {
	f, err := os.Open("testdata/" + file)
	if err != nil {
		return nil
	}
	defer f.Close()

	var b bytes.Buffer = *bytes.NewBuffer(make([]byte, 4096))
	zr := flate.NewReader(f)
	if _, err := io.Copy(&b, zr); err != nil {
		return err
	} else {
		fmt.Printf("%s", b.Bytes())
	}
	return nil
}
