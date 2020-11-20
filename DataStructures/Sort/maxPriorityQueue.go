package sort

import (
	"errors"
)

/*
优先级队列
插入元素：log(n)
删除元素：log(n)
排序：nlog(n)
*/

// MaxPriorityQueue struct
type MaxPriorityQueue struct {
	PriorityQueue
}

// NewMaxPriorityQueue create new MaxPriorityQueue
func NewMaxPriorityQueue(q []int64) *MaxPriorityQueue {
	l := []int64{0}
	if len(q) > 0 {
		l = append(l, q...)
	}
	mq := &MaxPriorityQueue{}
	mq.queue = l
	mq.count = len(q)
	return mq
}

// Swim 上浮堆化
// 父节点需要大于字节点
func (pq *MaxPriorityQueue) Swim(k int) {
	for k > 1 && pq.Less(k/2, k) {
		pq.Exch(k/2, k)
		k /= 2
	}
}

// Sink 下沉堆化
func (pq *MaxPriorityQueue) Sink(k, size int) {
	for 2*k <= size {
		j := 2 * k
		// 找到左右字节点最大的一个
		if j < size && pq.Less(j, j+1) {
			j++
		}
		// 是否比父节点大
		if !pq.Less(k, j) {
			break
		}
		pq.Exch(k, j)
		k = j
	}
}

// Insert 插入节点：将元素放到列表尾部，然后上浮
func (pq *MaxPriorityQueue) Insert(v int64) {
	pq.queue = append(pq.queue, v)
	pq.count++
	pq.Swim(pq.count)
}

// DelMax 删除最大节点：将最大元素与列表尾部元素交换，然后下沉首元素
func (pq *MaxPriorityQueue) DelMax() (int64, error) {
	if pq.count > 0 {
		max := pq.queue[1]
		pq.Exch(1, pq.count)
		pq.queue = pq.queue[:len(pq.queue)-1]
		pq.count--
		pq.Sink(1, pq.Len())
		return max, nil
	}
	return 0, errors.New("MaxPriorityQueue is Empty")
}

// GetMax get max value in MaxPriorityQueue
func (pq *MaxPriorityQueue) GetMax() (int64, error) {
	if pq.count > 0 {
		return pq.queue[1], nil
	}
	return 0, errors.New("MaxPriorityQueue is Empty")
}

// Sort 堆排序
func (pq *MaxPriorityQueue) Sort() {
	// 无序序列结构: pq.queue[1:pq.Len()]
	size := pq.Len()

	// 将给定无序序列构造成一个大顶堆
	for k := size / 2; k >= 1; k-- {
		// 从第一个非叶子结点从下至上，从右至左调整结构
		pq.Sink(k, size)
	}

	for ; size > 1; size-- {
		pq.Exch(1, size)
		pq.Sink(1, size-1)
	}
}
