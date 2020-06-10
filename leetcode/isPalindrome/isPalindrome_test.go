package isPalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !isPalindrome(121) {
		t.Fatal("[*] TestIsPalindrome Error! 121 is palindrome")
	}

	if !isPalindrome(14641) {
		t.Fatal("[*] TestIsPalindrome Error! 122222221 is palindrome")
	}

	if !isPalindrome(122222221) {
		t.Fatal("[*] TestIsPalindrome Error! 122222221 is palindrome")
	}

	if isPalindrome(-122222221) {
		t.Fatal("[*] TestIsPalindrome Error! -122222221 is palindrome")
	}

	if isPalindrome(-121) {
		t.Fatal("[*] TestIsPalindrome Error! -121 is not palindrome")
	}
}
