package BinaryHeap

import (
	"testing"
)

func TestCreateMaxHeap(t *testing.T) {
	// t.Skip()
	testData := []int{3, 6, 7, 8, 2, 10, 9}
	b := NewBinaryHeap(testData)
	b.Print()
}

func TestPop(t *testing.T) {
	// t.Skip()
	testData := []int{3, 6, 7, 8, 2, 10, 9}
	b := NewBinaryHeap(testData)
	b.Print()
	b.Pop()
	b.Print()
}

func TestInsert(t *testing.T) {
	// t.Skip()
	testData := []int{3, 6, 7, 8, 2, 10, 9}
	b := NewBinaryHeap(testData)
	// b.print()
	b.Insert(100)
	b.Print()
}
