package binaryheap

import "testing"

func TestBinaryHeapSort(t *testing.T) {
	testData := []int{57, 40, 38, 11, 13, 34, 48, 75, 6, 19, 9, 7}
	bh := CreateBinaryHeap(testData)

	bh.sort()

	top := bh.pop()
	if top != 6 {
		t.Fatal("pop top failed!")
	}

}
