package topKFrequent

import "testing"

func TestTopKFrequent(t *testing.T) {

	nums := []int{1,1,1,2,2,3}
	k := 2
	arr := topKFrequent(nums, k)
	if len(arr) != 2 {
		t.Fatal(arr)
	}
}

func TestTopKFrequent2(t *testing.T) {

	k := 1
	nums := []int{3,0,1,0}

	arr := topKFrequent(nums, k)
	if len(arr) != 1 {
			t.Fatal(arr)
	}
}

func TestTopKFrequent3(t *testing.T) {

	k := 2
	nums := []int{4,1,-1,2,-1,2,3}

	arr := topKFrequent(nums, k)
	if len(arr) != 2 {
			t.Fatal(arr)
	}
}