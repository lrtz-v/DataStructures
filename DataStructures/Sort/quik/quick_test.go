package quik

import "testing"

func TestSort1(t *testing.T) {
	testData := []int{3, 6, 7, 8, 2, 10, 9}
	s := &quikSort{comparableList: testData}
	s.sort(0, len(testData)-1)
	s.show()
	if !s.isSorted() {
		t.Errorf("Not Sorted")
	}
}

func TestSort2(t *testing.T) {
	testData := []int{3, 3, 3, 3, 3, 4}
	s := &quikSort{comparableList: testData}
	s.sort(0, len(testData)-1)
	s.show()
	if !s.isSorted() {
		t.Errorf("Not Sorted")
	}
}
