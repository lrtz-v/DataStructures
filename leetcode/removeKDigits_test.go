package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveKDigits(t *testing.T) {
	testArr := "541270936"
	res := removeKDigits(testArr, 1)
	fmt.Println(res)
}

func TestRemoveKDigits2(t *testing.T) {
	testArr := "1593212"
	res := removeKDigits(testArr, 3)
	fmt.Println(res)
}

func TestRemoveKDigits3(t *testing.T) {
	testArr := "10"
	res := removeKDigits(testArr, 2)
	fmt.Println(res)
}

func TestRemoveKDigits4(t *testing.T) {
	testArr := "1111114132"
	res := removeKDigits(testArr, 2)
	fmt.Println(res)
}
