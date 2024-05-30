### Go Standard library Summary      <a href="https://pkg.go.dev/std"><img src="./_rsc/link-src.drawio.png"/></a>

- 官网链接 <img src="./_rsc/link-src.drawio.png"/> 
- 补充说明  <img  src="./_rsc/link-others.drawio.png"/>
- 代码  <img src="./_rsc/link-code.drawio.png"/>
- 示例  <img src="./_rsc/link-exam.drawio.png"/>

---

- [x] bufio 包装了 `io.Reader` 和 `io.Writer` 并提供了缓冲和一些文本 I/O 帮助。<a href="https://pkg.go.dev/bufio"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./bufio/code/bufio_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

- [x] bytes 实现了操作字节切片的函数。它类似于 strings 包的功能。       <a href="https://pkg.go.dev/bytes"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./bytes/code/bytes_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

- [x] context 定义了 Context 上下文类型，它携带 *deadlines*、*cancellation signals* 和跨 API 边界和进程之间的其他 *request-scoped  values*。       <a href="https://pkg.go.dev/context"><img src="./_rsc/link-src.drawio.png" 
  id="other"/></a><a href="./context/context.md"  ><img  src="./_rsc/link-others.drawio.png" 
  id="code"/></a><a href="./context/code/context_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  id="exam"/></a><a href="./context/context.md#exam"   ><img src="./_rsc/link-exam.drawio.png"
  /></a>

- [x] errors 实现一些函数来处理错误。       <a href="https://pkg.go.dev/errors"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./errors/code/errors_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>


- [x] fmt 使用类 C 的 `printf` 和 `scanf` 的函数实现格式化 I/O。“*verbs*” 格式从 C 派生的。       <a href="https://pkg.go.dev/"><img src="./_rsc/link-src.drawio.png" 
  id="other"/></a><a href="./fmt/fmt.md"  ><img  src="./_rsc/link-others.drawio.png" 
  id="code"/></a><a href="./fmt/code/fmt_test.go"><img src="./_rsc/link-code.drawio.png" 
  /></a>


- [x] io 提供 I/O 原语的基本接口。       <a href="https://pkg.go.dev/io"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./io/code/io_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

  - [x] fs 定义了文件系统的基本接口。<a href="https://pkg.go.dev/io/fs"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./io/code/fs_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

- [x] log 实现了一个简单的日志记录包。       <a href="https://pkg.go.dev/log"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./log/code/log_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  id="exam"/></a><a href="./log/log.md#exam.md"   ><img src="./_rsc/link-exam.drawio.png"
  /></a>

  - [x] syslog 为系统日志服务提供了一个简单的接口。它可以使用 UNIX 域套接字、UDP 或 TCP 向系统日志守护程序发送消息。       <a href="https://pkg.go.dev/log/syslog"><img src="./_rsc/link-src.drawio.png" /></a>

  - [x] slog 提供结构化日志记录，其中日志记录包括消息、严重性级别和以键值对表示的各种其他属性。       <a href="https://pkg.go.dev/log/slog"><img src="./_rsc/link-src.drawio.png" 
  id="other"/></a><a href="./log/slog.md"  ><img  src="./_rsc/link-others.drawio.png" 
  id="code"/></a><a href="./log/code/slog_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  id="exam"/></a><a href="./log/slog.md#exam"   ><img src="./_rsc/link-exam.drawio.png"
  /></a>

- [x] strings 实现了一些函数来操作 UTF-8 编码的字符串。      <a href="https://pkg.go.dev/strings"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./strings/strings.md"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

- [x] testing 为 Go 包提供自动化测试支持。<a href="https://pkg.go.dev/testing"><img src="./_rsc/link-src.drawio.png" 
  id="other"/></a><a href="./testing/testing.md"  ><img  src="./_rsc/link-others.drawio.png" 
  id="code"/></a><a href="./testing/code/testing_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>





<!-- 

- [ ]         <a href="https://pkg.go.dev/#"><img src="./_rsc/link-src.drawio.png" 
  id="other"/></a><a href="#"  ><img  src="./_rsc/link-others.drawio.png" 
  id="code"/></a><a href="#"   ><img src="./_rsc/link-code.drawio.png" 
  id="exam"/></a><a href="#exam"   ><img src="./_rsc/link-exam.drawio.png"
  /></a>

-->
