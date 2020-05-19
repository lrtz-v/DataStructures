package validPalindrome

import (
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	if validPalindrome("abc") {
		t.Fatal("[*] abc ValidPalindrome Error!")
	}
}

func TestValidPalindrome2(t *testing.T) {
	if !validPalindrome("abca") {
		t.Fatal("[*] abca ValidPalindrome Error!")
	}
}

func TestValidPalindrome4(t *testing.T) {
	if !validPalindrome("deee") {
		t.Fatal("[*] deee ValidPalindrome Error!")
	}
}


func TestValidPalindrome5(t *testing.T) {
	if !validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga") {
		t.Fatal("[*] aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga ValidPalindrome Error!")
	}
}
