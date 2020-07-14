package findNumberIn2DArray

import "testing"

func TestFindNumberIn2DArray(t *testing.T) {
	matrix := make([][]int, 5)
	matrix[0] = []int{1, 4, 7, 11, 15}
	matrix[1] = []int{2, 5, 8, 12, 19}
	matrix[2] = []int{3, 6, 9, 16, 22}
	matrix[3] = []int{10, 13, 14, 17, 24}
	matrix[4] = []int{18, 21, 23, 26, 30}

	if !findNumberIn2DArray(matrix, 5) {
		t.Fatalf("[*] findNumberIn2DArray Error! Expected: true")
	}

	if findNumberIn2DArray(matrix, 20) {
		t.Fatalf("[*] findNumberIn2DArray Error! Expected: false")
	}

	if !findNumberIn2DArray2(matrix, 5) {
		t.Fatalf("[*] findNumberIn2DArray2 Error! Expected: true")
	}

	if findNumberIn2DArray2(matrix, 20) {
		t.Fatalf("[*] findNumberIn2DArray2 Error! Expected: false")
	}
}
