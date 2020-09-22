package maxProfit

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 空间优化
func maxProfit2(prices []int, fee int) int {

	if len(prices) == 0 {
		return 0
	}

	n := len(prices)
	var dp [2]int
	dp[0] = -prices[0]

	for i := 1; i < n; i++ {
		tmp := dp[1]
		// 变为非持有：max(今天卖出, 前一天也是非持有)
		dp[1] = max(dp[0] + prices[i] - fee, tmp)

		// 变为持有： max(前一天买的， 今天买=前一天非持有-今天价格)
		dp[0] = max(dp[0], tmp - prices[i])
	}

	return dp[1]
}