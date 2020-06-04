package productExceptSelf

import "testing"

func TestProductExceptSelf(t *testing.T) {
	s := productExceptSelf([]int{1, 2, 3, 4})
	if s[0] != 24 || s[1] != 12 || s[2] != 8 || s[3] != 6 {
		t.Fatal("[*] productExceptSelf Error")
	}
}
