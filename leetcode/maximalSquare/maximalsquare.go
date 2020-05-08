package maximalsquare

func maximalSquare(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	// init
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	maxSide := 0

	// first row && col
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = int(matrix[i][j])
				continue
			}
			if matrix[i][j] == 0 || matrix[i-1][j] == 0 || matrix[i][j-1] == 0 || matrix[i-1][j-1] == 0 {
				dp[i][j] = int(matrix[i][j])
				continue
			}

			min := dp[i-1][j]
			if min > dp[i][j-1] {
				min = dp[i][j-1]
			}
			if matrix[i-min][j-min] == 0 {
				dp[i][j] = min	
			} else {
				dp[i][j] = 1 + min
			}
			if dp[i][j] > maxSide {
				maxSide = dp[i][j]
			}
		}
	}

	return maxSide*maxSide
}

func min(a, b, c int) int {
	if a < b {
		a = b
	}
	if a < c {
		return c
	}
	return a
}
