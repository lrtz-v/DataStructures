package Stack

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	testArr := []int{1, 3, 9, 2, 0}
	minStack := NewMinStack()
	for _, val := range testArr {
		minStack.Push(val)
	}
	min := minStack.GetMin()
	if min != 0 {
		t.Fatalf("Expected Min: %d, Got: %d", 0, min)
	}

	popData := minStack.Pop()
	if popData != 0 {
		t.Fatalf("Expected Pop Data: %d, Got: %d", 0, popData)
	}

	min = minStack.GetMin()
	if min != 1 {
		t.Fatalf("Expected Min: %d, Got: %d", 1, min)
	}

	popData = minStack.Pop()
	if popData != 2 {
		t.Fatalf("Expected Pop Data: %d, Got: %d", 2, popData)
	}
}
