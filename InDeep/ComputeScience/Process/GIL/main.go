package main

import (
	"fmt"
	"time"
)

// var wg sync.WaitGroup

// Count to 100000000
func Count() bool {
	var x int
	for i := 0; i < 100000000; i++ {
		x = x + 1
	}
	// wg.Done()
	return true
}

func process() {
	start := time.Now().UnixNano()
	Count()
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}

func main() {
	process()
}
