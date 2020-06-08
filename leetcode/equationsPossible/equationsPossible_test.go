package equationsPossible

import "testing"

func TestEquationsPossible(t *testing.T) {
	if equationsPossible([]string{"a==b", "b!=a"}) {
		t.Fatalf("[*] TestEquationsPossible Error!")
	}
}

func TestEquationsPossible2(t *testing.T) {
	if !equationsPossible([]string{"b==a", "a==b"}) {
		t.Fatalf("[*] TestEquationsPossible2 Error!")
	}
}

func TestEquationsPossible3(t *testing.T) {
	if equationsPossible([]string{"a==b","e==c","b==c","a!=e"}) {
		t.Fatalf("[*] TestEquationsPossible3 Error!")
	}
}

func TestEquationsPossible4(t *testing.T) {
	if equationsPossible([]string{"a==b","b!=c","c==a"}) {
		t.Fatalf("[*] TestEquationsPossible4 Error!")
	}
}

func TestEquationsPossible5(t *testing.T) {
	if !equationsPossible([]string{"c==c", "b==d", "x!=z"}) {
		t.Fatalf("[*] TestEquationsPossible5 Error!")
	}
}

func TestEquationsPossible6(t *testing.T) {
	if !equationsPossible([]string{"a!=b","b!=c","c!=a"}) {
		t.Fatalf("[*] TestEquationsPossible6 Error!")
	}
}
