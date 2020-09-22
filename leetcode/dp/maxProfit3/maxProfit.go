package maxProfit

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	var dfs func(prices []int, nums int, has bool) int
	dfs = func(prices []int, nums int, has bool) int {
		if len(prices) == 0 || nums == 2 {
			return 0
		}

		// 买入，卖出，观望
		a, b, c := 0, 0, 0

		// 观望 
		c = dfs(prices[1:], nums, has)

		if has {
			// 变为卖出+观望
			b = prices[0] + dfs(prices[1:], nums+1, false)
			return max(b, c)
		}

		// 变为买入+观望
		a = -prices[0] + dfs(prices[1:], nums, true)
		return max(a, c)
	}

	return dfs(prices, 0, false)
}
