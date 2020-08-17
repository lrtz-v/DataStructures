package cuttingrope

func cuttingRope(n int) int {
	maxMap := make(map[int]int, n)
	maxMap[1] = 1
	return dp(maxMap, n)
}

func dp(maxMap map[int]int, n int) int {
	if n < 2 {
		return maxMap[n]
	}

	if _, ok := maxMap[n]; ok {
		return maxMap[n]
	}

	res := 1
	for i := 1; i < n; i++ {
		res = max(max(i * dp(maxMap, n - i), i*(n-i)), res)
	}

	maxMap[n] = res
	return maxMap[n]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
