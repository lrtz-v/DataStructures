package massage

func massage(nums []int) int {

	n := len(nums)
	if (n == 0) {
		return 0
	}
	dp0 := 0
	dp1 := nums[0]

	for i := 1; i < n; i++ {
		tdp0 := max(dp0, dp1)
		tdp1 := dp0 + nums[i]

		dp0 = tdp0
		dp1 = tdp1
	}
	return max(dp0, dp1)
}

func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}
