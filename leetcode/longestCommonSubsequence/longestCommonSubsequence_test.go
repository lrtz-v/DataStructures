package longestCommonSubsequence

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	l := longestCommonSubsequence("abcde", "ace")
	if l != 3 {
		t.Fatalf("[*] Get longestCommonSubsequence failed: abcde - ace, longest: %d\n", l)
	}

	l = longestCommonSubsequence("abc", "abc")
	if l != 3 {
		t.Fatalf("[*] Get longestCommonSubsequence failed: abc - abc, longest: %d\n", l)
	}

	l = longestCommonSubsequence("ezupkr", "ubmrapg")
	if l != 2 {
		t.Fatalf("[*] Get longestCommonSubsequence failed: ezupkr - ubmrapg, longest: %d\n", l)
	}

	l = longestCommonSubsequence("bsbininm", "jmjkbkjkv")
	if l != 1 {
		t.Fatalf("[*] Get longestCommonSubsequence failed: bsbininm - jmjkbkjkv, longest: %d\n", l)
	}
}