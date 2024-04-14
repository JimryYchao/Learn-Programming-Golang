package main

import "fmt"

type IFace interface {
	Func()
}
type Embed struct{}

func (Embed) Func() {
	fmt.Print("Hello World")
}

type S struct {
	Embed
}

func foo(f IFace) {
	f.Func()
}

func main() {
	S.Func(S{})
}
