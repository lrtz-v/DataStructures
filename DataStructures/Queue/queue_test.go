package Queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	// t.Skip()
	// testData := []int{1}
	testData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// testData := []int{1, 2, 3, 4, 5, 6, 7}
	// testData := []int{4, 2, 7, 1, 3, 6, 9}
	queue := NewQueue()
	for _, data := range testData {
		queue.Push(data)
	}

	if size := queue.size(); size != len(testData) {
		t.Errorf("push err")
	}

	for _, data := range testData {
		val, err := queue.Pop()
		fmt.Print(val)
		if err != nil {
			t.Error("pop err")
		}
		if val != data {
			t.Errorf("Expected %d, Got %d", data, val)
		}
	}
}
