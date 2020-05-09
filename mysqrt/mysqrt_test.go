package mysqrt

import "testing"

func TestMidSearch(t *testing.T) {
	i := midSearch(8)
	if i != 2 {
		t.Fatalf("[*] Error! Expetced 2, Got %d\n", i)
	}
}

func TestNewtonRaphson(t *testing.T) {
	i := int(newtonRaphson(8, 1e-7))
	if i != 2 {
		t.Fatalf("[*] Error! Expetced 2, Got %d\n", i)
	}
}
