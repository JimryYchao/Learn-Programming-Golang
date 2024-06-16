package main

import "C"
import "fmt"

// set GOARCH=386
// set CGO_ENABLED=1
// go build -ldflags "-s -w" -buildmode=c-shared -o CSgo.Interop.dll netCgo.go

func main() {
	Hello()
}

//export Hello
func Hello() {
	fmt.Println("Hello CSharp, I am Go!")
}

//export Println
func Println(format string, args []any) {
	fmt.Printf(format+"\n", args...)
}

//export IterIntSlice
func IterIntSlice(is []int) {
	for _, i := range is {
		fmt.Println(i)
	}
}
