package ch1

import "testing"

func TestParenthesesValid(t *testing.T) {
	src := "[()]{}{[()()]()}"
	if !parenthesesValid(src) {
		t.Fatal("check failed on \"[()]{}{[()()]()}\"")
	}

	src = "[(])"
	if parenthesesValid(src) {
		t.Fatal("check failed on \"[(])\"")
	}

	src = "[()]"
	if !parenthesesValid(src) {
		t.Fatal("check failed on \"[()]\"")
	}

	src = ")[)]"
	if parenthesesValid(src) {
		t.Fatal("check failed on \")[)]\"")
	}

	src = "((((((()))))))"
	if !parenthesesValid(src) {
		t.Fatal("check failed on \"((((((()))))))\"")
	}
}
