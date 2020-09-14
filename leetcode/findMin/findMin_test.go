package findMin

import "testing"

func TestFindMin(t *testing.T) {
	nums := []int{3,4,5,1,2}
	num := findMin(nums)
	if num != 1 {
		t.Fatal("[*] TestFindMin Error.")
	}
}

func TestFindMin2(t *testing.T) {
	nums := []int{4,5,6,7,0,1,2}
	num := findMin(nums)
	if num != 0 {
		t.Fatal("[*] TestFindMin Error.")
	}
}

func TestFindMin3(t *testing.T) {
	nums := []int{2, 1}
	num := findMin(nums)
	if num != 1 {
		t.Fatal("[*] TestFindMin Error.")
	}
}
