package fib

import "testing"

func TestFib(t *testing.T) {
	n := fib(2)
	if n != 1 {
		t.Fatalf("[*] TestFib Error! Expected 1, Got %d\n", n)
	}

	n = fib(5)
	if n != 5 {
		t.Fatalf("[*] TestFib Error! Expected 5, Got %d\n", n)
	}
}
