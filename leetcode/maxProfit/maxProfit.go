package maxProfit

func maxProfit(prices []int) int {
    l := len(prices)
    if l == 0 {
        return 0
    }
    m := prices[0]
    
    profit := 0
    for i := 1; i < l; i++ {
        if prices[i] - m > profit {
            profit = prices[i] - m
        }
        m = min(m, prices[i])
    }
    
    return profit
}

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}
