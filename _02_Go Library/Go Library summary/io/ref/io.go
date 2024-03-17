package io

import "errors"

const (
	SeekStart   = 0
	SeekCurrent = 1
	SeekEnd     = 2
)

var (
	EOF              = errors.New("EOF")
	ErrClosedPipe    = errors.New("io: read/write on closed pipe")
	ErrNoProgress    = errors.New("multiple Read calls return no data or error")
	ErrShortBuffer   = errors.New("short buffer")
	ErrShortWrite    = errors.New("short write")
	ErrUnexpectedEOF = errors.New("unexpected EOF")
)

type (
	Writer interface {
		Write(p []byte) (n int, err error)
	}
	Reader interface {
		Read(p []byte) (n int, err error)
	}
	WriterAt interface {
		WriteAt(p []byte, off int64) (n int, err error)
	}
	ReaderAt interface {
		ReadAt(p []byte, off int64) (n int, err error)
	}
	WriterTo interface {
		WriteTo(w Writer) (n int64, err error)
	}
	ReaderFrom interface {
		ReadFrom(r Reader) (n int64, err error)
	}
	RuneReader interface {
		ReadRune() (r rune, size int, err error)
	}
	RuneScanner interface {
		RuneReader
		UnreadRune() error
	}
	Closer interface {
		Close() error
	}
	Seeker interface {
		Seek(offset int64, whence int) (int64, error)
	}
	ByteWriter interface {
		WriteByte(c byte) error
	}
	ByteReader interface {
		ReadByte() (byte, error)
	}
	ByteScanner interface {
		ByteReader
		UnreadByte() error
	}
	StringWriter interface {
		WriteString(s string) (n int, err error)
	}
	ReadCloser interface {
		Reader
		Closer
	}
	ReadSeekCloser interface {
		Reader
		Seeker
		Closer
	}
	ReadSeeker interface {
		Reader
		Seeker
	}
	ReadWriteCloser interface {
		Reader
		Writer
		Closer
	}
	ReadWriteSeeker interface {
		Reader
		Writer
		Seeker
	}
	ReadWriter interface {
		Reader
		Writer
	}
	WriteCloser interface {
		Writer
		Closer
	}
	WriteSeeker interface {
		Writer
		Seeker
	}
)

var Discard Writer

func MultiWriter(writers ...Writer) Writer
func MultiReader(readers ...Reader) Reader
func TeeReader(r Reader, w Writer) Reader
func NopCloser(r Reader) ReadCloser

func Copy(dst Writer, src Reader) (written int64, err error)
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
func ReadAll(r Reader) ([]byte, error)
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
func ReadFull(r Reader, buf []byte) (n int, err error)
func WriteString(w Writer, s string) (n int, err error)

func Pipe() (*PipeReader, *PipeWriter)

type PipeReader struct{}

func (r *PipeReader) Close() error
func (r *PipeReader) CloseWithError(err error) error
func (r *PipeReader) Read(data []byte) (n int, err error)

type PipeWriter struct{}

func (w *PipeWriter) Close() error
func (w *PipeWriter) CloseWithError(err error) error
func (w *PipeWriter) Write(data []byte) (n int, err error)

type LimitedReader struct {
	R Reader
	N int64
}

func LimitReader(r Reader, n int64) Reader

func (l *LimitedReader) Read(p []byte) (n int, err error)

type OffsetWriter struct{}

func NewOffsetWriter(w WriterAt, off int64) *OffsetWriter
func (o *OffsetWriter) Seek(offset int64, whence int) (int64, error)
func (o *OffsetWriter) Write(p []byte) (n int, err error)
func (o *OffsetWriter) WriteAt(p []byte, off int64) (n int, err error)

type SectionReader struct{}

func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader
func (s *SectionReader) Outer() (r ReaderAt, off int64, n int64)
func (s *SectionReader) Read(p []byte) (n int, err error)
func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)
func (s *SectionReader) Seek(offset int64, whence int) (int64, error)
func (s *SectionReader) Size() int64
