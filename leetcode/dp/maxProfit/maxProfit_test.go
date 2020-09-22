package maxProfit

import "testing"

func TestMaxProfit(t *testing.T) {
	p := maxProfit([]int{1, 2, 3, 0, 2})
	if p != 3 {
		t.Fatalf("TestMaxProfit Failed. Gor: %d\n", p)
	}
}

func TestMaxProfit2(t *testing.T) {
	p := maxProfit2([]int{1, 2, 3, 0, 2})
	if p != 3 {
		t.Fatalf("TestMaxProfit Failed. Gor: %d\n", p)
	}
}
