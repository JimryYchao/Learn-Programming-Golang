package ref

import (
	"errors"
	"io"
)

// ===== Scanner ===================================================

const MaxScanTokenSize = 64 * 1024

type Scanner struct {
	// contains filtered or unexported fields
}

var (
	ErrTooLong         = errors.New("bufio.Scanner: token too long")
	ErrNegativeAdvance = errors.New("bufio.Scanner: SplitFunc returns negative advance count")
	ErrAdvanceTooFar   = errors.New("bufio.Scanner: SplitFunc returns advance count beyond input")
	ErrBadReadCount    = errors.New("bufio.Scanner: Read returned impossible count")
)

var ErrFinalToken = errors.New("final token")

type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

func NewScanner(r io.Reader) *Scanner

func (s *Scanner) Buffer(buf []byte, max int)
func (s *Scanner) Bytes() []byte
func (s *Scanner) Err() error
func (s *Scanner) Scan() bool
func (s *Scanner) Split(split SplitFunc)
func (s *Scanner) Text() string

var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)

func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)

// ===== Reader ===================================================

type Reader struct {
	// contains filtered or unexported fields
}

func NewReaderSize(rd io.Reader, size int) *Reader
func NewReader(rd io.Reader) *Reader
func (b *Reader) Size() int
func (b *Reader) Reset(r io.Reader)
func (b *Reader) Peek(n int) ([]byte, error)
func (b *Reader) Discard(n int) (discarded int, err error)
func (b *Reader) Read(p []byte) (n int, err error)
func (b *Reader) ReadByte() (byte, error)
func (b *Reader) UnreadByte() error
func (b *Reader) ReadRune() (r rune, size int, err error)
func (b *Reader) UnreadRune() error
func (b *Reader) Buffered() int
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
func (b *Reader) ReadBytes(delim byte) ([]byte, error)
func (b *Reader) ReadString(delim byte) (string, error)
func (b *Reader) WriteTo(w io.Writer) (n int64, err error)

// ===== Writer ===================================================

type Writer struct {
	// contains filtered or unexported fields
}

func NewWriterSize(w io.Writer, size int) *Writer
func NewWriter(w io.Writer) *Writer
func (b *Writer) Size() int
func (b *Writer) Reset(w io.Writer)
func (b *Writer) Flush() error
func (b *Writer) Available() int
func (b *Writer) AvailableBuffer() []byte
func (b *Writer) Buffered() int
func (b *Writer) Write(p []byte) (nn int, err error)
func (b *Writer) WriteByte(c byte) error
func (b *Writer) WriteRune(r rune) (size int, err error)
func (b *Writer) WriteString(s string) (int, error)
func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)

// ===== ReadWriter ===================================================

type ReadWriter struct {
	*Reader
	*Writer
}

func NewReadWriter(r *Reader, w *Writer) *ReadWriter
