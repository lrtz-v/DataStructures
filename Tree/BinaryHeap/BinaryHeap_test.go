package BinaryHeap

import (
	"testing"
)

func TestCreateMaxHeap(t *testing.T) {
	// t.Skip()
	testData := []int{3, 6, 7, 8, 2, 10, 9}
	testHeap, err := NewMaxHeap(testData)
	if err != nil {
		t.Errorf("Build error")
	}
	if testHeap.Heap[0].Value != 10 {
		t.Errorf("Max heap built failed")
	}
	testHeap.Clear()
}

func TestInsert(t *testing.T) {
	// t.Skip()
	testData := []int{3, 6, 7, 8, 2, 10, 9}
	testHeap, err := NewMaxHeap(testData)
	if err != nil {
		t.Errorf("Build error")
	}
	testHeap.Insert(11)
	if testHeap.Heap[0].Value != 11 {
		t.Errorf("Max heap insert failed")
	}
	testHeap.Clear()
}

func TestExtractMax(t *testing.T) {
	// t.Skip()
	testData := []int{3, 6, 7, 8, 2, 10, 9}
	testHeap, err := NewMaxHeap(testData)
	if err != nil {
		t.Errorf("Build error")
	}
	testHeap.Insert(11)
	max := testHeap.extractMax()
	if max != 11 {
		t.Errorf("Extract failed")
	}
	if len(testHeap.Heap) != len(testData) {
		t.Errorf("Max heap length error")
	}
	if testHeap.Heap[0].Value != 10 {
		t.Errorf("Max heap remove failed")
	}

	testHeap.Clear()
}

func TestRemove(t *testing.T) {
	// t.Skip()
	testData := []int{3, 6, 7, 8, 2, 10, 9}
	testHeap, err := NewMaxHeap(testData)
	if err != nil {
		t.Errorf("Build error")
	}
	testHeap.Insert(11)
	testHeap.remove(2)
	testHeap.Clear()
}
