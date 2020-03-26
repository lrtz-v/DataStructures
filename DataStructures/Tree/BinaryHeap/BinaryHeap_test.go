package binaryheap

import (
	"testing"
)

func TestCreate(t *testing.T) {
	testData := []int{57, 40, 38, 11, 13, 34, 48, 75, 6, 19, 9, 7}
	aim := []int{6, 7, 9, 11, 13, 19, 34, 38, 40, 48, 57, 75}
	bh := CreateBinaryHeap(testData)
	if len(testData) != len(bh.heap) || len(aim) != len(bh.heap) {
		t.Fatal("elements missing!")
	}
	for i := 0; i < len(aim); i++ {
		if aim[i] != bh.heap[i] {
			t.Fatal("init failed!")
		}
	}
	t.Log(bh.toString())
}

func TestInsert(t *testing.T) {
	testData := []int{57, 40, 38, 11, 13, 34, 48, 75, 6, 19, 9, 7}
	bh := CreateBinaryHeap(testData)
	bh.insert(1)
	t.Log(bh.toString())
	if bh.heap[0] != 1 {
		t.Fatal("insert 1 failed!")
	}
	bh.insert(0)
	if bh.heap[0] != 0 {
		t.Fatal("insert 0 failed!")
	}
}
