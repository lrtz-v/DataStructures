package longestCommonSubsequence

// 暴力破解(超时警告⚠️)
func longestCommonSubsequence(text1 string, text2 string) int {

	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				return longestCommonSubsequence(text1[i+1:], text2[j+1:]) + 1
			}

			return max(longestCommonSubsequence(text1[i:], text2[j+1:]), longestCommonSubsequence(text1[i+1:], text2[j:]))
		}
	}

	return 0
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// use dp
func longestCommonSubsequence2(text1 string, text2 string) int {

	m, n := len(text1), len(text2)

	// dp[i][j] 的含义是：对于 s1[1..i] 和 s2[1..j]，它们的 LCS 长度是 dp[i][j]
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]

}

/*
# 初始化 base case
dp[0][0][...] = base
# 进行状态转移
for 状态1 in 状态1的所有取值：
    for 状态2 in 状态2的所有取值：
        for ...
			dp[状态1][状态2][...] = 求最值(选择1，选择2...)

*/
