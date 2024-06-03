### Go Standard library Summary      <a href="https://pkg.go.dev/std" target="_blank"><img src="./_rsc/link-src.drawio.png"/></a>

- 官网链接 <img src="./_rsc/link-src.drawio.png"/> 
- 补充说明  <img  src="./_rsc/link-others.drawio.png"/>
- 代码  <img src="./_rsc/link-code.drawio.png"/>
- 示例  <img src="./_rsc/link-exam.drawio.png"/>

---

- [x] bufio 包装了 `io.Reader` 和 `io.Writer` 并提供了缓冲和一些文本 I/O 帮助。<a href="https://pkg.go.dev/bufio" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./bufio/code/bufio_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

- [x] bytes 实现了操作字节切片的函数。它类似于 strings 包的功能。       <a href="https://pkg.go.dev/bytes" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./bytes/code/bytes_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

- [x] context 定义了 Context 上下文类型，它携带 *deadlines*、*cancellation signals* 和跨 API 边界和进程之间的其他 *request-scoped  values*。       <a href="https://pkg.go.dev/context" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="other"/></a><a href="./context/context.md"  ><img  src="./_rsc/link-others.drawio.png" 
  id="code"/></a><a href="./context/code/context_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  id="exam"/></a><a href="./context/context.md#exam"   ><img src="./_rsc/link-exam.drawio.png"
  /></a>

- [x] embed 提供了对嵌入在运行的 Go 程序中的文件的访问。      <a href="https://pkg.go.dev/embed" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="other"/></a><a href="./embed/embed.md"  ><img  src="./_rsc/link-others.drawio.png" 
  id="code"/></a><a href="./embed/code/embed_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

- [x] errors 实现一些函数来处理错误。       <a href="https://pkg.go.dev/errors" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./errors/code/errors_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>


- [x] fmt 使用类 C 的 `printf` 和 `scanf` 的函数实现格式化 I/O。“*verbs*” 格式从 C 派生的。       <a href="https://pkg.go.dev/" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="other"/></a><a href="./fmt/fmt.md"  ><img  src="./_rsc/link-others.drawio.png" 
  id="code"/></a><a href="./fmt/code/fmt_test.go"><img src="./_rsc/link-code.drawio.png" 
  /></a>


- [x] io 提供 I/O 原语的基本接口。       <a href="https://pkg.go.dev/io" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./io/code/io_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

  - [x] fs 定义了文件系统的基本接口。<a href="https://pkg.go.dev/io/fs" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./io/code/fs_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

- [x] log 实现了一个简单的日志记录包。       <a href="https://pkg.go.dev/log" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./log/code/log_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  id="exam"/></a><a href="./log/log.md#exam.md"   ><img src="./_rsc/link-exam.drawio.png"
  /></a>

  - [x] syslog 为系统日志服务提供了一个简单的接口。它可以使用 UNIX 域套接字、UDP 或 TCP 向系统日志守护程序发送消息。       <a href="https://pkg.go.dev/log/syslog"  target="_blank"><img src="./_rsc/link-src.drawio.png" /></a>

  - [x] slog 提供结构化日志记录，其中日志记录包括消息、严重性级别和以键值对表示的各种其他属性。       <a href="https://pkg.go.dev/log/slog" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="other"/></a><a href="./log/slog.md"  ><img  src="./_rsc/link-others.drawio.png" 
  id="code"/></a><a href="./log/code/slog_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  id="exam"/></a><a href="./log/slog.md#exam"   ><img src="./_rsc/link-exam.drawio.png"
  /></a>

- [x] maps 定义了各种对任何类型的映射的辅助函数。      <a href="https://pkg.go.dev/maps"  target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./maps/code/maps_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

- [x] math 提供基本常量和数学函数。        <a href="https://pkg.go.dev/math" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./math/code/math_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

  - [x] big 实现了任意精度的算术（大数字）。        <a href="https://pkg.go.dev/math/big" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="other"/></a><a href="./math/big.md"  ><img  src="./_rsc/link-others.drawio.png" 
  id="code"/></a><a href="./math/code/big_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  id="exam"/></a><a href="./math/big.md#exam"   ><img src="./_rsc/link-exam.drawio.png"
  /></a>

  - [x] bits 为预先声明的无符号整数类型实现位计数和操作函数。        <a href="https://pkg.go.dev/math/bits" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./math/code/bits_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

  - [x] cmplx 为复数提供基本常数和数学函数。        <a href="https://pkg.go.dev/math/cmplx" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./math/code/cmplx_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

  - [x] rand 实现了适合模拟等任务的伪随机数生成器。       <a href="https://pkg.go.dev/math/rand" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="other"/></a><a href="./math/rand.md"  ><img  src="./_rsc/link-others.drawio.png" 
  id="code"/></a><a href="./math/code/rand_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>
      - [x] rand/v2 实现了适合模拟等任务的伪随机数生成器。       <a href="https://pkg.go.dev/math/rand/v2" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./math/code/randv2_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

- [x] slices 定义了对任何类型的切片的辅助函数。      <a href="https://pkg.go.dev/slices"  target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./slices/code/slices_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>


- [x] strings 实现了一些函数来操作 UTF-8 编码的字符串。      <a href="https://pkg.go.dev/strings "  target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./strings/strings.md"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

- [x] sync 提供基本的同步原语。        <a href="https://pkg.go.dev/sync" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./sync/code/sync_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  id="exam"/></a><a href="./sync/sync.md#exam"   ><img src="./_rsc/link-exam.drawio.png"
  /></a>
  - [x] atomic 提供了用于实现同步算法的低级原子内存原语。        <a href="https://pkg.go.dev/sync/atomic" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="other"/></a><a href="./sync/atomic.md"  ><img  src="./_rsc/link-others.drawio.png" 
  id="code"/></a><a href="./sync/code/atomic_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

- [x] testing 为 Go 包提供自动化测试支持。<a href="https://pkg.go.dev/testing" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="other"/></a><a href="./testing/testing.md"  ><img  src="./_rsc/link-others.drawio.png" 
  id="code"/></a><a href="./testing/code/testing_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>
  
  - [x] fstest 实现了对测试实现和文件系统用户的支持。      <a href="https://pkg.go.dev/testing/fstest" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./testing/code/fstest_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>
  
  - [x] iotest 实现了主要用于测试的 Readers 和 Writers。        <a href="https://pkg.go.dev/testing/iotest" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./testing/code/iotest_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

  - [x] quick 实现了一些实用函数来帮助进行黑盒测试。        <a href="https://pkg.go.dev/testing/quick" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./testing/code/quick_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>

  - [x] slogtest 实现了对 `log/slog.Handler` 的测试实现的支持。       <a href="https://pkg.go.dev/testing/slogtest" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="code"/></a><a href="./testing/code/slogtest_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  /></a>


- [x] time 提供测量和显示时间的功能。      <a href="https://pkg.go.dev/time" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="other"/></a><a href="./time/time.md"  ><img  src="./_rsc/link-others.drawio.png" 
  id="code"/></a><a href="./time/code/time_test.go"   ><img src="./_rsc/link-code.drawio.png" 
  id="exam"/></a><a href="./time/time.md#exam"   ><img src="./_rsc/link-exam.drawio.png"
  /></a>



<!-- 

- [ ]         <a href="https://pkg.go.dev/#" target="_blank"><img src="./_rsc/link-src.drawio.png" 
  id="other"/></a><a href="#"  ><img  src="./_rsc/link-others.drawio.png" 
  id="code"/></a><a href="#"   ><img src="./_rsc/link-code.drawio.png" 
  id="exam"/></a><a href="#exam"   ><img src="./_rsc/link-exam.drawio.png"
  /></a>

-->
