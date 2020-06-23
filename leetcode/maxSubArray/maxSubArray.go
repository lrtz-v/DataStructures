package maxSubArray

import "math"

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

// 分治法
func maxSubArrayDivide(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return divide(nums, 0, len(nums) - 1)
}

func divide(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	mid := (left + right) / 2

	// 1. 最大数列和在左边
	sumLeft := divide(nums, left, mid)
	// 2. 最大数列和在右边
	sumRight := divide(nums, mid+1, right)
	
	// 3. 最大数列和在中间
	// 先求左边的最大和
	leftSum := 0
	leftMaxSum := math.MinInt32
	for i := mid; i >= left; i-- {
		leftSum += nums[i];
		leftMaxSum = max(leftMaxSum, leftSum)
	}
	// 求右边的最大和
	rightSum := 0
	rightMaxSum := math.MinInt32
	for i := mid + 1; i <= right; i++ {
		rightSum += nums[i]
		rightMaxSum = max(rightMaxSum,rightSum)
	}
	return max(max(sumLeft, sumRight), leftMaxSum + rightMaxSum)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}