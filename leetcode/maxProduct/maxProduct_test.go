package maxProduct

import "testing"

func TestMaxProduct(t *testing.T) {

	m := maxProduct([]int{2, 3, -2 ,4, -1})
	if m != 48 {
		t.Fatalf("[*] MaxProduct Error! Expected %d Got %d\n", 48, m)
	}
}
