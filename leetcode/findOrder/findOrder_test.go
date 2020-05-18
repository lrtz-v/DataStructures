package findOrder

import "testing"

func TestFindOrder(t *testing.T) {
	d := make([][]int, 4)
	d[0] = []int{1, 0}
	d[1] = []int{2, 0}
	d[2] = []int{3, 1}
	d[3] = []int{3, 2}
	r := findOrderDFS(4, d)
	if len(r) != 4 {
		t.Fatal("[*] FindOrder Error!")
	}
	r = findOrderBFS(4, d)
	if len(r) != 4 {
		t.Fatal("[*] FindOrder Error!")
	}
}


func TestFindOrder2(t *testing.T) {
	d := make([][]int, 1)
	d[0] = []int{1, 0}
	r := findOrderDFS(2, d)
	if len(r) != 2 {
		t.Fatal("[*] FindOrder Error!")
	}
	r = findOrderBFS(2, d)
	if len(r) != 2 {
		t.Fatal("[*] FindOrder Error!")
	}
}


func TestFindOrder3(t *testing.T) {
	d := make([][]int, 0)
	r := findOrderDFS(2, d)
	if len(r) != 2 {
		t.Fatal("[*] FindOrder Error!")
	}
	r = findOrderBFS(2, d)
	if len(r) != 2 {
		t.Fatal("[*] FindOrder Error!")
	}
}