package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Printf("Hello World")
	}()
	time.Sleep(50000)
}
