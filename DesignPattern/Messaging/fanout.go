package messaging

import "fmt"

// Split elements into n chans
func Split(ch <-chan int, n int) map[string]chan int {

	cs := make(map[string]chan int)
	for i := 0; i < n; i++ {
		key := fmt.Sprintf("Chan Num_%d\n", i)
		cs[key] = make(chan int)
	}

	close := func(cs map[string]chan int) {
		for key, c := range cs {
			close(c)
			fmt.Printf("[*] CLose %s\n", key)
		}
	}

	go func(ch <-chan int, cs map[string]chan int) {
		defer close(cs)

		for {
			for _, c := range cs {
				select {
				case val, ok := <-ch:
					if !ok {
						return
					}
					c <- val
					fmt.Printf("[*] Translate %d\n", val)
				}
			}
		}
	}(ch, cs)

	return cs
}
