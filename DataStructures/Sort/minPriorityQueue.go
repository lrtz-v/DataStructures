package sort

/*
* 最小堆
* */

// MinPriorityQueue struct
type MinPriorityQueue struct {
	PriorityQueue
}

// 上浮堆化, 父节点小于字节点
func (pq MinPriorityQueue) Swim(k int) {
	for k > 1 && pq.Less(k, k/2) {
		pq.Exch(k/2, k)
		k /= 2
	}
}

// 下沉堆化，父节点小于字节点
func (pq MinPriorityQueue) Sink(k int) {
	size := pq.Len()
	for 2*k <= size {
		j := 2 * k
		// 找到字节点最小值
		if j < size && pq.Less(j+1, j) {
			j++
		}
		// 比较父节点与子节点
		if pq.Less(k, j) {
			break
		}
		// 交换
		pq.Exch(k, j)
		k = j
	}
}

// Insert 插入元素，上浮调整
func (pq MinPriorityQueue) Insert(v int64) {
	pq.queue = append(pq.queue, v)
	pq.count++
	pq.Swim(pq.count)
}

// DelMin 删除最小元素
func (pq MinPriorityQueue) DelMin(v int64) int64 {
	min := pq.queue[1]
	pq.Exch(1, pq.count)
	pq.Sink(1)
	return min
}

// Sort 最小堆排序G
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
