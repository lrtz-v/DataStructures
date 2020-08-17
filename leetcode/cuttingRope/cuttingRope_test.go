package cuttingrope

import "testing"

func TestCuttingRope(t *testing.T) {
	max := cuttingRope(4)
	if max != 4 {
		t.Fatalf("[*] testCuttingRope Failed. Expected 4, Got %d\n", max)
	}
}

func TestCuttingRope2(t *testing.T) {
	max := cuttingRope(10)
	if max != 36 {
		t.Fatalf("[*] testCuttingRope Failed. Expected 36, Got %d\n", max)
	}
}

func TestCuttingRope3(t *testing.T) {
	max := cuttingRope(8)
	if max != 18 {
		t.Fatalf("[*] testCuttingRope Failed. Expected 18, Got %d\n", max)
	}
}

func TestCuttingRope4(t *testing.T) {
	max := cuttingRope(3)
	if max != 2 {
		t.Fatalf("[*] testCuttingRope Failed. Expected 2, Got %d\n", max)
	}
}