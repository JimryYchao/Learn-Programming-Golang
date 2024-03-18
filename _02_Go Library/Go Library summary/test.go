package main

import (
	"fmt"
	"io/fs"
	"os"
)

// type cusReadDirFile struct {
// 	file *os.File
// }

// func (f *cusReadDirFile) ReadDir(n int) ([]fs.DirEntry, error) {
// 	return f.file.ReadDir(n)
// }
// func (f *cusReadDirFile) Close() error {
// 	return f.file.Close()
// }
// func (f *cusReadDirFile) Read(data []byte) (int, error) {
// 	return f.file.Read(data)
// }
// func (f *cusReadDirFile) Stat() (fs.FileInfo, error) {
// 	return f.file.Stat()
// }

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
