package gostd

import (
	"bytes"
	"compress/bzip2"
	"io"
	"os"
	"testing"
)

// ! NewReader 从 r 返回一个 Bzip2 解压缩 reader
// ? go test -v -run=^TestBzip2$
func TestBzip2(t *testing.T) {
	decompress(t, "e.txt.bz2", "e.txt")
	decompress(t, "Isaac.Newton-Opticks.txt.bz2", "Isaac.Newton-Opticks.txt")
}

func decompress(t *testing.T, file string, uncpName string) {
	bz2f, err := os.ReadFile("testdata/" + file)
	if err != nil {
		t.Fatal(err)
	}

	bzr := bzip2.NewReader(bytes.NewReader(bz2f))

	wr, err := os.OpenFile("testdata/"+uncpName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil || wr == nil {
		t.Fatal(err)
	}

	if _, err := io.Copy(wr, bzr); err != nil {
		t.Fatal(err)
	}
	wr.Close()
}
