package leetcode

import (
	"testing"
)

func TestBinaryGap(t *testing.T) {
	s := binaryGap(22)
	if s != 2 {
		t.Fatalf("Expect 2, Got, %d", s)
	}
}

func TestBinaryGap2(t *testing.T) {
	s := binaryGap(5)
	if s != 2 {
		t.Fatalf("Expect 2, Got, %d", s)
	}
}
