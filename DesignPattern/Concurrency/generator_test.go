package concurrency

import (
	"fmt"
	"testing"
)

func TestGenerator(t *testing.T) {
	count := Count(1, 20, 2)
	for i := range count {
		fmt.Printf("i: %d\n", i)
	}
}
