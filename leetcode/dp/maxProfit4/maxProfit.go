package maxProfit

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {

	n := len(prices)
	if n == 0 {
		return 0
	}

	min := prices[0]
	profit := 0

	for i := 1; i < n; i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i] - min > profit {
			profit = prices[i] - min
		}
	}

	return profit
}
