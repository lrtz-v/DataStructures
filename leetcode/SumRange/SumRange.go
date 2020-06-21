package sumRange

/*
给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。
示例：
给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()
sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3

说明:
你可以假设数组不可变。
会多次调用 sumRange 方法。
*/
type NumArray struct {
	sum []int
}

// Constructor init NumArray
func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums))
    if len(nums) > 0 {
        sum[0] = nums[0]
        for i := 1; i < len(nums); i++ {
            sum[i] = nums[i] + sum[i-1]
        }        
    }
	return NumArray{sum: sum}
}

// SumRange return sum of i to j
func (n *NumArray) SumRange(i int, j int) int {
    if len(n.sum) == 0 {
        return 0
    }
	if i == 0 {
		return n.sum[j]
	}
	return n.sum[j] - n.sum[i-1]
}
