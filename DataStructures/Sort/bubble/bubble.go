package bubble

/*
冒泡排序：把相邻的元素两两比较，当一个元素大于右侧相邻元素时，交换它们对的位置；
时间复杂度：O(n^2)
拓展：鸡尾酒排序(先从左到右排序，再从右到左排序，以此反复)，适用于大部分元素已经有序的情况
*/

func bubbleSort(arrayList []int) {
	for i := 0; i < len(arrayList)-1; i++ {
		isSorted := true                 // 没有元素交换说明已经排好序了
		sortBorder := len(arrayList) - 1 // 记录排好序的边界
		tmp := 0
		for j := 0; j < sortBorder; j++ {
			if arrayList[j] > arrayList[j+1] {
				arrayList[j], arrayList[j+1] = arrayList[j+1], arrayList[j]
				isSorted = false
				tmp = j
			}
		}
		sortBorder = tmp
		if isSorted {
			break
		}
	}
}

func cockTailSort(arrayList []int) {

	for i := 0; i < len(arrayList)/2; i++ {
		// 从左往右
		isSorted := true
		sortBorder := len(arrayList) - 1
		for j := i; j < sortBorder; j++ {
			if arrayList[j] > arrayList[j+1] {
				arrayList[j], arrayList[j+1] = arrayList[j+1], arrayList[j]
				isSorted = false
				sortBorder = j
			}
		}
		if isSorted {
			break
		}

		// 从右往左
		isSorted = true
		sortBorder = len(arrayList) - 1
		for j := sortBorder; j > i; j-- {
			if arrayList[j] < arrayList[j-1] {
				arrayList[j], arrayList[j-1] = arrayList[j-1], arrayList[j]
				isSorted = false
				sortBorder = j
			}
		}
		if isSorted {
			break
		}
	}
}
