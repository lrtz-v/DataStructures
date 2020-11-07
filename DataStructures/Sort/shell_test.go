package sort

import "testing"


func TestShellSort1(t *testing.T) {
	testData := []int64{3, 6, 7, 8, 2, 10, 9, 1, 2, 3, 6, 8, 2, 3, 5, 0}
	s := Shell{}
	s.Sort(testData)
	s.Show(testData)
	if !s.IsSorted(testData) {
		t.Errorf("Not Sorted")
	}
}

func TestShellSort2(t *testing.T) {
	testData := []int64{3, 3, 3, 3, 3, 4}
	s := Shell{}
	s.Sort(testData)
	s.Show(testData)
	if !s.IsSorted(testData) {
		t.Errorf("Not Sorted")
	}
}
