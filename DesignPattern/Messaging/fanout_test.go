package messaging

import (
	"fmt"
	"sync"
	"testing"
)

func send() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func TestSplit(t *testing.T) {
	send := send()

	var wg sync.WaitGroup
	listen := func(name string, s <-chan int) {
		fmt.Printf("[*] Listen: %s\n", name)
		for {
			val, ok := <-s
			if !ok {
				break
			}
			fmt.Printf("[*]Get %d\n", val)
		}
		wg.Done()
	}

	cs := Split(send, 3)

	for name, s := range cs {
		go listen(name, s)
		wg.Add(1)
	}

	wg.Wait()
}
