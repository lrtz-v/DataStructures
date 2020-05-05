package dynamic

// 使用动态规划解决0-1背包问题, 找到最大能装的重量
// weight:物品重量，n:物品个数，w:背包可承载重量
func knapsack(weight []int, n, w int) int {
	states := make([][]bool, n)
	for i := 0; i < n; i++ {
		states[i] = make([]bool, w+1)
	}

	states[0][0] = true
	if weight[0] <= w {
		states[0][weight[0]] = true
	}

	for i := 1; i < n; i++ { // 动态规划状态转移
		for j := 0; j <= w; j++ { // 不把第i个物品放入背包
			if states[i-1][j] == true {
				states[i][j] = states[i-1][j]
			}
		}
		for j := 0; j <= w-weight[i]; j++ { //把第i个物品放入背包
			if states[i-1][j] == true {
				states[i][j+weight[i]] = true
			}
		}
	}

	for i := w; i >= 0; i-- { // 输出结果
		if states[n-1][i] == true {
			return i
		}
	}
	return 0
}

// 使用一位数组优化空间
func knapsack2(weight []int, n, w int) int {
	states := make([]bool, w+1)
	states[0] = true // 第一行的数据要特殊处理，可以利用哨兵优化

	if weight[0] <= w {
		states[weight[0]] = true
	}

	for i := 1; i < n; i++ { // 动态规划
		for j := w - weight[i]; j >= 0; j-- { //把第i个物品放入背包
			if states[j] == true {
				states[j+weight[i]] = true
			}
		}
	}

	for i := w; i >= 0; i-- { // 输出结果
		if states[i] == true {
			return i
		}
	}
	return 0
}

// 加入物品价值唯独，在重量允许下，找到价值最大
func knapsack3(weight, value []int, n, w int) int {
	states := make([][]int, n)
	for i := 0; i < n; i++ {
		states[i] = make([]int, w+1)
	}

	// init states
	for i := 0; i < n; i++ {
		for j := 0; j < w+1; j++ {
			states[i][j] = -1
		}
	}

	states[0][0] = 0
	if weight[0] <= w {
		states[0][weight[0]] = value[0]
	}

	for i := 1; i < n; i++ { //动态规划，状态转移

		for j := 0; j <= w; j++ { // 不选择第i个物品
			if states[i-1][j] >= 0 {
				states[i][j] = states[i-1][j]
			}
		}

		for j := 0; j <= w-weight[i]; j++ { // 选择第i个物品
			if states[i-1][j] >= 0 {
				v := states[i-1][j] + value[i]
				if v > states[i][j+weight[i]] {
					states[i][j+weight[i]] = v
				}
			}
		}
	}

	// 找出最大值
	maxvalue := -1
	for j := 0; j <= w; j++ {
		if states[n-1][j] > maxvalue {
			maxvalue = states[n-1][j]
		}
	}
	return maxvalue
}
