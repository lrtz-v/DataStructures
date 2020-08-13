package cuttingrope

func cuttingRope(n int) int {
	maxMap := make(map[int]int)
	maxMap[1], maxMap[2] = 1, 1

	for i := 1; i <= n; i++ {
		findMax(maxMap, i)
	}

	return maxMap[n]
}

func findMax(maxMap map[int]int, n int) {
	m := 1
	for i := 1; i < n; i++ {
		m = max(maxMap[i] * maxMap[n-i], m)
	}
	maxMap[n] = m
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
