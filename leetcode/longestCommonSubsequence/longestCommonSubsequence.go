package longestCommonSubsequence

func longestCommonSubsequence(text1 string, text2 string) int {

	m, n := len(text1), len(text2)
	mutrix := make([][]int, m+1)
	mutrix[0] = make([]int, n+1)

	for i := 1; i < m+1; i++ {

		mutrix[i] = make([]int, n+1)
		for j := 1; j < n+1; j++ {

			if text1[i-1] == text2[j-1] {
				mutrix[i][j] = 1 + mutrix[i-1][j-1]
			} else {
				mutrix[i][j] = max(mutrix[i-1][j], mutrix[i][j-1])
			}
		}
	}

	return mutrix[m][n]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
