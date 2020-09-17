package maxAreaOfIsland

func maxAreaOfIsland(grid [][]int) int {

	N := len(grid)
	if N == 0 {
		return 0
	}
	M := len(grid[0])
	visited := make([][]bool, N)
	for i := 0; i < N; i++ {
		tmp := []bool{}
		for j := 0; j < M; j++ {
			tmp = append(tmp, false)
		}
		visited[i] = tmp
	}

	max := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if grid[i][j] == 1 {
				tmp := dfs(grid, i, j , visited, N, M)
				if tmp > max {
					max = tmp
				}
			}
		}
	}

	return max
}

func dfs(grid [][]int, i, j int, visited [][]bool, N, M int) int {
	if visited[i][j] || grid[i][j] == 0 {
		return 0
	}
	visited[i][j] = true

	size := grid[i][j]

	if i - 1 >= 0 {
		size += dfs(grid, i-1, j, visited, N, M)
	}
	if i + 1 < N {
		size += dfs(grid, i+1, j, visited, N, M)
	} 
	if j - 1 >= 0 {
		size += dfs(grid, i, j - 1, visited, N, M)
	}
	if j + 1 < M {
		size += dfs(grid, i, j+1, visited, N, M)
	} 
	return size
}
