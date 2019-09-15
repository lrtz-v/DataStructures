package structuraldesignpattern

import (
	"testing"
)

func Double(n int) int {
	return n * 2
}

func TestLogDecorate(t *testing.T) {
	f := LogDecorate(Double)
	if f(5) != 10 {
		t.Fatal("Double Error")
	}
}
