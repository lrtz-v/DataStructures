package sort

import "testing"

func TestSort(t *testing.T) {
	q := NewMinPriorityQueue([]int64{})

	l := []int64{1, 4, 9, 2, 10, 0}
	q.Sort(l)
	if !IsSorted(l) {
		t.Fatal(l)
		t.Fatal("Sort failed")
	}

}

func IsSorted(l []int64) bool {
	for i := 1; i < len(l); i++ {
		if l[i] < l[i-1] {
			return false
		}
	}
	return true
}
