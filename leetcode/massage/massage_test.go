package massage

import "testing"

func TestMassage(t *testing.T) {
	n := massage([]int{1,2,3,1})
	if n != 4 {
		t.Fatal("[*] TestMassage Error!")
	}
}
