package maxSubArray

func maxSubArray(nums []int) int {

	n := len(nums)

	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}

	max := nums[0]
	sum := 0

	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		if sum > max {
			max = sum
		}
	}

	return max
}
