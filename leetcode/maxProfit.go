package leetcode

import "fmt"

/*
func maxProfit(prices []int) int {
	n := len(prices)
	max := 0
	reverse := true
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			fmt.Println(i)
			fmt.Println(i - 1)
			reverse = false
			break
		}
	}
	fmt.Println(reverse)
	if reverse {
		return 0
	}
	profitMap := make(map[int]map[int]int, n)
	maxList := make([]int, n)
	for i := 0; i <= n-2; i++ {
		tmp := 0
		for j := i + 1; j <= n-1; j++ {
			if prices[j]-prices[i] >= 0 {
				if _, ok := profitMap[i]; !ok {
					profitMap[i] = make(map[int]int, n-i)
				}
				profitMap[i][j] = prices[j] - prices[i]
				if profitMap[i][j] > tmp {
					tmp = profitMap[i][j]
				}
			}
		}
		maxList[i] = tmp
	}
	for _, profit1 := range profitMap {
		for day2, proft2 := range profit1 {
			stockA := proft2
			stockB := 0
			// check secoed stock
			if day2 <= n-2 {
				for i := day2 + 1; i < n; i++ {
					if maxList[i] > stockB {
						stockB = maxList[i]
					}
				}
			}
			stockA += stockB
			if stockA > max {
				max = stockA
			}
		}
	}
	return max
}
*/

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
			fmt.Printf("start: %d, end: %dval: %d\n", start, end, val)
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
