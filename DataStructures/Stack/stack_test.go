package Stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	// t.Skip()
	// testData := []int{1}
	testData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	testData2 := []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	// testData := []int{1, 2, 3, 4, 5, 6, 7}
	// testData := []int{4, 2, 7, 1, 3, 6, 9}
	Stack := NewStack()
	for _, data := range testData {
		Stack.Push(data)
	}

	if size := Stack.size(); size != len(testData) {
		t.Errorf("push err")
	}

	for _, data := range testData2 {
		val, err := Stack.Pop()
		if err != nil {
			t.Error("pop err")
		}
		if val != data {
			t.Errorf("Expected %d, Got %d", data, val)
		}
	}
	if !Stack.Empty() {
		t.Errorf("Something exist")
	}
}
