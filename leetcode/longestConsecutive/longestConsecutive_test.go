package longestConsecutive

import "testing"

func TestLongestConsecutive(t *testing.T) {
	var l int
	l = longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	if l != 4 {
		t.Fatal("[*] longestConsecutive error!")
	}

	l = longestConsecutive([]int{1, 2, 0, 1})
	if l != 3 {
		t.Fatal("[*] longestConsecutive error2!")
	}

	l = longestConsecutive([]int{0,3,7,2,5,8,4,6,0,1})
	if l != 9 {
		t.Fatal("[*] longestConsecutive error3!")
	}
}
