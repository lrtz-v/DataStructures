package addTwoNumbers

import (
	"fmt"
	"testing"
)

func TestCreateList(t *testing.T) {
	res := createNodeList([]int{2, 4, 3})
	// printNodeList(res)
	for _, val := range []int{2, 4, 3} {
		if val != res.Val {
			fmt.Println(res.Val)
			t.Errorf("Expect %d, Get %d\n", val, res.Val)
		}
		res = res.Next
	}
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := createNodeList([]int{2, 4, 3})
	l2 := createNodeList([]int{5, 6, 4})

	res := addTwoNumbers(l1, l2)
	for _, val := range []int{7, 0, 8} {
		if val != res.Val {
			t.Errorf("Expect %d, Get %d\n", val, res.Val)
		}
		res = res.Next
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	l1 := createNodeList([]int{2, 4, 3, 2})
	l2 := createNodeList([]int{5, 6, 4, 1, 5})

	res := addTwoNumbers(l1, l2)
	for _, val := range []int{7, 0, 8, 3, 5} {
		if val != res.Val {
			t.Errorf("Expect %d, Get %d\n", val, res.Val)
		}
		res = res.Next
	}
}

func TestAddTwoNumbers3(t *testing.T) {
	l1 := createNodeList([]int{5})
	// printNodeList(l1)
	l2 := createNodeList([]int{5})
	// printNodeList(l2)

	res := addTwoNumbers(l1, l2)
	for _, val := range []int{0, 1} {
		if val != res.Val {
			t.Errorf("Expect %d, Get %d\n", val, res.Val)
		}
		res = res.Next
	}
}

func TestAddTwoNumbers4(t *testing.T) {
	l1 := createNodeList([]int{1})
	l2 := createNodeList([]int{9, 9})

	res := addTwoNumbers2(l1, l2)
	// printNodeList(res)
	for _, val := range []int{0, 0, 1} {
		if val != res.Val {
			t.Errorf("Expect %d, Get %d\n", val, res.Val)
		}
		res = res.Next
	}
}

func TestAddTwoNumbers5(t *testing.T) {
	l1 := createNodeList([]int{2, 4, 3})
	l2 := createNodeList([]int{5, 6, 4})

	res := addTwoNumbers2(l1, l2)
	for _, val := range []int{7, 0, 8} {
		if val != res.Val {
			t.Errorf("Expect %d, Get %d\n", val, res.Val)
		}
		res = res.Next
	}
}

func TestAddTwoNumbers6(t *testing.T) {
	l1 := createNodeList([]int{2, 4, 3, 2})
	l2 := createNodeList([]int{5, 6, 4, 1, 5})

	res := addTwoNumbers2(l1, l2)
	for _, val := range []int{7, 0, 8, 3, 5} {
		if val != res.Val {
			t.Errorf("Expect %d, Get %d\n", val, res.Val)
		}
		res = res.Next
	}
}

func TestAddTwoNumbers7(t *testing.T) {
	l1 := createNodeList([]int{5})
	// printNodeList(l1)
	l2 := createNodeList([]int{5})
	// printNodeList(l2)

	res := addTwoNumbers2(l1, l2)
	for _, val := range []int{0, 1} {
		if val != res.Val {
			t.Errorf("Expect %d, Get %d\n", val, res.Val)
		}
		res = res.Next
	}
}

func TestAddTwoNumbers11(t *testing.T) {
	// t.Skip()
	l1 := createNodeList([]int{1})
	l2 := createNodeList([]int{9, 9})

	res := addTwoNumbers4(l1, l2)
	for _, val := range []int{0, 0, 1} {
		if val != res.Val {
			t.Errorf("Expect %d, Get %d\n", val, res.Val)
		}
		res = res.Next
	}
}

func TestAddTwoNumbers12(t *testing.T) {
	// t.Skip()
	l1 := createNodeList([]int{2, 4, 3, 2})
	l2 := createNodeList([]int{5, 6, 4, 1, 5})

	res := addTwoNumbers4(l1, l2)
	for _, val := range []int{7, 0, 8, 3, 5} {
		if val != res.Val {
			t.Errorf("Expect %d, Get %d\n", val, res.Val)
		}
		res = res.Next
	}
}

func TestAddTwoNumbers13(t *testing.T) {
	// t.Skip()
	l1 := createNodeList([]int{2, 4, 3})
	l2 := createNodeList([]int{5, 6, 4})

	res := addTwoNumbers4(l1, l2)
	// printNodeList(res)
	for _, val := range []int{7, 0, 8} {
		if val != res.Val {
			t.Errorf("Expect %d, Get %d\n", val, res.Val)
		}
		res = res.Next
	}
}
