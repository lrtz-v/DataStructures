package string

import "testing"

func Test_string_find(t *testing.T) {

	index := bruteForce("1234567890", "789")
	if index != 6 {
		t.Fatalf("[*] Expect %d, get %d", 6, index)
	}

	index = bruteForce("1234567890", "a89")
	if index != -1 {
		t.Fatalf("[*] Expect %d, get %d", -1, index)
	}
}
