package maxSubArray

func maxSubArray(nums []int) int {
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	m := nums[0]

	for i := 1; i < len(nums); i++ {
		tmp := max(sum[i-1] + nums[i], nums[i])
		sum[i] = tmp
		if tmp > m {
			m = tmp
		}
	}

	return m
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
