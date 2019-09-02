package lis

import (
	"testing"
)

func TestLongestIncreasingSubsequence(t *testing.T) {
	t.Skip()
	testArray := []int{1, 2, 4, 7, 1, 2}
	lis := longestIncreasingSubsequence(testArray, len(testArray))
	if lis != 4 {
		t.Errorf("Expect %d, Get %d\n", 4, lis)
	}
}

func TestLongestIncreasingSubsequence2(t *testing.T) {
	testArray := []int{1, 2, 1, 4, 7, 1, 2}
	lis := longestIncreasingSubsequence2(testArray, len(testArray))
	if lis != 3 {
		t.Errorf("Expect %d, Get %d\n", 3, lis)
	}
}
