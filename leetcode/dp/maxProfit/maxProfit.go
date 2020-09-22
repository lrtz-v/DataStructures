package maxProfit

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	n := len(prices)
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	
	for i := 1; i < n; i++ {

		// 变为持有： max(前一天买的， 今天买=前一天冷冻期-今天价格)
		dp[i][0] = max(dp[i-1][0], dp[i-1][2] - prices[i])
		
		// 变为冻结：今天卖出
		dp[i][1] = dp[i-1][0] + prices[i]

		// 变为非冻结: max(前一天冻结，前一天也是非冻结)
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}

	return max(dp[n-1][1], dp[n-1][2])
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 空间优化
func maxProfit2(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	n := len(prices)
	var dp [3]int
	dp[0] = -prices[0]
	
	for i := 1; i < n; i++ {

		tmp := dp[2]
		// 变为非冻结: max(前一天冻结，前一天也是非冻结)
		// dp[i][2] = max(dp[i-1][1], dp[i-1][2])
		dp[2] = max(dp[1], tmp)

		// 变为冻结：今天卖出
		// dp[i][1] = dp[i-1][0] + prices[i]
		dp[1] = dp[0] + prices[i]

		// 变为持有： max(前一天买的， 今天买=前一天冷冻期-今天价格)
		// dp[i][0] = max(dp[i-1][0], dp[i-1][2] - prices[i])
		dp[0] = max(dp[0], tmp - prices[i])
	}

	return max(dp[1], dp[2])
}