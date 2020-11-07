package sort

/*
自顶向下的归并排序
对于长度为N的数组，自顶向下的归并排序需要(NlgN)/2~NlgN次比较，需要访问数组6*(NlgN)次
所需时间与NlgN成正比，所需额外空间与N成正比

自底向上的归并排序
对于长度为N的数组，自底向上的归并排序需要(NlgN)/2~NlgN次比较，最多访问数组6*(NlgN)次
适合链表组织的数据

特点
当数组的长度是2的幂时，自底向上和自顶向下的归并排序所用的比较次数和数组访问次数相同，只是顺序不同
*/

// Merge 归并排序
type Merge struct {
	Template

	aux []int64
}

// Sort 自定向下的归并排序
func (merge Merge) Sort(l []int64) {
	size := merge.Len(l)
	merge.aux = make([]int64, size)
	merge.sort(l, 0, size-1)
}

func (merge Merge) sort(l []int64, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	merge.sort(l, lo, mid)
	merge.sort(l, mid+1, hi)
	merge.merge(l, lo, mid, hi)
}

// SortV2 自底向上的归并排序
func (merge Merge) SortV2(l []int64) {
	size := merge.Len(l)
	merge.aux = make([]int64, size)

	for sz := 1; sz < size; sz += sz {
		for lo := 0; lo < size-sz; lo += sz + sz {
			merge.merge(l, lo, lo+sz-1, min(lo+sz+sz-1, size-1))
		}
	}
}

// 原地归并, 将[lo: mid]与[mid+1:hi]两个有序自数组合并
func (merge Merge) merge(l []int64, lo, mid, hi int) {
	i := lo
	j := mid + 1

	for k := lo; k <= hi; k++ {  // 复制 l[lo ... hi] 到 aux[lo ... hi]
		merge.aux[k] = l[k]
	}

	for k := lo; k <= hi; k++ {
		if i > mid {
			l[k] = merge.aux[j]
			j++
		} else if j > hi {
			l[k] = merge.aux[i]
			i++
		} else if merge.Less(merge.aux, j, i) {
			l[k] = merge.aux[j]
			j++
		} else {
			l[k] = merge.aux[i]
			i++
		}
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
