## 输入输出：io/fs

包 `os/fs` 定义了文件系统的基本接口。文件系统可以由主机操作系统提供，也可以由其他包提供。

---
### var Err

文件系统错误，可以使用 `errors.Is` 针对文件系统返回的错误进行测试。

```go
var (
    ErrInvalid    = errInvalid()    // "invalid argument"
    ErrPermission = errPermission() // "permission denied"
    ErrExist      = errExist()      // "file already exists"
    ErrNotExist   = errNotExist()   // "file does not exist"
    ErrClosed     = errClosed()     // "file already closed"
)
```

>---
### var Skip

```go
var SkipAll = errors.New("skip everything and stop the walk")
var SkipDir = errors.New("skip this directory")
```

`SkipAll` 和 `SkipDir ` 作为函数 `WalkDirFunc` 的返回值。`SkipAll` 指示将跳过所有剩余的文件和目录。`SkipDir` 指示将跳过调用中命名的路径。它们不会被任何函数作为错误返回。

>---
### interface FS

```go
type FS interface {
    Open(name string) (File, error)
}
```

接口 `FS` 提供对层级文件系统的访问。`FS` 接口是文件系统所需的最低实现。文件系统可以实现额外的接口，例如 `ReadFileFS`，以提供额外的或优化的功能。它可以代表系统的指定路径下的单个文件，或是一个目录。

`Open` 打开指定路径。当 `Open` 返回一个错误时（例如指定的路径不存在或无法打开），它的类型应该是 `*PathError`, `Op` 字段设置为 `"open"`，`Path` 字段设置为 `name`, `Err` 字段描述问题。

`Open` 的参数 `name` 不能为 `""`，否则将返回 `nil` 和一个非空错误。`name` 为 `"."` 表示为 `FS` 的根。

`Open` 应该拒绝尝试打开不满足 `ValidPath(name)` 的 `name`，并返回一个 `*PathError`, `Err` 设置为 `ErrInvalid` 或 `ErrNotExist`。

>---

#### type PathError

```go
type PathError struct {
    Op   string
    Path string
    Err  error
}

func (e *PathError) Error() string
func (e *PathError) Timeout() bool
func (e *PathError) Unwrap() error
```

`PathError` 记录错误以及导致错误的操作和路径。当 `Fs.Open` 发生错误时，返回一个实现 `error` 接口的 `*PathError` 类型，其中 `Op` 字段设置为 `"open"`，`Path` 字段设置为 `name`, `Err` 字段设置为 `Open` 的错误描述。

`PathError.Error` 方法实现 `error` 接口，并将路径错误信息格式化返回。

`PathError.Timeout` 报告此错误是否表示超时。

`PathError.Unwrap` 返回未包装的底层 `error` 对象 `PathError.Err`。

```go
func main() {
	fileSys := os.DirFS(".")
	if _, err := fileSys.Open("WrongFile"); err != nil {
		e := err.(*fs.PathError)
		fmt.Println(e.Error())   // e.Op + " " + e.Path + ": " + e.Err.Error()
		fmt.Println(e.Timeout()) // false
		fmt.Println(e.Unwrap())  // e.Err
		fmt.Println(e)
		// open WrongFile: The system cannot find the file specified.
	}
}
```

>---
### func ValidPath

```go
func ValidPath(name string) bool
```

`ValidPath` 报告给定的路径名是否可用于 `Open` 调用。传递给 `Open` 的路径名是 UTF-8 编码的、无根的、斜杠分隔的路径元素序列，如 `"x/y/z"`。路径名不能包含 `"."`、`".."` 元素或空字符串 `""`，根路径名为 `"."` 的特殊情况除外（单个 `"."` 表示以当前文件系统 `FS` 的根路径）。路径不能以斜杠开头或结尾：`"/x"` 和 `"x/"` 无效。

路径在所有系统上都是用斜杠 `/` 分隔的，甚至在 Windows 上也是如此。包含其他字符（如反斜杠和冒号）的路径被认为是有效的，但这些字符决不能被 `FS` 实现解释为路径元素分隔符。

```go
func main() {
	path := "_02_Go Library/Go Library summary/io/ref"
	if fs.ValidPath(path) {
		if file, err := os.DirFS(path).Open("hello.file"); err != nil || file == nil {
			log.Fatalln(err)
		} else {
			io.Copy(os.Stdout, file) // Hello World!
		}
	}
}
```

>---
### interface File

```go
type File interface {
    Stat() (FileInfo, error)
    Read([]byte) (int, error)
    Close() error
}
```

`File` 提供对单个文件的访问。`File` 接口是文件所需的最低实现。目录文件也应该实现 `ReadDirFile`。`File` 也可以实现 `io.ReaderAt` 或 `io.Seeker` 作为优化扩展。

`Stat` 返回对当前文件的信息 `FileInfo`。

`Read` 读取当前打开的文件流，并保存在字节数组中。`Close` 用于关闭当前打开的文件流，之后的读取操作将返回错误。

```go
func main() {
	path := "_02_Go Library/Go Library summary/io/ref"
	if file, err := os.DirFS(path).Open("hello.file"); err != nil || file == nil {
		log.Fatalln(err)
	} else {
		buf := make([]byte, 512)
		if n, err := file.Read(buf); err == nil {
			info, _ := file.Stat()
			fmt.Printf("Read %d bytes in the file [%v]: %s", n, info.Name(), buf)
		} // Read 12 bytes in the file [hello.file]: Hello World!
	}
}
```

>---
### type FileMode

```go
type FileMode uint32
```

`FileMode` 表示文件的模式和权限位。这些位在所有系统上具有相同的定义，因此有关文件的信息可以从一个系统移动到另一个系统。并非所有位都适用于所有系统。唯一需要的位是目录的 `ModeDir`。

定义的文件模式位是 `FileMode` 的最高有效位。9 个最低有效位是标准的 Unix `rwxrwxrwx` 权限。这些位的值应该被认为是公共 API 的一部分，并且可以在有线协议或磁盘表示中使用：它们不能被更改，尽管可能会添加新的位。

```go
const (
    // 单个字母是 String 方法格式化时使用的缩写。
    ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
    ModeAppend                                     // a: append-only
    ModeExclusive                                  // l: exclusive use
    ModeTemporary                                  // T: temporary file; Plan 9 only
    ModeSymlink                                    // L: symbolic link
    ModeDevice                                     // D: device file
    ModeNamedPipe                                  // p: named pipe (FIFO)
    ModeSocket                                     // S: Unix domain socket
    ModeSetuid                                     // u: setuid
    ModeSetgid                                     // g: setgid
    ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
    ModeSticky                                     // t: sticky
    ModeIrregular                                  // ?: non-regular file; nothing else is known about this file

    // Mask for the type bits. For regular files, none will be set.
    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular

    ModePerm FileMode = 0777 // Unix permission bits
)
```

以下方法是对 `FileMode` 值的校验方法：

```go
func (m FileMode) IsDir() bool
func (m FileMode) IsRegular() bool
func (m FileMode) Perm() FileMode
func (m FileMode) Type() FileMode
func (m FileMode) String() string
```

`m.IsDir` 报告 `m` 是否描述为目录。也就是说，它测试 `m` 是否设置了 `ModeDir` 位。

`m.IsRegular` 报告 `m` 是否描述为常规文件。也就是说，它测试是否没有设置模式类型位。

`m.Perm` 返回 `m` 中的 Unix 权限位 `m & ModePerm`。

`m.Type` 返回 `m` 中的类型位 `m & ModeType`。

`m.String` 返回 `m` 的完整的文件模式和权限位。

```go
func main() {
	path := "_02_Go Library/Go Library summary/io/ref"

	if info, err := fs.Stat(os.DirFS(path), "."); err == nil {
		mode := info.Mode()
		fmt.Printf("IsDir : %v\n", mode.IsDir())
		fmt.Printf("IsRegular : %v\n", mode.IsRegular())
		fmt.Printf("Perm : %v\n", mode.Perm())
		fmt.Printf("String : %v\n", mode.String())
		fmt.Printf("Type : %v\n", mode.Type())
	}
}
/*
IsDir : true
IsRegular : false
Perm : -rwxrwxrwx
String : drwxrwxrwx
Type : d---------
*/
```

>---

### interface FileInfo

```go
type FileInfo interface {
    Name() string       // base name of the file
    Size() int64        // length in bytes for regular files; system-dependent for others
    Mode() FileMode     // file mode bits
    ModTime() time.Time // modification time
    IsDir() bool        // abbreviation for Mode().IsDir()
    Sys() any           // underlying data source (can return nil)
}
```

`FileInfo` 描述给定路径的文件或目录信息，并由 `FS.Stat` 返回：
- `Name` 返回文件或目录的名称。
- `Size` 返回常规文件的字节大小。
- `Mode` 返回文件的模式和权限位。
- `ModTime` 返回最近一次修改时间。
- `IsDir` 检查当前是否为目录。
- `Sys` 返回系统的底层数据源。

>---
### interface DirEntry

```go
type DirEntry interface {

    Name() string
    IsDir() bool
    Type() FileMode
    Info() (FileInfo, error)
}
```

`DirEntry` 是从目录中读取的（使用 `ReadDir` 函数或 `ReadDirFile` 的 `ReadDir` 方法）。

`DirEntry.Name` 返回条目 `DirEntry` 所描述的文件（或子目录）的名称。该名称只是路径的最后一个元素（基本名称），而不是整个路径。例如，`Name` 将返回 `"hello.go"` 而不是 `"home/gopher/hello.go"`。

`DirEntry.IsDir` 报告条目是否描述了一个目录。

`DirEntry.Type` 返回该项的类型位。类型位是通常 `FileMode` 位的子集，由 `FileMode.Type` 返回。

`DirEntry.Info` 返回该条目描述的文件或子目录的 `FileInfo`。返回的 `FileInfo` 可以是从读取原始目录的时间开始，也可以是从调用 `Info` 的时间开始。如果文件在读取目录后被删除或重命名，`Info` 可能会返回一个错误的描述（例如 `ErrNotExist`）。如果条目表示一个符号链接，`Info` 报告关于链接本身的信息，而不是链接的目标。

>---
### func FileInfoToDirEntry

```go
func FileInfoToDirEntry(info FileInfo) DirEntry
```

`FileInfoToDirEntry` 返回一个从 `info` 返回信息的 `DirEntry`。如果 `info` 为 `nil`，则 `FileInfoToDirEntry` 返回 `nil`。

```go
// 查找子目录
func main() {
	dirPath := "_02_Go Library/Go Library summary/io"
	fsys := os.DirFS(dirPath)

	if files, err := fs.Glob(fsys, "*"); err == nil {
		for _, f := range files {
			if file, _ := fsys.Open(f); file != nil {
				defer file.Close()
				info, _ := file.Stat()
				if dir := fs.FileInfoToDirEntry(info); dir != nil && dir.IsDir() {
					fmt.Printf("Dir : %s\n", dir.Name())
				}
			}
		}
	}
	// Dir : ref
}
```

>---
### func FormatFileInfo

```go
func FormatFileInfo(info FileInfo) string
```

`FormatFileInfo` 返回 `info` 的人类可读的格式化版本。`FileInfo` 的实现可以从 `String` 方法调用它。

```go
func main() {
	path := "_02_Go Library/Go Library summary/io/ref"
	if file, err := os.DirFS(path).Open("hello.file"); err != nil || file == nil {
		log.Fatalln(err)
	} else {
		finfo, _ := file.Stat()
		fmt.Printf("FORMAT: %s\n", fs.FormatFileInfo(finfo))
		fmt.Printf("FileInfo : \n\tName : %s\n\tSize : %d\n\tMode : %v\n\tModTime : %v\n\tIsDir : %v\n\tSys : %v",
			finfo.Name(), finfo.Size(), finfo.Mode(), finfo.ModTime(), finfo.IsDir(), finfo.Sys())
	}
}
/*
FORMAT: -rw-rw-rw- 12 2024-03-17 22:35:52 hello.file
FileInfo :
        Name : hello.file
        Size : 12
        Mode : -rw-rw-rw-
        ModTime : 2024-03-17 22:35:52.4547126 +0800 CST
        IsDir : false
        Sys : &{32 {1699659782 31094904} {3733097101 31094925} {1772287542 31094904} 0 12}
*/
```

>---
### func FormatDirEntry

```go
func FormatDirEntry(dir DirEntry) string
```

`FormatDirEntry` 返回 `dir` 人类可读的格式化版本。`DirEntry` 的实现可以从 `String` 方法调用它。

```go
func main() {
	dirPath := "_02_Go Library/Go Library summary/io"
	fsys := os.DirFS(dirPath)
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			fmt.Printf("Dir : %v\n", fs.FormatDirEntry(d))
		} 
		return err
	})
}

/*
Dir : d ./
Dir : d ref/
*/
```

>---
### interface GlobFs

```go
type GlobFS interface {
    FS
    Glob(pattern string) ([]string, error)
}
```

`GlobFS` 是带有一个 `Glob` 方法的文件系统。`Glob` 返回匹配 `pattern` 的所有文件的名称，以提供顶级 `Glob` 函数的实现。 

>---
### func Glob

```go
func Glob(fsys FS, pattern string) (matches []string, err error)
```

`Glob` 返回所有匹配 `pattern` 的文件或路径，如果没有匹配的名称，则返回 `nil`。模式的语法与 `path.Match` 中的语法相同。该模式可以描述诸如 `usr/*/bin/ed` 之类的分层名称。

`Glob` 忽略文件系统错误，例如 I/O 错误（读取目录）。唯一可能返回的错误是 `path.ErrBadPattern`，报告模式格式不正确。

如果 `fsys` 实现 `GlobFS`，则 `Glob` 调用 `fsys.Glob`。否则，`Glob` 使用 `ReadDir` 遍历路径树并查找模式的匹配项。

```go
func main() {
	path := "_02_Go Library/Go Library summary/io"
	if fs.ValidPath(path) {
		fileSys := os.DirFS(path)
		if files, err := fs.Glob(fileSys, "*.md"); err == nil {
			for _, f := range files {
				fmt.Println(f)
			}
		}
	}
}
/*
Summary.md
io.md
io_fs.md
ioutil.md
*/
```

>---
### interface StatFS

```go
type StatFS interface {
    FS
    Stat(name string) (FileInfo, error)
}
```

`StatFS` 是一个带有 `Stat` 方法的文件系统。`Stat` 返回描述指定路径的 `FileInfo`。如果有错误，它的类型应该是 `*PathError`。

>---
### func Stat

```go
func Stat(fsys FS, name string) (FileInfo, error)
```

`Stat` 从文件系统返回 `name` 路径描述的 `FileInfo`。如果 `fsys` 实现 `StatFS`，`Stat` 将调用 `fsys.Stat`。否则，`Stat` 将尝试打开文件并检查其状态。

```go
func main() {
	path := "_02_Go Library/Go Library summary/io"
	if fs.ValidPath(path) {
		fileSys := os.DirFS(path)
		if files, err := fs.Glob(fileSys, "*"); err == nil {
			var info fs.FileInfo
			for _, fname := range files {
				if info, _ = fs.Stat(fileSys, fname); info != nil {
					fmt.Println(fs.FormatFileInfo(info))
				}
			}
		}
	}
}
/*
-rw-rw-rw- 0 2024-03-17 11:23:00 Summary.md
-rw-rw-rw- 24535 2024-03-18 01:02:09 io.md
-rw-rw-rw- 16686 2024-03-18 03:12:25 io_fs.md
-rw-rw-rw- 0 2024-03-17 10:09:01 ioutil.md
drwxrwxrwx 0 2024-03-17 23:58:44 ref/
*/
```

>---
### interface ReadFileFS 

```go
type ReadFileFS interface {
    FS
    ReadFile(name string) ([]byte, error)
}
```

`ReadFileFS` 是由文件系统实现的接口，它提供 `ReadFile` 的优化实现。

`ReadFile` 读取 `name` 文件并返回其内容。一个成功的调用返回 `nil` 错误，而不是 `io.EOF`。因为 `ReadFile` 读取整个文件，所以最终 `Read` 的预期 `EOF` 不会被视为要报告的错误。允许调用方修改返回的字节片段。此方法应返回底层数据的副本。


>---
### func ReadFile

```go
func ReadFile(fsys FS, name string) ([]byte, error)
```

`ReadFile` 从文件系统 `fsys` 中读取指定 `name` 的文件并返回其内容。成功的调用返回 `nil` 错误，而不是 `io.EOF`。（由于 `ReadFile` 读取整个文件，因此最终读取的预期 `EOF` 不会被视为要报告的错误。）

如果 `fsys` 实现接口 `ReadFileFS`，则 `ReadFile` 调用 `fsys.ReadFile`。否则 `ReadFile` 调用 `fsys.Open` 并对返回的 `File` 使用 `Read` 和 `Close`。

```go
func main() {
	fileSystem := os.DirFS("./_02_Go Library/Go Library summary/io/ref")

	if bytes, err := fs.ReadFile(fileSystem, "hello.file"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s\n", bytes) // Hello World!
	}
}
```

>---
### interface SubFS

```go
type SubFS interface {
    FS
    Sub(dir string) (FS, error)
}
```

`SubFS` 是一个具有 `Sub` 方法的文件系统。

`SubFS.Sub` 返回一个对应于以 `dir` 为根的子树的 `FS`。

>---
### func Sub

```go
func Sub(fsys FS, dir string) (FS, error)
```

`Sub` 返回一个对应于 `fsys` 目录下以 `dir` 为根的子树的文件系统 `FS`。

如果 `dir` 是 `"."`，`Sub` 返回 `fsys` 不变。或是，如果 `fs` 实现 `SubFS`，`Sub` 将返回 `fsys.Sub(dir)`。或是，`Sub` 返回一个实现 `sub` 的新的 `FS`，它实际上将 `sub.Open(name)` 实现为 `fsys.Open(path.Join(dir，name))`。该实现还适当地转换对 `ReadDir`、`ReadFile` 和 `Glob` 的调用。

`Sub(os.DirFS("/"),"prefix")` 等效于 `os.DirFS(“/prefix”)`，它们都不能保证避免操作系统访问 `"/prefix"` 之外的目录，因为 `os.DirFS` 的实现不会检查 `"/prefix"` 内部指向其他目录的符号链接。也就是说，`os.DirFS` 不是 *chroot-style* 安全机制的一般替代品，`Sub` 也不会改变这一事实。

```go
func main() {
	path := "_02_Go Library/Go Library summary/io"

	subfs, _ := fs.Sub(os.DirFS(path), ".")

	fs.WalkDir(subfs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			fmt.Printf("Dir : %s\n", path)
		} else {
			fmt.Printf("File : %s\n", path)
		}
		return err
	})
}
```

>---
### interface ReadDirFS
  
```go
type ReadDirFS interface {
    FS
    ReadDir(name string) ([]DirEntry, error)
}
```
`ReadDirFS` 是由文件系统实现的接口，该文件系统提供 `ReadDir` 的优化实现。

`ReadDirFS.ReadDir` 读取 `name` 目录并返回按文件名排序的目录项列表。

>---
### func ReadDir

```go
func ReadDir(fsys FS, name string) ([]DirEntry, error)
```

`ReadDir` 读取指定的目录并返回按文件名排序的目录条目列表。如果 `fsys` 实现 `ReadDirFS`，则 `ReadDir` 调用 `fsys.ReadDir`。否则 `ReadDir` 调用 `fsys.Open` 并对返回的文件使用 `ReadDir` 和 `Close`。

```go
func main() {
	dirPath := "_02_Go Library/Go Library summary/"
	if dirs, err := fs.ReadDir(os.DirFS(dirPath), "."); err == nil {
		for _, v := range dirs {
			fmt.Println(fs.FormatDirEntry(v))
		}
	}
}
/*
- Summary.md
- io.md
- io_fs.md
- ioutil.md
d ref/
*/
```

>---
### interface ReadDirFile

```go
type ReadDirFile interface {
    File
    ReadDir(n int) ([]DirEntry, error)
}
```

`ReadDirFile` 是一个目录文件，其条目可以用 `ReadDir` 方法读取。每个目录文件都应该实现这个接口。它允许任何文件实现此接口，但如果是这样，`ReadDir` 应该返回非目录的错误。

`ReadDir` 读取目录的内容，并按目录顺序返回最多 `n` 个 `DirEntry` 值的切片。对同一 `File` 的后续调用将产生更多的 `DirEntry` 值。

如果 `n > 0`，`ReadDir` 最多返回 `n` 个 `DirEntry` 结构。在这种情况下，如果 `ReadDir` 返回空切片，它将返回一个非 `nil` 错误来解释原因。到达在目录末尾，错误是 `io.EOF`。`ReadDir` 必须返回 `io.EOF` 本身，而不是包装 `io.EOF` 的错误。

如果 `n <= 0`，`ReadDir` 将在单个切片中返回目录中所有的 `DirEntry` 值。在这种情况下，如果 `ReadDir` 成功（一直读取到目录的末尾），它将返回切片和 `nil` 错误。如果在目录结束之前遇到错误，`ReadDir` 将返回到该点之前读取的 `DirEntry` 列表和一个非 `nil` 错误。

```go
func main() {
	var fsdir fs.ReadDirFile
	var err error
	if fsdir, err = os.Open("_02_Go Library/Go Library summary/io"); err == nil {
		defer fsdir.Close()
		entries, _ := fsdir.ReadDir(-1)
		for _, entry := range entries {
			fmt.Println(fs.FormatDirEntry(entry))
		}
	}
}
```

>---
### funcType WalkDirFunc

```go
type WalkDirFunc func(path string, d DirEntry, err error) error
```

`WalkDirFunc` 是 `WalkDir` 调用的函数类型，用于访问每个文件或目录。

`path` 参数包含 `WalkDir` 的参数作为前缀。也就是说，如果使用根参数 `"dir"` 调用 `WalkDir`，并在该目录中找到名为 `"a"` 的文件，则将使用参数 `"dir/a"` 调用 `walk` 函数。

参数 `d` 是命名路径的 `DirEntry`。

函数返回的错误结果控制 `WalkDir` 如何继续。如果函数返回特殊值 `SkipDir`，`WalkDir` 将跳过当前目录（如果 `d.IsDir()` 为 `true`，则为 `path`，否则为 `path` 的父目录）。如果函数返回特殊值 `SkipAll`，`WalkDir` 将跳过所有剩余的文件和目录。否则，如果函数返回非 `nil` 错误，`WalkDir` 将完全停止并返回该错误。

`err` 参数报告一个与 `path` 相关的错误，表示 `WalkDir` 不会进入该目录。该函数可以决定如何处理该错误；如前所述，返回错误将导致 `WalkDir` 停止遍历整个树。

`WalkDir` 在两种情况下使用非 `nil` 参数 `err` 调用函数：
- 首先，如果根目录上的初始 `Stat` 失败，`WalkDir` 将调用函数，其中 `path` 设置为 `root`，`d` 设置为 `nil`，`err` 设置为 `fs.Stat` 中的错误。

+ 其次，如果目录的 `ReadDir` 方法（请参阅 `ReadDirFile`）失败，`WalkDir` 将调用该函数，其中 `path` 设置为目录的路径，`d` 设置为描述目录的 `DirEntry`，`err` 设置为 `ReadDir` 的错误。在第二种情况下，使用目录的路径调用该函数两次：第一次调用是在尝试读取目录之前，并将 `err` 设置为 `nil`，使函数有机会返回 `SkipDir` 或 `SkipAll` 并完全避免使用 `ReadDir`。第二个调用发生在 `ReadDir` 失败之后，并报告来自 `ReadDir` 的错误。（如果 `ReadDir` 成功，则没有第二次调用。）

`WalkDirFunc` 与 `path/filepath.WalkFunc` 的区别在于：
- 第二个参数的类型是 `DirEntry` 而不是 `FileInfo`。
- 在读取目录之前调用该函数，以允许 `SkipDir` 或 `SkipAll` 绕过完全读取的目录或分别跳过所有剩余的文件和目录。
- 如果目录读取失败，则会再次调用该函数，以使该目录报告错误。

>---
### func WalkDir

```go
func WalkDir(fsys FS, root string, fn WalkDirFunc) error
```

`WalkDir` 遍历以 `root` 为根的文件树，为树中的每个文件或目录（包括 `root`）调用 `fn`。

访问文件和目录时出现的所有错误都由 `fn` 过滤。这些文件是按词法顺序遍历的，这使得输出具有确定性，但要求 `WalkDir` 在继续遍历该目录之前将整个目录读入内存。`WalkDir` 不会跟踪目录中的符号链接，但如果根目录本身是一个符号链接，则会遍历其目标。

```go
func main() {
	root := "_02_Go Library/Go Library summary/io"
	if fs.ValidPath(root) {
		fileSys := os.DirFS(root)

		fs.WalkDir(fileSys, ".", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				log.Fatal(err)
			}
			// 输出该目录下的子目录和文件
			fmt.Println(path)
			return nil
		})
	}
}
```
