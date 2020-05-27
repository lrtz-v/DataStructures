package subarraysDivByK

import "testing"

func TestSubarraysDivByK(t *testing.T) {

	n := subarraysDivByK([]int{4,5,0,-2,-3,1}, 5)
	if n != 7 {
		t.Fatalf("[*] SubarraysDivByK Error! Expected: %d, Got %d\n", 7, n)
	}
}

func TestSubarraysDivByK2(t *testing.T) {

	n := subarraysDivByK([]int{-1, 2, 9}, 2)
	if n != 2 {
		t.Fatalf("[*] SubarraysDivByK Error! Expected: %d, Got %d\n", 2, n)
	}
}
