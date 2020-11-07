package sort

/*
选择排序
	首先，找到数组中最小的元素，其次，将它和数组的第一个元素交换
	在剩下的元素中找到最小的元素，将它与数组对的第二个元素交换，如此往复

对于长度为N的数组，需要：
	比较: (N-1)+(N-2)+(N-3)+...+2+1=N^2/2次
	交换: N次

特点
	运行时间和输入无关，乱序数组和有序数组的运行时间一样
	数据移动次数为线性增长
*/

// Selection 选择排序
type Selection struct {
	Template
}

// Sort 选择排序
func (s Selection) Sort(l []int64) {
	n := s.Len(l)
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if s.Less(l, j, minIndex) {
				minIndex = j
			}
		}
		s.Exch(l, i, minIndex)
	}
}
