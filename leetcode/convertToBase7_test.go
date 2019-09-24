package leetcode

import (
	"testing"
)

func TestConvertBaseN(t *testing.T) {
	s := convertToBase7(-7)
	if s != "-10" {
		t.Fatal("-7 err")
	}
}

func TestConvertBase100(t *testing.T) {
	s := convertToBase7(100)
	if s != "202" {
		t.Fatal("202 err")
	}
}
