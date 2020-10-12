package mintime

var visited map[int]bool

func minTime(n int, edges [][]int, hasApple []bool) int {
    if n == 1 {
        return 0
    }
	
	visited = make(map[int]bool)
	
	return dfs(0, edges, hasApple)
}

func dfs(n int, edges [][]int, hasApple []bool) int {
	if visited[n] {
		return 0
	}
	visited[n] = true

	step := 0
	for _, v := range edges {
		if v[0] == n {
			step += dfs(v[1], edges, hasApple)
		} else if v[1] == n {
			step += dfs(v[0], edges, hasApple)
		}
	}

	if n > 0 && (step > 0 || hasApple[n]) {
		return step + 2
	}

	return step
}
