package concurrency

func Count(start, end, step int) chan int {
	ch := make(chan int)

	go func(ch chan int) {
		for i := start; i < end; i += step {
			// Blocks on the operation
			ch <- i
		}

		close(ch)
	}(ch)

	return ch
}
