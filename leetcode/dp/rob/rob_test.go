package rob

import "testing"

func TestRob(t *testing.T) {
	nums := []int{1,2,3,1}
	s := rob(nums) 
	if s != 4 {
		t.Fatalf("[*] TestRob Failed. Got %d.\n", s)
	}
}
