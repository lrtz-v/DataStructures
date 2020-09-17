package maxAreaOfIsland

import "testing"

func TestMaxAreaOfIsland(t *testing.T) {
	
	grid := [][]int{}
	grid = append(grid, []int{0,0,1,0,0,0,0,1,0,0,0,0,0})
	grid = append(grid, []int{0,0,0,0,0,0,0,1,1,1,0,0,0})
	grid = append(grid, []int{0,1,1,0,1,0,0,0,0,0,0,0,0})
	grid = append(grid, []int{0,1,0,0,1,1,0,0,1,0,1,0,0})
	grid = append(grid, []int{0,1,0,0,1,1,0,0,1,1,1,0,0})
	grid = append(grid, []int{0,0,0,0,0,0,0,0,0,0,1,0,0})
	grid = append(grid, []int{0,0,0,0,0,0,0,1,1,1,0,0,0})
	grid = append(grid, []int{0,0,0,0,0,0,0,1,1,0,0,0,0})

	max := maxAreaOfIsland(grid)
	if max != 6 {
		t.Fatalf("[*] TestMaxAreaOfIsland Error, Got: %d.\n", max)
	}
}

func TestMaxAreaOfIsland2(t *testing.T) {
	
	grid := [][]int{}
	grid = append(grid, []int{0,0,0,0,0,0,0,0})

	max := maxAreaOfIsland(grid)
	if max != 0 {
		t.Fatalf("[*] TestMaxAreaOfIsland Error, Got: %d.\n", max)
	}
}
