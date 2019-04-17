package insertion

import (
	"fmt"
)

// 插入排序
// 将一个元素插入到有序的数组中的合适位置
// 当索引左边的元素都是有序的，但最终位置还不确定，为了给更小的元素腾出空间，它们可能被移动
// 但当索引到达数组的最右边时，数组排序就完成了

// 对于长度为N的数组，需要：
// 比较: N-1 ~ N^2/2次
// 交换: 0 ~ N^2/2次

// 特点
// 运行时间取决于输入的初始顺序
// 交换次数和数组中倒置的数量相同
// 比较次数大于等于倒置的数量，小于等于倒置的数量加数组的大小减一
// 当倒置的数量很少时，效率很高

type insertionSort struct {
	comparableList []int
}

func (s *insertionSort) isSorted() bool {
	for i := 1; i < len(s.comparableList); i++ {
		if s.less(i, i-1) {
			return false
		}
	}
	return true
}

func (s *insertionSort) less(a, b int) bool {
	return s.comparableList[a] < s.comparableList[b]
}

func (s *insertionSort) exch(a, b int) {
	s.comparableList[a], s.comparableList[b] = s.comparableList[b], s.comparableList[a]
}

func (s *insertionSort) sort() {
	n := len(s.comparableList)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && s.less(j, j-1); j-- {
			s.exch(j, j-1)
		}
	}
}

func (s *insertionSort) show() {
	fmt.Println("Final Arr: ")
	for i := 0; i < len(s.comparableList); i++ {
		fmt.Println(s.comparableList[i])
	}
}
