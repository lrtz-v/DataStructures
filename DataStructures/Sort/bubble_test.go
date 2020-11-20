package sort

import (
	"testing"
)

func TestBubble1(t *testing.T) {
	// t.Skip()
	testArray := []int64{5, 8, 6, 3, 9, 2, 1, 7}
	b := Bubble{}
	b.Sort(testArray)
	if !b.IsSorted(testArray) {
		t.Fatal(testArray)
		t.Fatal("Buffer sort failed")
	}
}

func TestBubble2(t *testing.T) {
	testArray := []int64{3, 4, 2, 1, 5, 6, 7, 8}
	b := Bubble{}
	b.Sort(testArray)
	if !b.IsSorted(testArray) {
		t.Fatal(testArray)
		t.Fatal("Buffer sort failed")
	}
}

func TestCockTailSort(t *testing.T) {
	testArray := []int64{5, 8, 6, 3, 9, 2, 1, 7}
	b := Bubble{}
	b.SortV2(testArray)
	if !b.IsSorted(testArray) {
		t.Fatal(testArray)
		t.Fatal("cockTail Bubble sort failed")
	}
}
