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
	len int
}

func CreateBinaryHeap(numbers []int) (heap *BinaryHeap) {
	heap = &BinaryHeap{heap:numbers, len:len(numbers)}

	if heap.len == 0 || heap.len == 1 {
		return
	}

	// 构建堆O(n), 满足堆中的每个节点的值必须大于等于其子树中每个节点的值
	heap.buildHeap()

	// 排序
	heap.sort()

	return
}

func (bh *BinaryHeap) buildHeap() {
	// 从非叶子节点开始
	for i := bh.len / 2; i >= 0; i-- {
		bh.heapify(bh.len, i)
	}
}

func (bh *BinaryHeap) heapify(n, i int) {
	for {
		maxPos := i
		leftIndex := i*2+1
		rightIndex := i*2+2
		// 父节点必须大于左右子节点
		if leftIndex < n && bh.less(i, leftIndex) {
			maxPos = leftIndex
		}
		if rightIndex < n && bh.less(maxPos,rightIndex) {
			maxPos = rightIndex
		}
		if maxPos == i {
			break
		}
		bh.swap(i, maxPos)
		i = maxPos
	}
}

func (bh *BinaryHeap) swap(i, j int) {
	bh.heap[i], bh.heap[j] = bh.heap[j], bh.heap[i]
}

func (bh *BinaryHeap) less(i, j int) bool {
	return bh.heap[i] < bh.heap[j]
}

func (bh *BinaryHeap) sort() {
	for i := bh.len-1; i > 0; i-- {
		bh.swap(i, 0)
		bh.heapify(i, 0)
	}
}

func (bh *BinaryHeap) toString() string {
	var str []string
	for _, v := range bh.heap {
		str = append(str, strconv.Itoa(v))
	}
	return strings.Join(str, ",")
}

func (bh *BinaryHeap) insert(n int) {
	bh.heap = append(bh.heap, n)
	bh.len++

	for i:=bh.len-1; i/2>=0 && bh.heap[i] < bh.heap[i/2]; i = i/2 {
		bh.swap(i, i/2)
	}
}

// todo delete
