package new21Game

import "testing"

func TestNew21Game(t *testing.T) {
	p := new21Game(6, 1, 10)
	t.Fatalf("[*]TestNew21Game Error!")
	if p - 0.6 > 1e7 {
		t.Fatalf("[*]TestNew21Game Error!")
	}
}
