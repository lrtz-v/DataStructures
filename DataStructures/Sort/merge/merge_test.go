package merge

import "testing"

func TestSortD1(t *testing.T) {
	testData := []int{3, 6, 7, 8, 2, 10, 9}
	s := &mergeSort{comparableList: testData}
	s.sortD(0, len(testData)-1)
	s.show()
	if !s.isSorted() {
		t.Errorf("Not Sorted")
	}
}

func TestSortD2(t *testing.T) {
	testData := []int{3, 3, 3, 3, 3, 4}
	s := &mergeSort{comparableList: testData}
	s.sortD(0, len(testData)-1)
	s.show()
	if !s.isSorted() {
		t.Errorf("Not Sorted")
	}
}

func TestSortU1(t *testing.T) {
	testData := []int{3, 6, 7, 8, 2, 10, 9}
	s := &mergeSort{comparableList: testData}
	s.sortU()
	s.show()
	if !s.isSorted() {
		t.Errorf("Not Sorted")
	}
}

func TestSortU2(t *testing.T) {
	testData := []int{3, 3, 3, 3, 3, 4}
	s := &mergeSort{comparableList: testData}
	s.sortU()
	s.show()
	if !s.isSorted() {
		t.Errorf("Not Sorted")
	}
}
