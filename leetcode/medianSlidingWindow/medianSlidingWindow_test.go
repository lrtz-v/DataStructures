package medianSlidingWindow

import (
	"fmt"
	"testing"
)

func TestMedianSlidingWindow(t *testing.T) {
	t.Skip()
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	s := medianSlidingWindow(nums, 3)
	fmt.Println(s)
}
