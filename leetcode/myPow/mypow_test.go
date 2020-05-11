package mypow

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	p := myPow(2.00000, 10)
	if math.Abs(p - 1024.00000) > 1e-7 {
		t.Fatalf("[*] Pow Error! Expected 1024.00000 Got %f\n", p)
	}
	p = myPow(2.10000, 3)
	if math.Abs(p - 9.26100) > 1e-7 {
		t.Fatalf("[*] Pow Error! Expected 9.26100 Got %f\n", p)
	}
	p = myPow(2.00000, -2)
	if math.Abs(p - 0.25) > 1e-7 {
		t.Fatalf("[*] Pow Error! Expected 0.25 Got %f\n", p)
	}
	p = myPow(0.00001, 200)
	if math.Abs(p - 0.0000000000) > 1e-7 {
		t.Fatalf("[*] Pow Error! Expected 0.0000000000 Got %f\n", p)
	}
}
