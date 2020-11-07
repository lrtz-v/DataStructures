package sort

import (
	"testing"
)

func TestMerge(t *testing.T) {
	data := []int64{5, 6, 7, 1, 2, 3, 4}
	m := Merge{}
	m.aux = make([]int64, len(data))
	m.merge(data, 0, 2, len(data)-1)
}

func TestSortD1(t *testing.T) {
	testData := []int64{3, 6, 7, 8, 2, 10, 9}
	s := Merge{}
	s.SortV2(testData)
	s.Show(testData)
	if !s.IsSorted(testData) {
		t.Errorf("Not Sorted")
	}
}

func TestSortD2(t *testing.T) {
	testData := []int64{3, 3, 3, 3, 3, 4}
	s := Merge{}
	s.Sort(testData)
	s.Show(testData)
	if !s.IsSorted(testData) {
		t.Errorf("Not Sorted")
	}
}

func TestSortU1(t *testing.T) {
	testData := []int64{3, 6, 7, 8, 2, 10, 9}
	s := Merge{}
	s.Sort(testData)
	s.Show(testData)
	if !s.IsSorted(testData) {
		t.Errorf("Not Sorted")
	}
}

func TestSortU2(t *testing.T) {
	testData := []int64{3, 3, 3, 3, 3, 4}
	s := Merge{}
	s.Sort(testData)
	s.Show(testData)
	if !s.IsSorted(testData) {
		t.Errorf("Not Sorted")
	}
}
