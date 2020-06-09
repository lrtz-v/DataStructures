package translateNum

import "testing"

func TestTranslateNum(t *testing.T) {
	n := translateNum(12258)
	if n != 5 {
		t.Fatalf("[*] TestTranslateNum Error! Expected 5, Got: %d", n)
	}
	
	n = translateNum(506)
	if n != 1 {
		t.Fatalf("[*] TestTranslateNum Error! Expected 1, Got: %d", n)
	}
}