package productExceptSelf

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	
	left[0] = nums[0]
	right[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		left[i] = nums[i]*left[i-1]
	}
	for i := len(nums)-2; i > 0; i-- {
		right[i] = nums[i]*right[i+1]
	}

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = right[i+1]
		} else if i == len(nums)-1 {
			result[i] = left[i-1]
		} else {
			result[i] = left[i-1]*right[i+1]
		}
	}

	return result
}
