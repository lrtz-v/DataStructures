package heapSort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	testArray := []int{1, 3, 2, 6, 5, 7, 8, 9, 10, 0}
	heapSort(testArray)
	fmt.Println(testArray)
}
