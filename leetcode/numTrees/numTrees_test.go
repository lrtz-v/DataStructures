package numTrees

import "testing"

func TestNumTrees(t *testing.T) {
	s := numTrees(3)
	if s != 5 {
		t.Fatal("[*] TestNumTrees Error!")
	}
}
