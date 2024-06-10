package gostd

import (
	"compress/zlib"
	"io"
	"os"
	"testing"
)

/* reader
decompress
! NewReader 从 r 构造一个 ReaderCloser 解压缩器，可能会从 r 中读取更多的数据；返回的 io.ReadCloser 也实现了 Resetter
! NewReaderDict 类似于 NewReader，但使用预设的 dict 作为支持
! flate.Reader 是 NewReader 实际需要的 reader 接口，但如果提供的 io.Reader 没有实现 ReadByte，则使用自己的 buffer
! flate.Resetter 使用 r 和 dict 重置 ReadCloser 解压缩器以重用
compress
! NewWriter 返回一个在给定级别压缩数据的新 Writer;
! NewWriterLevel 类似于 NewWriter，但使用指定的压缩级别 [-2,9]
! NewWriterLevelDict 类似于 NewWriter 但使用预设的 dict, level, 写入 w 的压缩数据只能由使用相同 dict 初始化的 reader 解压缩
! flate.Writer 写入数据的压缩形式到底层 writer
	Close 刷新并关闭写入器；完成后调用 Writer 上的 Close 是调用方的责任
	Flush 将所有挂起的数据刷新到底层写入器。在 zlib 库的术语中，Flush 相当于 Z_SYNC_FLUSH。
	Reset 丢弃 writer 的状态，并使其等效于使用 dst writer 和 w 的压缩级别与 dict 重置 Writer 压缩器以重用
*/

func TestZlibCompress(t *testing.T) {
	for _, f := range fileNames {
		zlibCompress(t, f)
	}
}

func zlibCompress(t *testing.T, file string) {
	os.Remove("testdata/" + file + ".zl")
	in, err := os.Open("testdata/" + file)
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()
	out, err := os.Create("testdata/" + file + ".zl")
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()
	zw, _ := zlib.NewWriterLevel(out, zlib.DefaultCompression)

	if _, err = io.Copy(zw, in); err != nil {
		zw.Close()
		t.Fatal(err)
	}
	zw.Close()
}

func TestZlibReader(t *testing.T) {
	f, err := os.Open("testdata/e.txt.zl")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	zr, err := zlib.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}
	if bs, err := io.ReadAll(zr); err != nil {
		t.Fatal(err)
	} else {
		zr.Close()
		os.Stdout.Write(bs)
	}
}
