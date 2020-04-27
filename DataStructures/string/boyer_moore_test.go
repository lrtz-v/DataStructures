package string

import "testing"

func Test_boyer_moore(t *testing.T) {
	index := 0
	index = boyer_moore("1234567890", "789")
	if index != 6 {
		t.Fatalf("[*] Expect %d, get %d", 6, index)
	}

	index = boyer_moore("1234567890", "a89")
	if index != -1 {
		t.Fatalf("[*] Expect %d, get %d", -1, index)
	}

	index = boyer_moore("aaabaaabaaabaaab", "aaaa")
	if index != -1 {
		t.Fatalf("[*] Expect %d, get %d", -1, index)
	}

	index = boyer_moore("aaaaaaaaaaaaaaaa", "baaa")
	if index != -1 {
		t.Fatalf("[*] Expect %d, get %d", -1, index)
	}

}
