package minCostClimbingStairs

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	var m int

	m = minCostClimbingStairs([]int{10, 15, 20})
	if m != 15 {
		t.Fatal("[*] TestMinCostClimbingStairs Error!")
	}
}
