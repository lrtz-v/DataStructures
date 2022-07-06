package ch1

import "testing"

func TestQuickUnion(t *testing.T) {
	idPairs := [][]int{
		[]int{4, 3},
		[]int{3, 8},
		[]int{6, 5},
		[]int{6, 8},
		[]int{1, 3},
	}

	max := 8
	ids := []int{}
	for i := 0; i <= max; i++ {
		ids = append(ids, i)
	}

	u := QuickUnion{ids: ids}
	for _, v := range idPairs {
		if u.Connected(v[0], v[1]) {
			continue
		}
		u.Union(v[0], v[1])
	}
	if !u.Connected(5, 8) {
		t.Fatalf("check connected failed on %d & %d\n", 5, 8)
	}
	if !u.Connected(1, 3) {
		t.Fatalf("check connected failed on %d & %d\n", 1, 3)
	}
}
