package rob

func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	first := nums[0]
	second := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		first, second = second, max(nums[i]+first, second)
	}

	return max(first, second)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
