package convertToBase7

import (
	"testing"
)

func TestConvertBaseN(t *testing.T) {
	t.Skip()
	s := convertToBase7(-7)
	if s != "-10" {
		t.Fatal("-7 err")
	}
}

func TestConvertBase100(t *testing.T) {
	t.Skip()
	s := convertToBase7(100)
	if s != "202" {
		t.Fatal("202 err")
	}
}
