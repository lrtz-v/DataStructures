package package01

import "testing"

func TestPackage01(t *testing.T) {
	n := package01(4, 5, []int{2, 4, 4, 5}, []int{1, 2, 3, 4})
	if n != 8 {
		t.Fatalf("[*] TestPackage01 Error! Gor %d\n", n)
	}

	m := zeroOnePack(4, 5, []int{2, 4, 4, 5}, []int{1, 2, 3, 4})
	if m != 8 {
		t.Fatalf("[*] TestPackage01 Error! Gor %d\n", m)
	}
}
