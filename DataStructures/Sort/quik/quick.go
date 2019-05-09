package quik

import (
	"github.com/DataStructures/DataStructures/Stack"
)

/*
快速排序
每一轮挑选一个基准元素，并让其他比它大的元素移动到数列的一边，比它小的元素移动到数列的另一边
从而把数列拆分成两个部分
平均时间复杂度：O(nlogn), 最差：O(n^2)
*/

// partition 双边循环 递归
func partition(arr []int, startIndex, endIndex int) int {
	pivot := arr[startIndex]
	left := startIndex
	right := endIndex

	for left != right {
		for left < right && arr[right] > pivot {
			right--
		}
		for left < right && arr[left] <= pivot {
			left++
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	// left == right
	arr[startIndex], arr[left] = arr[left], arr[startIndex]
	return left
}

// 单边循环
func partitionV2(arr []int, startIndex, endIndex int) int {
	pivot := arr[startIndex]
	mark := startIndex

	for i := startIndex + 1; i <= endIndex; i++ {
		if arr[i] < pivot {
			mark++
			arr[mark], arr[i] = arr[i], arr[mark]
		}
	}

	arr[startIndex], arr[mark] = arr[mark], arr[startIndex]
	return mark
}

func quickSort(arr []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	pivotIndex := partitionV2(arr, startIndex, endIndex)
	quickSort(arr, startIndex, pivotIndex-1)
	quickSort(arr, pivotIndex+1, endIndex)
}

// quickSort2 使用栈代替递归
func quickSort2(arr []int, startIndex, endIndex int) {
	stack := Stack.NewStack()
	rootParam := make(map[string]int)
	rootParam["startIndex"] = startIndex
	rootParam["endIndex"] = endIndex
	stack.Push(rootParam)

	for !stack.Empty() {
		data, err := stack.Pop()
		if err != nil {
			panic("Pop Error")
		}
		param := data.(map[string]int)
		pivotIndex := partitionV2(arr, param["startIndex"], param["endIndex"])

		if param["startIndex"] < pivotIndex-1 {
			leftParam := make(map[string]int)
			leftParam["startIndex"] = param["startIndex"]
			leftParam["endIndex"] = pivotIndex - 1
			stack.Push(leftParam)
		}

		if pivotIndex+1 < param["endIndex"] {
			rightParam := make(map[string]int)
			rightParam["startIndex"] = pivotIndex + 1
			rightParam["endIndex"] = param["endIndex"]
			stack.Push(rightParam)
		}
	}
}
