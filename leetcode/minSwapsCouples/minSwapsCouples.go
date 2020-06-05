package minSwapsCouples

func minSwapsCouples(row []int) int {

	swapTimes := 0
	indexMap := make(map[int]int)

	for i := 0; i < len(row); i++ {
		indexMap[row[i]] = i
	}

	aim := row[0]
	aimIndex := 0
	for i := 0; i < len(row)-1; i++ {
		if row[i]%2 == 0 {
			aim = row[i] + 1
		} else {
			aim = row[i] - 1
		}
		aimIndex = indexMap[aim]
		if aimIndex-i == 1 || aimIndex-i == -1 {
			continue
		} else {
			row[i+1], row[aimIndex] = row[aimIndex], row[i+1]
			indexMap[row[i+1]] = i + 1
			indexMap[row[aimIndex]] = aimIndex
			swapTimes++
		}
	}

	return swapTimes
}
