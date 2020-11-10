package letterCasePermutation

import (
	"testing"
)

func TestLetterCasePermutationI(t *testing.T) {
	ans := letterCasePermutation("a1b2")
	t.Fatal(ans)
}

func TestStringChange(t *testing.T) {
	str := "abcd"
	str2 := letterChange(str, 1)
	if str2 != "aBcd" {
		t.Fatalf("Change str failed. %s ==> %s\n", str, str2)
	}
}
