package minstack

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	minStack := Constructor()

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	min := minStack.GetMin();
	if min != -3 {
		t.Fatalf("[*] GetMin Error! Expected -3 Got %d\n", min)
	}
	minStack.Pop();
	top := minStack.Top();
	if top != 0 {
		t.Fatalf("[*] Top Error! Expected 0 Got %d\n", min)
	}
	min = minStack.GetMin();
	if min != -2 {
		t.Fatalf("[*] GetMin Error! Expected -2 Got %d\n", min)
	}
}

func TestMinStack2(t *testing.T) {
	minStack := Constructor()

	minStack.Push(2147483646)
	minStack.Push(2147483646)
	minStack.Push(2147483647)
	
	minStack.Top()
	minStack.Pop()
	minStack.GetMin()
	minStack.Pop()
	minStack.GetMin()
	minStack.Pop()
	minStack.Push(2147483647)
	minStack.Top()
	minStack.GetMin()
	minStack.Push(-2147483648)
	minStack.Top()
	minStack.GetMin()
	minStack.Pop()
	minStack.GetMin();
	
}
