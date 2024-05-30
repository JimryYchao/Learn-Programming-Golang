package gostd

import (
	"embed"
	"io/fs"
	"testing"
)

//go:embed img.png
var img []byte

// ? go test -v -run=^TestEmbedImg$
func TestEmbedImg(t *testing.T) {
	log(len(img))
}

//go:embed hello.txt
var hello string

// ? go test -v -run=^TestEmbedToString$
func TestEmbedToString(t *testing.T) {
	log(hello)
}

//go:embed \\
var embedfs embed.FS

// ? go test -v -run=^TestEmbedFS$
func TestEmbedFS(t *testing.T) {
	if dirs, err := fs.ReadDir(embedfs, "."); err == nil {
		for _, d := range dirs {
			log(d.Name())
		}
	}
}
