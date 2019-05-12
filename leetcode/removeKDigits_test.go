package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveKDigits(t *testing.T) {
	testArr := []int{5, 4, 1, 2, 7, 0, 9, 3, 6}
	res := removeKDigits(testArr, 1)
	fmt.Println(res)
}

func TestRemoveKDigits2(t *testing.T) {
	testArr := []int{1, 5, 9, 3, 2, 1, 2}
	res := removeKDigits(testArr, 3)
	fmt.Println(res)
}

func TestRemoveKDigits3(t *testing.T) {
	testArr := []int{1, 0}
	res := removeKDigits(testArr, 2)
	fmt.Println(res)
}
