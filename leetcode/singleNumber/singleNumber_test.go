package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	s := singleNumber([]int{2,2,1})
	if s != 1 {
		t.Fatalf("[*] Find singleNumber error! Expected 1 Got %d\n", s)
	}
}