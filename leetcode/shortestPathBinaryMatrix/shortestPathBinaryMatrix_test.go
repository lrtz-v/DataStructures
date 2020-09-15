package shortestPathBinaryMatrix

import "testing"

func TestShortestPathBinaryMatrix(t *testing.T) {
	grid := make([][]int, 2)
	grid[0] = []int{0, 1}
	grid[1] = []int{1, 0}

	n := shortestPathBinaryMatrix(grid)
	if n != 2 {
		t.Fatalf("[*] TestShortestPathBinaryMatrix Error. Got: %d\n", n)
	}
}

func TestShortestPathBinaryMatrix2(t *testing.T) {
	grid := make([][]int, 3)
	grid[0] = []int{0, 0, 0}
	grid[1] = []int{1, 1, 0}
	grid[2] = []int{1, 1, 0}

	n := shortestPathBinaryMatrix(grid)
	if n != 4 {
		t.Fatalf("[*] TestShortestPathBinaryMatrix Error. Got: %d\n", n)
	}
}

func TestShortestPathBinaryMatrix3(t *testing.T) {
	grid := make([][]int, 3)
	grid[0] = []int{1, 0, 0}
	grid[1] = []int{1, 1, 0}
	grid[2] = []int{1, 1, 0}

	n := shortestPathBinaryMatrix(grid)
	if n != -1 {
		t.Fatalf("[*] TestShortestPathBinaryMatrix Error. Got: %d\n", n)
	}
}