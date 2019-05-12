package leetcode

import (
	"fmt"
	"testing"
)

func TestGetMaxSortedDistance(t *testing.T) {
	t.Skip()
	testArr := []float64{4.12, 6.421, 0.0023, 3.0, 2.123, 8.122, 4.12, 10.09}
	maxDistance := getMaxSortedDistance(testArr)
	fmt.Println()
	fmt.Printf("maxDistance: %f", maxDistance)
	fmt.Println()
}

func TestGetMaxSortedDistance2(t *testing.T) {
	t.Skip()
	testArr := []float64{2, 6, 3, 4, 5, 10, 9}
	maxDistance := getMaxSortedDistance(testArr)
	fmt.Println()
	fmt.Printf("maxDistance: %f", maxDistance)
	fmt.Println()
}
