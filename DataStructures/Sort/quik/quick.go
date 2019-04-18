package quik

import (
	"fmt"
)

// 快速排序

// 特点
//

type quikSort struct {
	comparableList []int
	auk            []int
}

func (s *quikSort) isSorted() bool {
	for i := 1; i < len(s.comparableList); i++ {
		if s.less(s.comparableList[i], s.comparableList[i-1]) {
			return false
		}
	}
	return true
}

func (s *quikSort) less(a, b int) bool {
	return a < b
}

func (s *quikSort) sort(lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	s.sort(lo, mid)
	s.sort(mid+1, hi)
}

func (s *quikSort) show() {
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
