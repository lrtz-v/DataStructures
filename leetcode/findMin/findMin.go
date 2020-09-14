package findMin

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	// 7, 8, 9, 1, 2,    3, 4, 5, 6 在左边
	// 5, 6, 7, 8, 9,    1, 2  在右边
	// 7, 8, 9,          1, 2, 3  在中间

	left := 0
	right := len(nums) - 1

	for {
		mid := left + (right - left) / 2
		if mid > 0 && nums[mid] < nums[mid - 1] {
			return nums[mid]
		} else if mid + 1 <= right && nums[mid] > nums[mid + 1] {
			return nums[mid + 1]
		} else if nums[mid] > nums[right] {
			left = mid+1
		} else if nums[mid] < nums[left] {
			right = mid-1
		}
	}
}
