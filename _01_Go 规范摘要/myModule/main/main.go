package main

import "runtime/debug"

type (
	IFace interface {
		Fun()
	}

	S1 struct{}
	S2 struct {
		S1
	}
	S3 S2
	S4 = S1
)

func (S1) Fun() {
	test(S3{})
}

func test(f IFace) {
	f.Fun()
}

func main() {
	debug.SetGCPercent(100)
}
