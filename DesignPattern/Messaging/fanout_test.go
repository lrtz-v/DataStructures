package messaging

import (
	"fmt"
	"sync"
	"testing"
)

func TestSplit(t *testing.T) {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 20; i++ {
			ch <- i % 3
		}
	}()

	var wg sync.WaitGroup
	listen := func(s chan int) {
		for {
			val, ok := <-s
			if !ok {
				break
			}
			fmt.Printf("[*]Get %d\t", val)
		}
		wg.Done()
	}

	cs := Split(ch, 3)

	go func() {
		for _, s := range cs {
			go listen(s)
			wg.Add(1)
		}
		wg.Wait()
	}()
}
