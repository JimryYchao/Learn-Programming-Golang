package test

import (
	"fmt"
	"testing"
)

// 异步
func sliceToChan(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func TestChannel(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 66}
	dataChan := sliceToChan(nums)

	finalChan := sq(dataChan)

	for n := range finalChan {
		fmt.Println(n)
	}
}
