package minWindow

import "testing"

func TestMinWindow(t *testing.T) {
	s := minWindow("ADOBECODEBANCC", "ABC")
	if s != "BANC" {
		t.Fatal("[*] Find MinWindow Error!")
	}
}
