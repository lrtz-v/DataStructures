package sumRange

import "testing"

func TestSumRange(t *testing.T) {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	n := obj.SumRange(0, 2)
	if n != 1 {
		t.Fatal("[*] TestSumRange Error")
	}

	n = obj.SumRange(2, 5)
	if n != -1 {
		t.Fatal("[*] TestSumRange Error")
	}

	n = obj.SumRange(0, 5)
	if n != -3 {
		t.Fatal("[*] TestSumRange Error")
	}
}