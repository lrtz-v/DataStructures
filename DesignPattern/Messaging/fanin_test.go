package messaging

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	t.Skip()

	in1 := make(chan int)
	in2 := make(chan int)

	go func() {
		defer close(in1)
		defer close(in2)
		for i := 1; i < 200; i++ {
			switch {
			case i%2 == 0:
				in1 <- i
			case i%3 == 0:
				in2 <- i
			default:
			}
		}
	}()

	out := Merge(in1, in2)
	for i := range out {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()

}
