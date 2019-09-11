package leetcode

// dfs
// Time: O(2^(m+n))
// Space O(m+n)
func minPathSum(grid [][]int) int {

	max := 0xffffffff
	n := len(grid)
	m := len(grid[0])

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var calculate func(grid [][]int, i, j int) int

	calculate = func(grid [][]int, i, j int) int {
		if i == n || j == m {
			return max
		} else if i == n-1 && j == m-1 {
			return grid[i][j]
		}
		return grid[i][j] + min(calculate(grid, i+1, j), calculate(grid, i, j+1))
	}

	return calculate(grid, 0, 0)

}

// use dp table
// Time: O(mn)
// Space O(mn)
func minPathSum2(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j != n-1 {
				dp[i][j] = grid[i][j] + dp[i][j+1]
			} else if j == n-1 && i != m-1 {
				dp[i][j] = grid[i][j] + dp[i+1][j]
			} else if i != m-1 && j != n-1 {
				dp[i][j] = grid[i][j] + min(dp[i+1][j], dp[i][j+1])
			} else {
				dp[i][j] = grid[i][j]
			}
		}
	}
	return dp[0][0]
}

// use dp table
// Time: O(mn)
// Space: O(n)
func minPathSum3(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	dp := make([]int, n)

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j != n-1 {
				dp[j] = grid[i][j] + dp[j+1]
			} else if j == n-1 && i != m-1 {
				dp[j] = grid[i][j] + dp[j]
			} else if i != m-1 && j != n-1 {
				dp[j] = grid[i][j] + min(dp[j], dp[j+1])
			} else {
				dp[j] = grid[i][j]
			}
		}
	}
	return dp[0]
}

// use dp table
// Time: O(mn)
// Space: O(1)
func minPathSum4(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j != n-1 {
				grid[i][j] = grid[i][j] + grid[i][j+1]
			} else if j == n-1 && i != m-1 {
				grid[i][j] = grid[i][j] + grid[i+1][j]
			} else if i != m-1 && j != n-1 {
				grid[i][j] = grid[i][j] + min(grid[i+1][j], grid[i][j+1])
			} else {
				// grid[i][j] = grid[i][j]
			}
		}
	}
	return grid[0][0]
}
