package longestPalindrome

import "testing"

func TestLongestPalindrome(t *testing.T) {
	s := longestPalindrome("babad")
	if s != "bab" && s != "aba" {
		t.Fatalf("[*] Find LongestPalindrome Error! Expected: bab/aba, Got %s\n", s)
	}
}

func TestLongestPalindrome2(t *testing.T) {
	s := longestPalindrome("cbbd")
	if s != "bb" {
		t.Fatalf("[*] Find LongestPalindrome Error! Expected: bb, Got %s\n", s)
	}
}
