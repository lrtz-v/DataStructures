package sort

type MinPriorityQueue struct {
	PriorityQueue
}

// 上浮堆化
func (pq MinPriorityQueue) Swim(k int) {
	for k > 1 && pq.Less(k/2, k) {
		pq.Exch(k/2, k)
		k /= 2
	}
}

// 下沉堆化
func (pq MinPriorityQueue) Sink(k int) {
	size := pq.Len()
	for 2*k <= size {
		j := 2 * k
		if j < size && pq.Less(j, j+1) {
			j++
		}
		if !pq.Less(k, j) {
			break
		}
		pq.Exch(k, j)
		k = j
	}
}

func (pq MinPriorityQueue) Insert(v int64) {}

func (pq MinPriorityQueue) DelMin(v int64) int64 {
	return 0
}

// Sort 最小堆排序
func (pq MinPriorityQueue) Sort(l []int64) {
	size := pq.Len()
	for k := size / 2; k >= 1; k-- {
		sink(l, k, size)
	}
	for size > 1 {
		pq.Exch(1, size)
		size--
		sink(l, 1, size)
	}
}

func sink(l []int64, k, n int) {
	for 2*k <= n {
		j := 2 * k
		if j < n && l[j] < l[j+1] {
			j++
		}
		if l[k] >= l[j] {
			break
		}
		l[k], l[j] = l[j], l[k]
		k = j
	}
}