package messaging

import "time"

func Split(ch <-chan int, n int) []chan int {

	cs := make([]chan int, n)
	for i := 0; i < n; i++ {
		cs = append(cs, make(chan int))
	}

	close := func([]chan int) {
		for _, c := range cs {
			close(c)
		}
	}

	go func(ch <-chan int, cs []chan int) {
		defer close(cs)

		for {
			for _, c := range cs {
				select {
				case val := <-ch:
					c <- val
				case <-time.After(1 * time.Second):
					return
				}
			}
		}
	}(ch, cs)

	return cs
}
