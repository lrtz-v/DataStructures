package minArray

import "testing"

func TestMinArray(t *testing.T) {
	n := minArray([]int{3, 4, 5, 1, 2})
	if n != 1 {
		t.Fatalf("[*] TestMinArray Error! Expected 1 Got %d\n", n)
	}

	n = minArray([]int{2, 2, 2, 0, 1})
	if n != 0 {
		t.Fatalf("[*] TestMinArray Error! Expected 0 Got %d\n", n)
	}
}