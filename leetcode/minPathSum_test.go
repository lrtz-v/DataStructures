package leetcode

import (
	"testing"
)

func TestMinPathSum(t *testing.T) {
	var arr [][]int
	arr = [[1,3,1], [1,5,1], [4,2,1]]
	res := minPathSum(arr)
	if arr != 7 {
		t.Fatal("Error")
	}
}
