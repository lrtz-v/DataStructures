package twoSum

func twoSum(nums []int, target int) []int {
	numIndex := make(map[int]int)
	indexList := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		tmp := target - num
		if _, ok := numIndex[tmp]; ok {
			if numIndex[tmp] != i {
				indexList[0] = numIndex[tmp]
				indexList[1] = i
				break
			}
		} else {
			numIndex[num] = i
		}
	}
	return indexList
}
