package twoSum

func twoSum(nums []int, target int) []int {
	numHash := make(map[int]int)
	for index, value := range nums {
		numHash[value] = index
	}
	for index, value := range nums {
		tmpVal := target - value
		if v, ok := numHash[tmpVal]; ok {
			if index != v {
				return []int{index, v}
			}
		}
	}

	return nil
}
