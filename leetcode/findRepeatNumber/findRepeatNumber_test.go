package findRepeatNumber

import "testing"

func TestFindRepeatNumber(t *testing.T) {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	n := findRepeatNumber(nums)
	if n != 2 && n != 3 {
		t.Fatalf("[*] findRepeatNumber error! Expected 2/3, Got %d\n.", n)
	}
}
