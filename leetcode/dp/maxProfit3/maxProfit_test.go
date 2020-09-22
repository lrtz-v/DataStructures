package maxProfit

import "testing"

func TestMaxProfit(t *testing.T) {
	p2 := maxProfit([]int{3,3,5,0,0,3,1,4})
	if p2 != 6 {
		t.Fatalf("TestMaxProfit Failed. Got: %d\n", p2)
	}
}
