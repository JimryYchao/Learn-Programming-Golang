package ref

import (
	"errors"
	"io/fs"
	"time"
)

var (
	ErrInvalid    = fs.ErrInvalid    // "invalid argument"
	ErrPermission = fs.ErrPermission // "permission denied"
	ErrExist      = fs.ErrExist      // "file already exists"
	ErrNotExist   = fs.ErrNotExist   // "file does not exist"
	ErrClosed     = fs.ErrClosed     // "file already closed"
)

var SkipAll = errors.New("skip everything and stop the walk")
var SkipDir = errors.New("skip this directory")

type FileMode uint32

func (m FileMode) IsDir() bool
func (m FileMode) IsRegular() bool
func (m FileMode) Perm() FileMode
func (m FileMode) Type() FileMode
func (m FileMode) String() string

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

type (
	FileInfo interface {
		Name() string       // base name of the file
		Size() int64        // length in bytes for regular files; system-dependent for others
		Mode() FileMode     // file mode bits
		ModTime() time.Time // modification time
		IsDir() bool        // abbreviation for Mode().IsDir()
		Sys() any           // underlying data source (can return nil)
	}
	File interface {
		Stat() (FileInfo, error)
		Read([]byte) (int, error)
		Close() error
	}
	FS interface {
		Open(name string) (File, error)
	}
	DirEntry interface {
		Name() string
		IsDir() bool
		Type() FileMode
		Info() (FileInfo, error)
	}
	GlobFS interface {
		FS
		Glob(pattern string) ([]string, error)
	}
	StatFS interface {
		FS
		Stat(name string) (FileInfo, error)
	}
	ReadFileFS interface {
		FS
		ReadFile(name string) ([]byte, error)
	}
	SubFS interface {
		FS
		Sub(dir string) (FS, error)
	}
	ReadDirFS interface {
		FS
		ReadDir(name string) ([]DirEntry, error)
	}
	ReadDirFile interface {
		File
		ReadDir(n int) ([]DirEntry, error)
	}
)

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string
func (e *PathError) Timeout() bool
func (e *PathError) Unwrap() error

func ValidPath(name string) bool
func FileInfoToDirEntry(info FileInfo) DirEntry
func FormatFileInfo(info FileInfo) string
func FormatDirEntry(dir DirEntry) string
func Glob(fsys FS, pattern string) (matches []string, err error)
func Stat(fsys FS, name string) (FileInfo, error)
func ReadFile(fsys FS, name string) ([]byte, error)
func Sub(fsys FS, dir string) (FS, error)
func ReadDir(fsys FS, name string) ([]DirEntry, error)

type WalkDirFunc func(path string, d DirEntry, err error) error

func WalkDir(fsys FS, root string, fn WalkDirFunc) error
