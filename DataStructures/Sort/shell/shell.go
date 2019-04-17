package shell

import (
	"fmt"
)

// 希尔排序
// 基于插入排序的快速排序算法
// 插入排序只会交换相邻的元素，元素只能一点一点的从数组的一端移动到另一端，希尔排序加速了这一过程。
// 使数组中任意间隔为h的元素都是有序的，这样的数组被称为h有序数组
// 如果h很大，我们就能将元素移动到很远的地方
//

// 特点
// 简单实现，对于中等大小的数组其运行时间可接受

type shellSort struct {
	comparableList []int
}

func (s *shellSort) isSorted() bool {
	for i := 1; i < len(s.comparableList); i++ {
		if s.less(i, i-1) {
			return false
		}
	}
	return true
}

func (s *shellSort) less(a, b int) bool {
	return s.comparableList[a] < s.comparableList[b]
}

func (s *shellSort) exch(a, b int) {
	s.comparableList[a], s.comparableList[b] = s.comparableList[b], s.comparableList[a]
}

func (s *shellSort) sort() {
	n := len(s.comparableList)
	h := 1

	for h < n/3 {
		h = 3*h + 1 // 1, 4, 13, 40, 121, 364, 1093
	}

	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && s.less(j, j-h); j -= h {
				s.exch(j, j-h)
			}
		}
		h = h / 3
	}
}

func (s *shellSort) show() {
	fmt.Println("Final Arr: ")
	for i := 0; i < len(s.comparableList); i++ {
		fmt.Println(s.comparableList[i])
	}
}
