package largest1BorderedSquare

import "testing"

func TestLargest1BorderedSquare(t *testing.T) {
	grid := make([][]int, 3)
	grid[0] = []int{1, 1, 1}
	grid[1] = []int{1, 0, 1}
	grid[2] = []int{1, 1, 1}
	s := largest1BorderedSquare(grid)
	if s != 9 {
		t.Fatal("[*] TestLargest1BorderedSquare Error!")
	}
}

func TestLargest1BorderedSquare2(t *testing.T) {
	grid := make([][]int, 2)
	grid[0] = []int{0, 1}
	grid[1] = []int{1, 1}
	s := largest1BorderedSquare(grid)
	if s != 1 {
		t.Fatalf("[*] TestLargest1BorderedSquare Error!, Got %d\n", s)
	}
}

func TestLargest1BorderedSquare3(t *testing.T) {
	grid := make([][]int, 1)
	grid[0] = []int{0}

	s := largest1BorderedSquare(grid)
	if s != 0 {
		t.Fatalf("[*] TestLargest1BorderedSquare Error!, Got %d\n", s)
	}
}