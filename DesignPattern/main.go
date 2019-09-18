package main

import (
	"fmt"
	"time"
)

func Split(ch chan int, n int) {

	cs := make([]chan int, n)
	for i := 0; i < n; i++ {
		cs = append(cs, make(chan int))
	}

	close := func([]chan int) {
		for _, c := range cs {
			close(c)
		}
	}

	go func() {
		defer close(cs)

		for {
			select {
			case val := <-ch:
				fmt.Printf("[*]Get: %d\t", val)
			case <-time.After(2 * time.Second):
				return
			}
		}
	}()

}

func main() {
	ch := make(chan int, 20)

	go func() {
		defer close(ch)
		for i := 0; i < 20; i++ {
			ch <- i
			fmt.Printf("[*]Put %d\t", i)
		}
	}()

	// var wg sync.WaitGroup
	// listen := func(s chan int) {
	// 	for {
	// 		val, ok := <-s
	// 		if !ok {
	// 			break
	// 		}
	// 		fmt.Printf("[*]Get %d\t", val)
	// 	}
	// 	wg.Done()
	// }

	Split(ch, 3)

	// for _, s := range cs {
	// 	go listen(s)
	// 	wg.Add(1)
	// }

	// go func() {
	// 	wg.Wait()
	// }()
	time.Sleep(10 * time.Second)
}
