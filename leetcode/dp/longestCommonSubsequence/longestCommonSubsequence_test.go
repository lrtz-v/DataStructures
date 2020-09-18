package longestCommonSubsequence

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	l := longestCommonSubsequence("abcde", "ace")
	if l != 3 {
		t.Fatalf("[*] TestLongestCommonSubsequence Error. Got %d.\n", l)
	}
}


func TestLongestCommonSubsequence2(t *testing.T) {
	l := longestCommonSubsequence2("abcde", "ace")
	if l != 3 {
		t.Fatalf("[*] TestLongestCommonSubsequence Error. Got %d.\n", l)
	}
}
