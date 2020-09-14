package findLongestWord

import "testing"

func TestFindLongestWord(t *testing.T) {
	s := "abpcplea"
	d := []string{"ale","apple","monkey","plea"}

	result := findLongestWord(s, d)
	if result != "apple" {
		t.Fatalf("[*]TestFindLongestWord Error, Expected apple, Got: [%s]", result)
	}
}

func TestFindLongestWord2(t *testing.T) {
	s := "bab"
	d := []string{"ba","ab","a","b"}

	result := findLongestWord(s, d)
	if result != "ab" {

		t.Fatalf("[*]TestFindLongestWord Error, Expected ab, Got: [%s]", result)
	}
}
