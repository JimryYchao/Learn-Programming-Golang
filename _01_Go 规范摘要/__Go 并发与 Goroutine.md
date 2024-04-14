## Go 并发与 Goroutine


在许多环境中，并发编程由于实现对共享变量的正确访问所需的微妙之处而变得困难。Go 语言鼓励一种不同的方法，在这种方法中，共享值在通道上传递，事实上，从来没有被单独的执行线程主动共享过。在任何给定的时间，只有一个 *goroutine* 可以访问该值。根据设计，不会发生数据争用。

---

### Goroutine

*goroutine* 有一个简单的模型：它是一个在同一地址空间中与其他 *goroutine* 并发执行的函数。它是轻量级的，成本比堆栈空间的分配多一点。堆栈开始时很小，所以它们很便宜，并通过根据需要分配（和释放）堆存储来增长。

*goroutine* 被多路复用到多个 OS 线程上，所以如果一个线程阻塞，比如在等待 I/O 时，其他线程会继续运行。它们的设计隐藏了线程创建和管理的许多复杂性。

在函数或方法调用前加上 `go` 关键字，以在新的 *goroutine* 中运行该调用。当调用完成时，*goroutine* 会默默地退出。效果类似于 Unix shell 在后台运行命令的 & 符号。

```go
go list.Sort()  // run list.Sort concurrently; don't wait for it.
```

函数文本（匿名函数）在 *goroutine* 调用中很便捷。匿名函数是闭包的：实现确保函数引用的变量只要处于活动状态就能存活。

```go
func Announce(message string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println(message)
    }()  // Note the parentheses - must call the function.
}
```

---
### Channels

可以利用通道类型在多个 *goroutine* 之间进行通信。

与映射一样，通道也是用 `make` 分配的，结果值充当对底层数据结构的引用。如果提供了一个可选的整数参数，它将设置通道的缓冲区大小。对于无缓冲或同步通道，默认值为零。

```go
ci := make(chan int)            // unbuffered channel of integers
cj := make(chan int, 0)         // unbuffered channel of integers
cs := make(chan *os.File, 100)  // buffered channel of pointers to Files
```

无缓冲通道将通信（值的交换）与同步结合起来，保证两个计算（*goroutine*）处于已知状态。例如，通道可以允许启动的 *goroutine* 等待排序完成。

```go
c := make(chan int)  // Allocate a channel.

// Start the sort in a goroutine; when it completes, signal on the channel.
go func() {
    list.Sort()
    c <- 1  // Send a signal; value does not matter.
}()
doSomethingForAWhile()
<-c   // Wait for sort to finish; discard sent value.
```

在有数据要接收之前，接收方始终处于阻塞状态。如果通道无缓冲，则发送方会阻塞，直到接收方收到该值。如果通道有缓冲区，则发送方仅在值被复制到缓冲区之前阻塞；如果缓冲区已满，则意味着等待某个接收方检索到值。

缓冲通道可以像信号量一样使用，例如限制吞吐量。在下面例子中，传入的请求被传递到 `handle`，它向通道发送一个值，处理请求，然后从通道接收一个值，为下一个消费者准备 “信号量”。通道缓冲区的容量将同时调用的数量限制为 `process`。

```go
var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
    sem <- 1    // Wait for active queue to drain.
    process(r)  // May take a long time.
    <-sem       // Done; enable next request to run.
}

func Serve(queue chan *Request) {
    for {
        req := <-queue
        go handle(req)  // Don't wait for handle to finish.
    }
}
```
