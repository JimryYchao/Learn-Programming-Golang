package gostd

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync/atomic"
	"testing"
	"time"
)

// ! log Functions

func TestLogFatal(t *testing.T) {

}

func TestLogPanic(t *testing.T) {

}

/* standard Logger
! New 创建一个新的 `log.Logger` 并关联一个 `io.Writer out`。`prefix` 出现在每个生成日志行的开头。`flag` 定义日志记录的属性。
	Ldata : 2009/01/23
	Ltime : 01:23:23
	Lmicroseconds : 01:23:23.123123
	Llongfile : /a/b/c/d.go:23
	Lshortfile : d.go:23
	LUTC : 使用 UTC 时区
	Lmsgprefix : 将前缀 prefix 移动到 mess 前
	LstdFlags = Ldate | Ltime
! log.Logger 表示一个活动的日志记录对象，它生成到 io.Writer 的输出行。每个日志记录操作都调用 `Writer` 的 Write() 方法。Logger 可以同时在多个 goroutine 中使用；它保证了对 Writer 的序列化访问。
! Default 返回一个包级输出函数使用的标准日志 Logger。
*/
//? go test -v -run=^TestStdLogger$
func TestStdLogger(t *testing.T) {
	t.Run("std logger", func(t *testing.T) {
		beforeTest(t)
		t.Cleanup(func() {
			// reset std logger
			log.SetFlags(log.LstdFlags)
			log.SetOutput(os.Stdout)
			log.SetPrefix("")
		})

		log.Print("hello", "World")
		log.Print("hello", 1, 2, 3, "World")
		log.Println("hello", 1, 2, 3, "World")

		stdlogger := log.Default()

		fmt.Printf("LOG Prefix is `%s`\n", log.Prefix())
		log.SetPrefix("LOG_TEST: ") // 设置日志输出的前缀
		stdlogger.Print(time.Now())

		tmpf, _ := os.CreateTemp(t.TempDir(), "tmplogs")
		defer tmpf.Close()

		log.SetOutput(tmpf)
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lmsgprefix)

		stdlogger.Print("Hello, 世界")
		tmpf.Seek(0, io.SeekStart) // 移动文件位置指示器至开头
		io.Copy(os.Stdout, tmpf)
	})

	t.Run("parallel logger", func(t *testing.T) {
		beforeTest(t)
		compCh := make(chan bool)
		var ngorte atomic.Int32
		logInGoroutine := func(logger *log.Logger, n int) {
			ngorte.Add(1)
			for i := range n {
				logger.Printf("line %d", i)
				time.Sleep(1 * time.Second)
			}
			compCh <- true
		}
		go logInGoroutine(log.New(os.Stdout, "[DEBUG]", log.Lmsgprefix), 10)
		go logInGoroutine(log.New(os.Stdout, "[WARNING]", log.LstdFlags|log.Lmsgprefix), 15)
		go logInGoroutine(log.New(os.Stdout, "", log.Ldate|log.LUTC), 5)
		go logInGoroutine(log.New(os.Stdout, "[FATAL]", log.LstdFlags|log.Lmsgprefix|log.Lmicroseconds), 3)
		go logInGoroutine(log.New(os.Stdout, "[INFO]", log.Lmsgprefix), 8)

		var n int32 = 0
		for <-compCh {
			n++
			if ngorte.Load() <= n {
				break
			}
		}
	})
}
