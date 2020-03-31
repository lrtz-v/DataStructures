package binaryheap

import (
	"strconv"
	"strings"
)

/**
二叉堆
构建时间复杂度：O(n)
*/

// BinaryHeap 二叉堆
type BinaryHeap struct {
	heap []int
	len  int
}

func CreateBinaryHeap(numbers []int) (heap *BinaryHeap) {
	heap = &BinaryHeap{heap: []int{0xFFFF}, len: len(numbers) + 1}
	heap.heap = append(heap.heap, numbers...)

	if heap.len <= 2 {
		return
	}

	// 构建堆O(n)
	heap.buildHeap()

	return
}

func (bh *BinaryHeap) buildHeap() {

	for i := bh.size() / 2; i >= 1; i-- {
		bh.bubbleDown(i, bh.size()-1)
	}
}

func (bh *BinaryHeap) bubbleUp(k int) {
	for ; k > 1 && bh.less(k/2, k); k = k / 2 {
		bh.swap(k, k/2)
	}
}

func (bh *BinaryHeap) bubbleDown(k, n int) {
	for 2*k <= n {
		j := 2 * k
		if j < n && bh.less(j, j+1) {
			j++
		}
		if !bh.less(k, j) {
			break
		}
		bh.swap(k, j)
		k = j
	}
}

func (bh *BinaryHeap) swap(i, j int) {
	bh.heap[i], bh.heap[j] = bh.heap[j], bh.heap[i]
}

func (bh *BinaryHeap) less(i, j int) bool {
	return bh.heap[i] < bh.heap[j]
}

func (bh *BinaryHeap) insert(n int) {
	bh.heap = append(bh.heap, n)
	bh.increase()
	bh.bubbleUp(bh.size() - 1)
}

func (bh *BinaryHeap) pop() (top int) {
	top = bh.heap[1]

	bh.swap(1, bh.size()-1)

	bh.heap = bh.heap[0 : bh.size()-1]
	bh.decrease()

	bh.bubbleDown(1, bh.size())

	return
}

func (bh *BinaryHeap) size() int {
	return bh.len
}

func (bh *BinaryHeap) decrease() {
	bh.len--
}

func (bh *BinaryHeap) increase() {
	bh.len++
}

func (bh *BinaryHeap) toString() string {
	var str []string
	for i := 1; i < bh.size(); i++ {
		str = append(str, strconv.Itoa(bh.heap[i]))
	}
	return strings.Join(str, ",")
}
