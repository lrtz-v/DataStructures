package leetcode

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	num1 := []int{1, 2}
	num2 := []int{3, 4}
	res := findMedianSortedArrays(num1, num2)
	fmt.Println(res)
}
