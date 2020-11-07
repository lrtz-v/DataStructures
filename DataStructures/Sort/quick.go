package sort

/*
快速排序

每一轮挑选一个基准元素，并让其他比它大的元素移动到数列的一边，比它小的元素移动到数列的另一边
从而把数列拆分成两个部分
时间复杂度：O(nlogn) - O(n^2)
空间复杂度:O(logn)
*/

// Quick sort
type Quick struct {
	Template
}

// Sort work
func (quick Quick) Sort(l []int64) {
	size := quick.Len(l)
	quick.sort(l, 0, size-1)
}

func (quick Quick) sort(l []int64, lo, hi int) {
	if hi <= lo {
		return
	}

	lt, gt := quick.partition(l, lo, hi) // 关键在于切分操作
	quick.sort(l, lo, lt-1)
	quick.sort(l, gt+1, hi)
	// 不需要merge操作，已经有序
}

func (quick Quick) partition(l []int64, lo, hi int) (int, int) {

	lt := lo
	gt := hi

	i := lo + 1
	v := l[lo]

	for i <= gt {
		if quick.Equal(l[i], v) {
			i++
		} else if quick.Less(l[i], v) {
			quick.Exch(l, lt, i)
			lt++
			i++
		} else {
			quick.Exch(l, i, gt)
			gt--
		}
	}
	return lt, gt
}
