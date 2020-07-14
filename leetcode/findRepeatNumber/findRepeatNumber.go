package findRepeatNumber

func findRepeatNumber(nums []int) int {
	countMap := make(map[int]int)
	for _, i := range nums {
		if _, ok := countMap[i]; ok {
			return i
		}
		countMap[i] = 1
	}

	return -1
}

// todo 使用数组本身当做哈希表来用，降低空间复杂度，但会修改原数组
