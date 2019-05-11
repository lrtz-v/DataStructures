package heapSort

func max(array []int, indexA, indexB int) int {
	if array[indexA] >= array[indexB] {
		return indexA
	}
	return indexB
}

func downAdjust(array []int, parentIndex, length int) {
	// 左节点
	index := 2*parentIndex + 1
	if index >= length {
		return
	}

	// 右节点存在
	if index+1 < length {
		index = max(array, index, index+1)
	}

	// 父节点大于等于任意子节点，跳出
	if array[parentIndex] >= array[index] {
		return
	}
	array[parentIndex], array[index] = array[index], array[parentIndex]
	downAdjust(array, index, length)
}

func heapSort(array []int) {
	for i := (len(array) - 2) / 2; i >= 0; i-- {
		downAdjust(array, i, len(array)-1)
	}
	for i := len(array) - 1; i > 0; i-- {
		array[i], array[0] = array[0], array[i]
		downAdjust(array, 0, i)
	}
}
