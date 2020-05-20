package findTheLongestSubstring

import "testing"

func TestFindTheLongestSubstring(t *testing.T) {
	l := findTheLongestSubstring("eleetminicoworoep")
	if l != 13 {
		t.Fatalf("[*] findTheLongestSubstring Error! Expected %d Got %d\n", 13, l)
	}
}

func TestFindTheLongestSubstring2(t *testing.T) {
	l := findTheLongestSubstring("leetcodeisgreat")
	if l != 5 {
		t.Fatalf("[*] findTheLongestSubstring Error! Expected %d Got %d\n", 5, l)
	}
}

func TestFindTheLongestSubstring3(t *testing.T) {
	l := findTheLongestSubstring("bcbcbc")
	if l != 6 {
		t.Fatalf("[*] findTheLongestSubstring Error! Expected %d Got %d\n", 6, l)
	}
}