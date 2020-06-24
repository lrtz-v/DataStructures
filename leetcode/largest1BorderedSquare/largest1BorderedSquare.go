package largest1BorderedSquare

func largest1BorderedSquare(grid [][]int) int {
	t := make([][]int, len(grid))
	t1 := make([][]int, len(grid))
	for i := 0; i < len(t); i++ {
		t[i] = make([]int, len(grid[0]))
		t1[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 && j > 0 {
				t[i][j] = t[i][j-1] + grid[i][j]
			} else {
				t[i][j] = grid[i][j]
			}
		}
	}

	for j := 0; j < len(grid[0]); j++ {
		for i := 0; i < len(grid); i++ {
			if grid[i][j] == 1 && i > 0 {
				t1[i][j] = t1[i-1][j] + grid[i][j]
			} else {
				t1[i][j] = grid[i][j]
			}
		}
	}

	max := 0
	for i := 1; i < len(t); i++ {
		for j := 1; j < len(t[0]); j++ {
			if grid[i][j] > 0 {
				if grid[i][j] > max {
					max = grid[i][j]
				}
				if t[i][j] == t1[i][j] && grid[i-t[i][j]+1][j-t[i][j]+1] > 0 && t[i][j] > max {
					max = t[i][j]
				}
			}
		}
	}
	return max*max
}