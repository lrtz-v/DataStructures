package climbStairs

import "testing"


func TestClimbStairs(t *testing.T) {
	n := climbStairs(2)
	if n != 2 {
		t.Fatal("[*] TestClimbStairs Error! Expected 2")
	}
	n = climbStairs(3)
	if n != 3 {
		t.Fatal("[*] TestClimbStairs Error! Expected 3")
	}
	n = climbStairs(44)
	if n != 1134903170 {
		t.Fatal("[*] TestClimbStairs Error! Expected 1134903170")
	}
}
