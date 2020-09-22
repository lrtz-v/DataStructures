package maxProfit

import "testing"

func TestMaxProfit(t *testing.T) {
	p2 := maxProfit([]int{7,1,5,3,6,4})
	if p2 != 5 {
		t.Fatalf("TestMaxProfit Failed. Got: %d\n", p2)
	}
}
