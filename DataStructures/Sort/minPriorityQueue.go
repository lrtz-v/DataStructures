package sort

import "errors"

/*
最小堆

堆排序的基本思想是：
1.将待排序序列构造成一个大顶堆，此时，整个序列的最大值就是堆顶的根节点
2.将其与末尾元素进行交换，此时末尾就为最大值
3.然后将剩余n-1个元素重新构造成一个堆，这样会得到n个元素的次小值。如此反复执行，便能得到一个有序序列了

PS：一般升序采用大顶堆，降序采用小顶堆
*/

// MinPriorityQueue struct
type MinPriorityQueue struct {
	PriorityQueue
}

// NewMinPriorityQueue create new MinPriorityQueue
func NewMinPriorityQueue(q []int64) *MinPriorityQueue {
	l := []int64{0}
	if len(q) > 0 {
		l = append(l, q...)
	}
	mq := &MinPriorityQueue{}
	mq.queue = l
	mq.count = len(q)
	return mq
}

// Swim 上浮堆化, 父节点小于字节点
func (pq *MinPriorityQueue) Swim(k int) {
	for k > 1 && pq.Less(k, k/2) {
		pq.Exch(k/2, k)
		k /= 2
	}
}

// Sink 下沉堆化，父节点小于字节点
func (pq *MinPriorityQueue) Sink(k, size int) {
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
func (pq *MinPriorityQueue) Insert(v int64) {
	pq.queue = append(pq.queue, v)
	pq.count++
	pq.Swim(pq.count)
}

// DelMin 删除最小元素
func (pq *MinPriorityQueue) DelMin() (int64, error) {
	if pq.count > 0 {
		min := pq.queue[1]
		pq.Exch(1, pq.count)
		pq.queue = pq.queue[:len(pq.queue)-1]
		pq.count--
		pq.Sink(1, pq.Len())
		return min, nil
	}
	return 0, errors.New("MaxPriorityQueue is Empty")
}

// GetMin 查询最小值
func (pq *MinPriorityQueue) GetMin() (int64, error) {
	if pq.count > 0 {
		return pq.queue[1], nil
	}
	return 0, errors.New("MaxPriorityQueue is Empty")
}
