<a id="TOP"></a>

## Package bufio

<div id="top" style="z-index:99999999;position:fixed;bottom:35px;right:50px;float:right">
	<a href="./code/bufio_test.go" target="_blank"><img id="img-code" src="../_rsc/to-code.drawio.png" ></img></a>
	<!-- <a href="#TOP" ><img id="img-top" src="../_rsc/to-top.drawio.png" ></img></a> -->	
	<a href="https://pkg.go.dev/bufio" ><img id="img-link" src="../_rsc/to-link.drawio.png" ></img></a>
	<a href="..\README.md"><img id="img-back" src="../_rsc/back.drawio.png"></img></a>
</div>

包 `bufio` 实现了缓冲 I/O。它包装了一个 `io.Reader` 或 `io.Writer` 对象，创建了另一个对象（`Reader` 或 `Writer`），该对象也实现了该接口，但提供了缓冲和一些文本 I/O 帮助。


---


```go
func main() {
	file, _ := os.Open("_02_Go Library/Go Library summary/bufio/ref/readline.file")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 512), 512)
	scanner.Split(bufio.ScanLines) // 行扫描，默认行为
	for scanner.Scan() {
		fmt.Printf("Read line : %q\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
```

```go
// 读取标准输入
func ScanLines() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines) // 设置行拆分
	fmt.Println(`Start ScanLines in STDIN, input "^Z" to stop scanning.`)
	for scanner.Scan() {
		fmt.Printf("read a line : %q\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
```
```go
// 读取一段输入中的英文字符
func ScanRunes() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)
	fmt.Println(`Start ScanRunes in STDIN, input "^Z" to stop scanning.`)
	for scanner.Scan() {
		rangeTable := []*unicode.RangeTable{unicode.Lu, unicode.Ll}
		if l := scanner.Bytes(); unicode.IsOneOf(rangeTable, bytes.Runes(l)[0]) {
			fmt.Printf("read a letter : %q\n", string(l))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
```

```go
// 读取标准输入中的数字
func ScanWords() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	fmt.Println(`Start ScanWords in STDIN, input "^Z" to stop scanning.`)
	for scanner.Scan() {
		s := scanner.Text()
		if num, ok := func(s string) (string, bool) {
			_, err := strconv.ParseFloat(s, 64)
			return s, err == nil
		}(s); ok {
			fmt.Printf("read a number : %q\n", num)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
```