package maxSubArray

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	l := []int{-2,1,-3,4,-1,2,1,-5,4}
	max := maxSubArray(l)
	fmt.Println(max)
	if max != 6 {
		t.Fatal("Max Error!")
	}
}
