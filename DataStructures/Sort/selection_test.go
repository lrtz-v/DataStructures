package sort

import "testing"

func TestSelection1(t *testing.T) {
	testData := []int64{3, 6, 7, 8, 2, 10, 9}
	s := Selection{}
	s.Sort(testData)

	s.Show(testData)
	if !s.IsSorted(testData) {
		t.Errorf("Not Sorted")
	}
}

func TestSelection2(t *testing.T) {
	testData := []int64{3, 3, 3, 3, 3, 4}
	s := Selection{}
	s.Sort(testData)
	s.Show(testData)
	if !s.IsSorted(testData) {
		t.Errorf("Not Sorted")
	}
}
