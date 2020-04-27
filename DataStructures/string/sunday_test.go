package string

import "testing"

func Test_sunday(t *testing.T) {
	index := 0
	index = sunday("1234567890", "789")
	if index != 6 {
		t.Fatalf("[*] Expect %d, get %d", 6, index)
	}

	index = sunday("1234567890", "a89")
	if index != -1 {
		t.Fatalf("[*] Expect %d, get %d", -1, index)
	}

	index = sunday("aaabaaabaaabaaab", "aaaa")
	if index != -1 {
		t.Fatalf("[*] Expect %d, get %d", -1, index)
	}

	index = sunday("aaaaaaaaaaaaaaaa", "baaa")
	if index != -1 {
		t.Fatalf("[*] Expect %d, get %d", -1, index)
	}

}

