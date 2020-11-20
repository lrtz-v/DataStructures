package sort

import (
	"testing"
)

func TestQuickSort1(t *testing.T) {
	testData := []int64{3, 6, 7, 8, 2, 10, 9}
	q := Quick{}
	q.Sort(testData)
	// q.Show(testData)
	if !q.IsSorted(testData) {
		t.Errorf("Not Sorted")
	}
}

func TestQuickSort2(t *testing.T) {
	testData := []int64{3, 3, 3, 3, 3, 4}
	q := Quick{}
	q.Sort(testData)
	// q.Show(testData)
	if !q.IsSorted(testData) {
		t.Errorf("Not Sorted")
	}
}
