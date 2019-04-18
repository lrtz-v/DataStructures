package merge

import (
	"fmt"
)

// 自顶向下的归并排序
// 对于长度为N的数组，自顶向下的归并排序需要(NlgN)/2~NlgN次比较，需要访问数组6*(NlgN)次
// 所需时间与NlgN成正比，所需额外空间与N成正比

// 自底向上的归并排序
// 对于长度为N的数组，自底向上的归并排序需要(NlgN)/2~NlgN次比较，最多访问数组6*(NlgN)次
// 适合链表组织的数据，
//

// 特点
// 当数组的长度是2的幂时，自底向上和自顶向下的归并排序所用的比较次数和数组访问次数相同，只是顺序不同

type mergeSort struct {
	comparableList []int
	auk            []int
}

func (s *mergeSort) isSorted() bool {
	for i := 1; i < len(s.comparableList); i++ {
		if s.less(s.comparableList[i], s.comparableList[i-1]) {
			return false
		}
	}
	return true
}

func (s *mergeSort) less(a, b int) bool {
	return a < b
}

func (s *mergeSort) merge(lo, mid, hi int) {
	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		s.auk[k] = s.comparableList[k]
	}

	for k := lo; k <= hi; k++ {
		if i > mid {
			s.comparableList[k] = s.auk[j]
			j++
		} else if j > hi {
			s.comparableList[k] = s.auk[i]
			i++
		} else if s.less(s.auk[j], s.auk[i]) {
			s.comparableList[k] = s.auk[j]
			j++
		} else {
			s.comparableList[k] = s.auk[i]
			i++
		}
	}
}

// 自上向下
func (s *mergeSort) sortD(lo, hi int) {
	s.auk = make([]int, len(s.comparableList))
	s.sortDown(0, len(s.comparableList)-1)
}

func (s *mergeSort) sortDown(lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	s.sortDown(lo, mid)
	s.sortDown(mid+1, hi)
	s.merge(lo, mid, hi)
}

// 自底向上
func (s *mergeSort) sortU() {
	n := len(s.comparableList)
	s.auk = make([]int, n)
	for sz := 1; sz < n; sz += sz {
		for lo := 0; lo < n-sz; lo += sz + sz {
			s.merge(lo, lo+sz-1, min(lo+sz+sz-1, n-1))
		}
	}
}

func (s *mergeSort) show() {
	fmt.Println("Final Arr: ")
	for i := 0; i < len(s.comparableList); i++ {
		fmt.Println(s.comparableList[i])
	}
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
