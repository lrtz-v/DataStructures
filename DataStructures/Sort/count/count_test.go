package count

import (
	"testing"
)

func TestCountSort(t *testing.T) {
	// t.Skip()
	testArr := []int{9, 3, 5, 4, 9, 1, 2, 7, 8, 1, 3, 6, 5, 3, 4, 0, 10, 9, 7, 9}
	countSort(testArr)
}

func TestCountSort2(t *testing.T) {
	// t,Skip()
	testArr := []int{95, 94, 91, 98, 99, 90, 99, 93, 91, 92}
	countSort2(testArr)
}
