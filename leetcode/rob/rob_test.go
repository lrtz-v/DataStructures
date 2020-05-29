package rob

import "testing"

func TestRob(t *testing.T) {
	r := rob([]int{1,2,3,1})
	if r != 4 {
		t.Fatalf("[*] Rob Failed! Expected 4, Got %d\n", r)
	}
}

func TestRob2(t *testing.T) {
	r := rob([]int{2,7,9,3,1})
	if r != 12 {
		t.Fatalf("[*] Rob Failed! Expected 12, Got %d\n", r)
	}
}

func TestRob3(t *testing.T) {
	r := rob([]int{114,117,207,117,235,82,90,67,143,146,53,108,200,91,80,223,58,170,110,236,81,90,222,160,165,195,187,199,114,235,197,187,69,129,64,214,228,78,188,67,205,94,205,169,241,202,144,240})
	if r != 12 {
		t.Fatalf("[*] Rob Failed! Expected 12, Got %d\n", r)
	}
}