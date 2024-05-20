<div id="top" style="z-index:99999999;position:fixed;bottom:35px;right:50px;float:right">
	<a href="./code/ioutil_test.go" target="_blank"><img id="img-code" src="../_rsc/to-code.drawio.png" ></img></a>
	<a href="#TOP" ><img id="img-top" src="../_rsc/to-top.drawio.png" ></img></a>
	<a href="..\README.md"><img id="img-back" src="../_rsc/back.drawio.png"></img></a>
</div>
<a id="TOP"></a>

## Package [io/ioutil](https://pkg.go.dev/io/ioutil)


## Package io/ioutil (Deprecated, 1.16)

包 `ioutil` 实现了一些 I/O 实用功能函数。

已弃用：从 Go 1.16 开始，相同的功能现在由包 `io` 或包 `os` 提供，这些实现应该在新代码中优先使用。 

---
### var Discard (D, 1.16)

```go
var Discard io.Writer = io.Discard
```

>---
### func NopCloser (D, 1.16)

```go
func NopCloser(r io.Reader) io.ReadCloser {
	return io.NopCloser(r)
}
```

>---
### func ReadAll (D, 1.16)

```go
func ReadAll(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}
```

>---
### func ReadDir (D, 1.16)

```go
func ReadDir(dirname string) ([]fs.FileInfo, error)
```

`ReadDir` 读取由 `dirname` 命名的目录，并返回目录内容的 `fs.FileInfo` 列表，按文件名排序。如果在读取目录时发生错误，`ReadDir` 将不会返回任何目录条目。

已弃用：从 Go 1.16 开始，`os.ReadDir` 是一个更有效和正确的选择：它返回 `fs.DirEntry` 而不是 `fs.FileInfo` 的列表，并且在读取目录中途出错的情况下返回部分结果。

```go
func ReadDir(dirname string) ([]fs.FileInfo, error) {
	entries, err := os.ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			break
		}
		infos = append(infos, info)
	}
	return infos, err
}
```

>---
### func ReadFile (D, 1.16)

```go
func ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}
```

>---
### func TempDir (D, 1.17)

```go
func TempDir(dir, pattern string) (name string, err error) {
	return os.MkdirTemp(dir, pattern)
}
```

>---
### func TempFile (D, 1.17)

```go
func TempFile(dir, pattern string) (f *os.File, err error) {
	return os.CreateTemp(dir, pattern)
}
```

>---
### func WriteFile (D, 1.16)

```go
func WriteFile(filename string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(filename, data, perm)
}
```

---