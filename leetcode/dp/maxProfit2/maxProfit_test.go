package maxProfit

import "testing"

func TestMaxProfit(t *testing.T) {
	p2 := maxProfit2([]int{1, 3, 2, 8, 4, 9}, 2)
	if p2 != 8 {
		t.Fatalf("TestMaxProfit Failed. Got: %d\n", p2)
	}
}
