package binaryheap

import (
	"testing"
)

func TestCreate(t *testing.T) {
	testData := []int{57, 40, 38, 11, 13, 34, 48, 75, 6, 19, 9, 7}
	bh := CreateBinaryHeap(testData)
	t.Log(bh.toString())
}

func TestInsert(t *testing.T) {
	testData := []int{57, 40, 38, 11, 13, 34, 48, 75, 6, 19, 9, 7}
	bh := CreateBinaryHeap(testData)
	bh.insert(100)
	if bh.heap[1] != 100 {
		t.Fatal("insert 1 failed!")
	}
	bh.insert(0)
	if bh.heap[bh.size()-1] != 0 {
		t.Fatal("insert 0 failed!")
	}
}

func TestDelete(t *testing.T) {
	testData := []int{57, 40, 38, 11, 13, 34, 48, 75, 6, 19, 9, 7}
	bh := CreateBinaryHeap(testData)  // 75,57,48,40,19,34,38,11,6,13,9,7
	top := bh.pop()
	if bh.heap[1] != 57 || top != 75 {
		t.Fatal("pop 0 failed!")
	}
}