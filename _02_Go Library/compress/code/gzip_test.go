package gostd

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

/*
! gzip.Header 存储一个头文件，以给出压缩文件的元数据。该头文件公开为 Writer 和 Reader 结构的字段
	字符串必须是 UTF-8 编码的，并且由于 GZIP 文件格式的限制，只能包含 Unicode 码位 U+0001 到 U+00FF

Reader
! NewReader 从 r 构造一个解压缩器 gzip.Reader，调用方有责任在完成时调用 Reader 上的 Close。Header 在读取后有效
! gzip.Reader 读取 gzip 格式的压缩文件的未压缩数据
 	Close 关闭读取器。它不会关闭底层 io.Reader。为了验证 GZIP 校验和，读取器必须完全消耗，直到 io.EOF
	Multistream 	控制读取器是否支持多数据流文件；如果启用（默认值），读取器期望输入是一个单独的 gzip 压缩数据流序列，
					每个数据流都有自己的头部和尾部，结尾为 io.EOF。其效果是，对一系列 gzip 格式压缩文件的连接被视为等同于该序列连接的 gzip。
					这是 gzip 读取器的标准行为。调用 Multistream(false) 将禁用此行为;

					当读取文件格式区分单个 gzip 数据流或将 gzip 数据流与其他数据流混合时，禁用此行为非常有用。在这种模式下，当 Reader 到达数据流的末尾时，
					Reader.Read 返回 io.EOF。底层读取器必须实现 io.ByteReader，以便将其定位在 gzip 流之后。要启动下一个流，请调用 z.Reset(r)，
					然后调用 z.Multisream(false)。如果没有下一个流，z.Reset(r) 将返回 io.EOF。

Writer
! NewWriter 从 w 构造一个压缩器 gzip.Writer, 调用方有责任在完成时调用 Writer 上的 Close。
! NewWriterLevel 类似于 NewWriter, 可以设置压缩级别	[-2,9]
! gzip.Writer 压缩并写入到底层 writer；希望设置 Header 时，调用方必须在第一次调用 Write, Flush, Close 时设置
	Close 刷新并关闭写入器。它不会关闭底层 io.Writer
	Flush 刷新将任何挂起的压缩数据刷新到底层 io.Writer; 它主要用于压缩网络协议，以确保远程读取器有足够的数据来重建数据包。如果底层写入器返回错误，Flush 将返回该错误。
	Reset 用于重置 w 的状态。
*/
//?
func TestGzipCompress(t *testing.T) {
	for _, f := range fileNames {
		gzipCompress(t, f)
	}
}

func gzipCompress(t *testing.T, f string) {
	os.Remove("testdata/" + f + ".gz")
	in, err := os.Open("testdata/" + f)
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()
	out, err := os.Create("testdata/" + f + ".gz")
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()
	zw := gzip.NewWriter(out)
	zw.Name = "e.txt"
	zw.Comment = "author: JimryYchao"
	zw.ModTime = time.Now().UTC()
	zw.Extra = []byte("hello gzip")

	if _, err = io.Copy(zw, in); err != nil {
		zw.Close()
		t.Fatal(err)
	}
	zw.Close()
}

func TestGzipReader(t *testing.T) {
	f, err := os.Open("testdata/e.txt.gz")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	zr, err := gzip.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}
	if bs, err := io.ReadAll(zr); err != nil {
		t.Fatal(err)
	} else {
		zr.Close()
		fmt.Printf("Name: %s\nComment: %s\nModTime: %s\nExtra: %s\nOS: %d\n", zr.Name, zr.Comment, zr.ModTime, zr.Extra, zr.OS)
		os.Stdout.Write(bs)
	}
}

func TestGzipMultistream(t *testing.T) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	var files = []struct {
		name    string
		comment string
		modTime time.Time
		data    string
	}{
		{"file-1.txt", "file-header-1", time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC), "Hello Gophers - 1"},
		{"file-2.txt", "file-header-2", time.Date(2007, time.March, 2, 4, 5, 6, 1, time.UTC), "Hello Gophers - 2"},
	}

	for _, file := range files {
		zw.Name = file.name
		zw.Comment = file.comment
		zw.ModTime = file.modTime

		if _, err := zw.Write([]byte(file.data)); err != nil {
			log.Fatal(err)
		}

		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}

		zw.Reset(&buf)
	}

	zr, err := gzip.NewReader(&buf)
	if err != nil {
		log.Fatal(err)
	}

	for {
		zr.Multistream(false)
		fmt.Printf("Name: %s\nComment: %s\nModTime: %s\n\n", zr.Name, zr.Comment, zr.ModTime.UTC())

		if _, err := io.Copy(os.Stdout, zr); err != nil {
			log.Fatal(err)
		}

		fmt.Print("\n\n")

		err = zr.Reset(&buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestCompressingReader(t *testing.T) {
	// 这是一个编写压缩/读取器的示例。这对于 HTTP 客户端主体非常有用，如下所示。
	const testdata = "the data to be compressed"

	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		zr, err := gzip.NewReader(req.Body)
		if err != nil {
			log.Fatal(err)
		}

		// 输出示例的数据。
		if _, err := io.Copy(os.Stdout, zr); err != nil {
			log.Fatal(err)
		}
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// 要压缩的数据，作为 io.Reader
	dataReader := strings.NewReader(testdata)

	// bodyReader 作为 io.Reader 是 HTTP 请求的主体。
	// httpWriter 作为 io.Writer 是 HTTP 请求的主体。
	bodyReader, httpWriter := io.Pipe()

	//确保 bodyReader 始终处于关闭状态，以便下面的 goroutine 将永远退出。
	defer bodyReader.Close()

	// gzipWriter 将数据压缩到 httpWriter。
	gzipWriter := gzip.NewWriter(httpWriter)

	// errch 从写入 goroutine 中收集任何错误。
	errch := make(chan error, 1)

	go func() {
		defer close(errch)
		sentErr := false
		sendErr := func(err error) {
			if !sentErr {
				errch <- err
				sentErr = true
			}
		}
		// 复制数据到 gzipWriter, gzipWriter 会将数据压缩并发送到 bodyReader。
		if _, err := io.Copy(gzipWriter, dataReader); err != nil && err != io.ErrClosedPipe {
			sendErr(err)
		}
		if err := gzipWriter.Close(); err != nil && err != io.ErrClosedPipe {
			sendErr(err)
		}
		if err := httpWriter.Close(); err != nil && err != io.ErrClosedPipe {
			sendErr(err)
		}
	}()

	// 向测试服务器发送 HTTP 请求。
	req, err := http.NewRequest("PUT", ts.URL, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	// 将 request 传递给 http.Client.Do 将保证关闭 body，在本例中是 bodyReader。
	resp, err := ts.Client().Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// 检查压缩数据是否出错。
	if err := <-errch; err != nil {
		log.Fatal(err)
	}

	// 在这个例子中，不关心响应。
	resp.Body.Close()
}
