package longestCommonPrefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	str := longestCommonPrefix([]string{"flower","flow","flight"})
	if str != "fl" {
		t.Fatal("[*] TestLongestCommonPrefix Error! Expected 'fl'.")
	}
	str = longestCommonPrefix([]string{"dog","racecar","car"})
	if str != "" {
		t.Fatal("[*] TestLongestCommonPrefix Error! Expected ''.")
	}
	str = longestCommonPrefix([]string{"aca","cba"})
	if str != "" {
		t.Fatal("[*] TestLongestCommonPrefix Error! Expected ''.")
	}
}
