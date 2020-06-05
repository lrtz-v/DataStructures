package maxProfit

func maxProfit(prices []int) int {

	n := len(prices)

	start := 0
	end := 0
	first := 0
	second := 0

	for i := 1; i < n; i++ {
		if prices[i] < prices[i-1] || i == n-1 {
			if prices[i] < prices[i-1] {
				end = i - 1
			} else {
				end = i
			}
			val := prices[end] - prices[start]
			if val > first && val > second {
				first, second = val, first
			} else if val > second {
				second = val
			}
			start = i
		}
	}

	return first + second

}
