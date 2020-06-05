package minPathSum

import (
	"testing"
)

func TestMinPathSum(t *testing.T) {
	arr := make([][]int, 3)
	arr[0] = []int{1, 3, 1}
	arr[1] = []int{1, 5, 1}
	arr[2] = []int{4, 2, 1}
	res := minPathSum(arr)
	if res != 7 {
		t.Fatal("Error")
	}
}
