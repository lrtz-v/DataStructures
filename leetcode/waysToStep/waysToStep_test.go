package waysToStep

import "testing"

func TestWaysToStep(t *testing.T) {
	n := waysToStep(61)
	if n != 752119970 {
		t.Fatalf("[*] TestWaysToStep Error! Expected 752119970 Got %d\n", n)
	}
}
