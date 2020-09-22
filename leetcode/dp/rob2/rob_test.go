package rob2

import "testing"

func TestRob2(t *testing.T) {
	s := rob2([]int{1, 7, 9, 2})
	if s != 11 {
		t.Fatalf("TestRob Failed. Got: %d\n", s)
	}
}

func TestRob3(t *testing.T) {
	s := rob2([]int{2, 3, 2})
	if s != 3 {
		t.Fatalf("TestRob Failed. Got: %d\n", s)
	}
}
