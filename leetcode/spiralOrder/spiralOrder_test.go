package spiralOrder

import "testing"

func TestSpiralOrder(t *testing.T) {
	matrix := make([][]int, 3)
	var s []int
	matrix[0] = []int{1,2,3}
	matrix[1] = []int{4,5,6}
	matrix[2] = []int{7,8,9}

	s = spiralOrder(matrix)
	for i, v := range []int{1,2,3,6,9,8,7,4,5} {
		if s[i] != v {
			t.Fatal(s)
		}
	}

	matrix[0] = []int{1,2,3,4}
	matrix[1] = []int{5,6,7,8}
	matrix[2] = []int{9,10,11,12}
	s = spiralOrder(matrix)
	if len(s) != len([]int{1,2,3,4,8,12,11,10,9,5,6,7}) {
		t.Fatal(s)
	}
	for i, v := range []int{1,2,3,4,8,12,11,10,9,5,6,7} {
		if s[i] != v {
			t.Fatal(s)
		}
	}
}
