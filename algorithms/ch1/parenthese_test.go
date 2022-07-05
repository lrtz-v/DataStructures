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

func TestCompleteLeftParentheses(t *testing.T) {
	src := "1 + 2 ) * 3 - 4 ) * 5 - 6 ) ) )"
	res := completeLeftParentheses(src)
	if res != "( ( 1 + 2 ) * ( ( 3 - 4 ) * ( 5 - 6 ) ) )" {
		t.Fatalf("complete failed on %s", src)
	}
}
