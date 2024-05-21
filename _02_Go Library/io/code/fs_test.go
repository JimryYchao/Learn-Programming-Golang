package gostd

import (
	"io"
	"io/fs"
	"os"
	"testing"
)

var (
	dirPath  = "."
	filePath = "./fs_testing.file"
	fileName = "fs_testing.file"
)

func getRootFS() fs.FS {
	return os.DirFS(".")
}

func readFsFile(f fs.File) {
	readToStdout(f)
}

/*
! ValidPath 报告给定的路径名是否有效，有效时可用于 `Open` 调用。
! FS 提供对分层文件系统的访问。`FS` 接口是文件系统所需的最低实现。
	Open 以打开 `name` 指定的文件或目录，并返回一个 `File`。发生错误时返回一个 `*PathError` 类型的 `error`。
	PathError 记录错误以及导致错误的操作和文件路径。
! File 提供对单个文件的访问。`File` 接口是文件所需的最低实现。
*/
//? go test -v -run=^TestFSOpenFile$
func TestFSOpenFile(t *testing.T) {
	beforeTest(t)
	if fs.ValidPath(dirPath) {
		fsys := os.DirFS(dirPath)
		// 访问 File
		if file, err := fsys.Open(fileName); err == nil {
			file.Read(buf)
			defer file.Close()

			logf("Read file : %s", buf)
			info, _ := file.Stat()
			logf("Info of file : %s", fs.FormatFileInfo(info))
		} else {
			if pathErr, ok := err.(*fs.PathError); ok {
				t.Fatal(pathErr.Err)
			}
		}
	}
}

/*
! Stat 从文件系统 `fsys` 中返回描述 `name` 文件的 `FileInfo`。
	FileInfo 描述文件信息。
	StatFS 是一个带有 `Stat` 方法的文件系统。
*/
//?
func TestGetStatFS(t *testing.T) {
	s := []io.Writer{os.Stdout}
	wr := io.MultiWriter(s...)

	s[0] = nil
	wr.Write([]byte(hello))
}

/*


! Glob 返回指定文件系统 `fsys` 中所有匹配 `pattern` 的文件名称。
	GlobFs 是一个带有 `Glob` 方法的文件系统。

! Sub 返回一个对应于 `fsys` 系统下目录 `dir` 的子树文件系统。
	SubFS 是一个具有 `Sub` 方法的文件系统。

! ReadFile 从文件系统 `fsys` 中读取指定的 `name` 文件并返回其内容。
	ReadFileFS 是由文件系统实现的接口，它提供 `ReadFile` 的优化实现。

! ReadDir 读取指定的目录并返回按文件名排序的目录 `DirEntry` 条目列表。
	DirEntry 是目录中读取的条目。
    ReadDirFS 是由文件系统实现的接口，该文件系统提供 `ReadDir` 的优化实现。
	ReadDirFile 是一个目录文件，其目录条目可以用其 ReadDir 方法读取。

! WalkDir 遍历以 `root` 为根的文件树，为树中的每个文件或目录（包括 `root`）调用 `WalkDirFunc fn`。
	WalkDirFunc 是 `WalkDir` 调用的函数类型，用于访问每个文件或目录。

! FileInfoToDirEntry 返回一个从 `FileInfo` 返回信息的 `DirEntry`。
! FormatDirEntry 返回 `DirEntry` 格式化的字符串。
! FormatFileInfo 返回 `FileInfo` 格式化的字符串。







! FileMode 表示文件的模式和权限位。这些位在所有的系统上都有相同的定义，
	IsDir 描述是否为目录。
	IsRegular 描述是否为常规文件，即没有被设置模式类型位。
	Perm 返回 `fileMode` 的 Unix 权限位。
	String
	Type 返回 `fileMode & fs.ModeType` 中的类型位。


!

!




*/

// var SkipAll = errors.New("skip everything and stop the walk")
// var SkipDir = errors.New("skip this directory")

// type FileMode uint32

// func (m FileMode) IsDir() bool
// func (m FileMode) IsRegular() bool
// func (m FileMode) Perm() FileMode
// func (m FileMode) Type() FileMode
// func (m FileMode) String() string

// const (
// 	// 单个字母是 String 方法格式化时使用的缩写。
// 	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
// 	ModeAppend                                     // a: append-only
// 	ModeExclusive                                  // l: exclusive use
// 	ModeTemporary                                  // T: temporary file; Plan 9 only
// 	ModeSymlink                                    // L: symbolic link
// 	ModeDevice                                     // D: device file
// 	ModeNamedPipe                                  // p: named pipe (FIFO)
// 	ModeSocket                                     // S: Unix domain socket
// 	ModeSetuid                                     // u: setuid
// 	ModeSetgid                                     // g: setgid
// 	ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
// 	ModeSticky                                     // t: sticky
// 	ModeIrregular                                  // ?: non-regular file; nothing else is known about this file

// 	// Mask for the type bits. For regular files, none will be set.
// 	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular

// 	ModePerm FileMode = 0777 // Unix permission bits
// )

// type (
// 	FileInfo interface {
// 		Name() string       // base name of the file
// 		Size() int64        // length in bytes for regular files; system-dependent for others
// 		Mode() FileMode     // file mode bits
// 		ModTime() time.Time // modification time
// 		IsDir() bool        // abbreviation for Mode().IsDir()
// 		Sys() any           // underlying data source (can return nil)
// 	}
// 	File interface {
// 		Stat() (FileInfo, error)
// 		Read([]byte) (int, error)
// 		Close() error
// 	}
// 	FS interface {
// 		Open(name string) (File, error)
// 	}
// 	DirEntry interface {
// 		Name() string
// 		IsDir() bool
// 		Type() FileMode
// 		Info() (FileInfo, error)
// 	}
// 	GlobFS interface {
// 		FS
// 		Glob(pattern string) ([]string, error)
// 	}
// 	StatFS interface {
// 		FS
// 		Stat(name string) (FileInfo, error)
// 	}
// 	ReadFileFS interface {
// 		FS
// 		ReadFile(name string) ([]byte, error)
// 	}
// 	SubFS interface {
// 		FS
// 		Sub(dir string) (FS, error)
// 	}
// 	ReadDirFS interface {
// 		FS
// 		ReadDir(name string) ([]DirEntry, error)
// 	}
// 	ReadDirFile interface {
// 		File
// 		ReadDir(n int) ([]DirEntry, error)
// 	}
// )

// type PathError struct {
// 	Op   string
// 	Path string
// 	Err  error
// }

// func (e *PathError) Error() string
// func (e *PathError) Timeout() bool
// func (e *PathError) Unwrap() error

// func ValidPath(name string) bool
// func FileInfoToDirEntry(info FileInfo) DirEntry
// func FormatFileInfo(info FileInfo) string
// func FormatDirEntry(dir DirEntry) string
// func Glob(fsys FS, pattern string) (matches []string, err error)
// func Stat(fsys FS, name string) (FileInfo, error)
// func ReadFile(fsys FS, name string) ([]byte, error)
// func Sub(fsys FS, dir string) (FS, error)
// func ReadDir(fsys FS, name string) ([]DirEntry, error)

// type WalkDirFunc func(path string, d DirEntry, err error) error

// func WalkDir(fsys FS, root string, fn WalkDirFunc) error
