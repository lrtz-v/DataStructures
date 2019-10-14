package leetcode

import (
	"testing"
)

func TestGcd(t *testing.T) {
	t.Skip()
	res := gcd(16, 4)
	if res != 4 {
		t.Errorf("GCD Expected %d, Got %d", 4, res)
	}
}

func TestGcd2(t *testing.T) {
	t.Skip()
	res := gcd(16, 3)
	if res != 1 {
		t.Errorf("GCD Expected %d, Got %d", 1, res)
	}
}

func TestGcd3(t *testing.T) {
	t.Skip()
	res := gcd(3, 16)
	if res != 1 {
		t.Errorf("GCD Expected %d, Got %d", 1, res)
	}
}

func TestGcd4(t *testing.T) {
	t.Skip()
	res := gcd(9, 27)
	if res != 9 {
		t.Errorf("GCD Expected %d, Got %d", 9, res)
	}
}
