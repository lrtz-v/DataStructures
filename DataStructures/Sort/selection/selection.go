package selection

import (
	"fmt"
)

// 选择排序
// 首先，找到数组中最小的元素，其次，将它和数组的第一个元素交换
// 在剩下的元素中找到最小的元素，将它与数组对的第二个元素交换，如此往复

// 对于长度为N的数组，需要：
// 比较: (N-1)+(N-2)+(N-3)+...+2+1=N^2/2次
// 交换: N次

// 特点
// 运行时间和输入无关
// 数据移动次数（线性）

type selectSort struct {
	comparableList []int
}

func (s *selectSort) isSorted() bool {
	for i := 1; i < len(s.comparableList); i++ {
		if s.less(i, i-1) {
			return false
		}
	}
	return true
}

func (s *selectSort) less(a, b int) bool {
	return s.comparableList[a] < s.comparableList[b]
}

func (s *selectSort) exch(a, b int) {
	s.comparableList[a], s.comparableList[b] = s.comparableList[b], s.comparableList[a]
}

func (s *selectSort) sort() {
	n := len(s.comparableList)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if s.less(j, min) {
				min = j
			}
		}
		s.exch(i, min)
	}
}

func (s *selectSort) show() {
	fmt.Println("Final Arr: ")
	for i := 0; i < len(s.comparableList); i++ {
		fmt.Println(s.comparableList[i])
	}
}
