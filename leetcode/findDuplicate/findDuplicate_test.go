package findDuplicate

import (
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	s := findDuplicate([]int{1,3,4,2,2})
	if s != 2 {
		t.Fatalf("[*] FindDuplicate Error! Expected %d, Got %d\n", 2, s)
	}

	s = findDuplicate([]int{3,1,3,4,2})
	if s != 3 {
		t.Fatalf("[*] FindDuplicate Error! Expected %d, Got %d\n", 3, s)
	}
}
