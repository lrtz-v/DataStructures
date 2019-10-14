package leetcode

import (
	"testing"
)

func TestRemoveKDigits(t *testing.T) {
	t.Skip()
	testArr := "541270936"
	res := removeKDigits(testArr, 1)
	if res != "41270936" {
		t.Errorf("Expected: %s, Got: %s", "41270936", res)
	}
}

func TestRemoveKDigits2(t *testing.T) {
	t.Skip()
	testArr := "1593212"
	res := removeKDigits(testArr, 3)
	if res != "1212" {
		t.Errorf("Expected: %s, Got: %s", "1212", res)
	}
}

func TestRemoveKDigits3(t *testing.T) {
	t.Skip()
	testArr := "10"
	res := removeKDigits(testArr, 2)
	if res != "" {
		t.Errorf("Expected: %s, Got: %s", "", res)
	}
}

func TestRemoveKDigits4(t *testing.T) {
	t.Skip()
	testArr := "1111114132"
	res := removeKDigits(testArr, 2)
	if res != "11111112" {
		t.Errorf("Expected: %s, Got: %s", "11111112", res)
	}
}

func TestRemoveKDigits5(t *testing.T) {
	t.Skip()
	testArr := "1111114132"
	res := removeKDigits(testArr, 3)
	if res != "1111111" {
		t.Errorf("Expected: %s, Got: %s", "1111111", res)
	}
}
