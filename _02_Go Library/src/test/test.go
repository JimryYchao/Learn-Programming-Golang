package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	s_map := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
		"d": true,
		"e": true,
		"f": true,
		"h": true,
		"i": true,
		"j": true,
		"k": true,
		"m": true}

	var counter = struct {
		sync.RWMutex
		m map[string]bool
	}{m: s_map}

	keys := make([]string, 0, len(s_map))
	for k := range s_map {
		keys = append(keys, k)
	}

	for _, key := range keys {
		go func() {
			counter.RLock()
			counter.m[key] = false
			fmt.Println(key, counter.m[key])
			counter.RUnlock()
		}()
	}
	time.Sleep(10000)
}
