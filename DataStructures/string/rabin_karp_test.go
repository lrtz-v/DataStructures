package string

import "testing"

func Test_rabinKarp(t *testing.T) {

	index := rabinKarp("1234567890", "789")
	if index != 6 {
		t.Fatalf("[*] Expect %d, get %d", 6, index)
	}

	index = rabinKarp("1234567890", "a89")
	if index != -1 {
		t.Fatalf("[*] Expect %d, get %d", -1, index)
	}

}
