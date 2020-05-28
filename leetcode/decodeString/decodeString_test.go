package decodeString

import (
	"testing"
)

func TestDecodeString(t *testing.T) {
	s := decodeString("3[a]2[bc]")
	if s != "aaabcbc" {
		t.Fatalf("[*] DecodeString Error, Input: %s, Got :%s\n", "aaabcbc", s)
	}
}

func TestDecodeString2(t *testing.T) {
	s := decodeString("3[a2[c]]")
	if s != "accaccacc" {
		t.Fatalf("[*] DecodeString Error, Input: %s, Got :%s\n", "accaccacc", s)
	}
}
