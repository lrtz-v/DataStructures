package findNearestNumber

import (
	"errors"
	"sort"
)

/*
找出给定正整数全排列的下一个数
如给12354，输出为12435

分析：
	逆序组合最大，顺序组合最小
*/

// FindNearestNumber find NearestNumber
func FindNearestNumber(arr []int) ([]int, error) {
	// 判断逆序区间开始位置
	index := len(arr) - 1
	for i := index; i > 0; i-- {
		if arr[i] > arr[i-1] {
			index = i
			break
		}
	}
	if index == 0 {
		return nil, errors.New("already biggest")
	}

	// 交换
	for i := len(arr) - 1; i > 0; i-- {
		if arr[index-1] < arr[i] {
			arr[index], arr[index-1] = arr[index-1], arr[index]
			break
		}
	}

	// 转变成顺序
	sort.Sort(sort.IntSlice(arr[index:]))
	return arr, nil
}
