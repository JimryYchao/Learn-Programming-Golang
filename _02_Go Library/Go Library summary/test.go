package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}
type WriterAt interface {
	WriteAt(p []byte, off int64) (n int, err error)
}
type WriterTo interface {
	WriteTo(w Writer) (n int64, err error)
}
type ByteWriter interface {
	WriteByte(c byte) error
}

type Point2D struct {
	X, Y int
}
type Point3D struct {
	Point2D
	Z int
}

func ScanLines(n int) {
	var Origin3D = Point3D{Point2D: Point2D{0, 0}, Z: 0}
	var p = Origin3D
	p.X += 1 // 提升字段 X, 等效于 p.Point2D.X
	p.Y += 1 // 提升字段 Y

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
	{
		// ...
		goto Error
		// ...
	Error:
	}
}

func main() {
	return
	const input = "1,2,3,4,5,6,7,8,9"

	file, _ := os.Open("_02_Go Library/Go Library summary/bufio/ref/readline.file")
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// for i := 0; i < len(data); i++ {
		// 	if data[i] == ',' {
		// 		return i + 1, data[:i], nil
		// 	}
		// }
		// println(data, atEOF)
	scan:
		advance, token, err = bufio.ScanWords(data, atEOF)
		// println(advance, token, err, atEOF)
		if err == nil && token != nil {
			// println(string(token))
			if _, err = strconv.ParseInt(string(token), 10, 32); err != nil {
				// println("!!!")
				err = nil
				data = data[advance:]
				// advance = len(string(token)) + 1
				// token = nil
				goto scan
			}
		}
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(onComma)
	for scanner.Scan() {
		fmt.Printf("Read Integer : %s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
