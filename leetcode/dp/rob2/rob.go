package rob2

/*
所有的房屋都围成一圈;
两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警;
*/
func rob2(nums []int) int {

    n := len(nums)
    if n == 0 {
        return 0
    } else if n == 1 {
        return nums[0]
    }
    
    var robCircle func(nums []int) int
    robCircle = func(nums []int) int {
		n := len(nums)
		if n == 0 {
			return 0
		} else if n == 1 {
			return nums[0]
		}
	
		first := nums[0]
		second := max(nums[0], nums[1])
	
		for i := 2; i < n; i++ {
			first, second = second, max(nums[i]+first, second)
		}
	
		return max(first, second)
    }

    return max(robCircle(nums[1:]), robCircle(nums[:n-1]))
}

func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}