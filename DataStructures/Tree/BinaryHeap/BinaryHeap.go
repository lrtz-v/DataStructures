package BinaryHeap

import "fmt"

/*
二叉堆
构建时间复杂度：O(n)
*/

// BinaryHeap 最大二叉堆
type BinaryHeap struct {
	heap []int
}

// DownAdjust 下沉
func (b *BinaryHeap) DownAdjust(parentIndex int) {
	// 左节点
	index := 2*parentIndex + 1
	if index >= len(b.heap) {
		return
	}

	// 右节点存在
	if index+1 < len(b.heap) {
		index = b.max(index, index+1)
	}

	// 父节点大于等于任意子节点，跳出
	if b.heap[parentIndex] >= b.heap[index] {
		return
	}
	b.heap[parentIndex], b.heap[index] = b.heap[index], b.heap[parentIndex]
	b.DownAdjust(index)
}

// 上浮
func (b *BinaryHeap) upAdjust(index int) {
	if index <= 0 {
		return
	}
	parentIndex := (index - 1) / 2
	if b.heap[parentIndex] >= b.heap[index] {
		return
	}
	b.heap[parentIndex], b.heap[index] = b.heap[index], b.heap[parentIndex]
	b.upAdjust(parentIndex)
}

// Insert 插入新节点
func (b *BinaryHeap) Insert(value int) {
	b.heap = append(b.heap, value)
	b.upAdjust(len(b.heap) - 1)
}

// Pop the first element
func (b *BinaryHeap) Pop() int {
	lastIndex := len(b.heap) - 1
	firstIndex := 0
	b.heap[firstIndex], b.heap[lastIndex] = b.heap[lastIndex], b.heap[firstIndex]
	val := b.heap[lastIndex]
	b.heap = b.heap[0 : len(b.heap)-1]
	b.DownAdjust(0)
	return val
}

func (b *BinaryHeap) max(indexA, indexB int) int {
	if b.heap[indexA] >= b.heap[indexB] {
		return indexA
	}
	return indexB
}

// Print print
func (b *BinaryHeap) Print() {
	fmt.Println(b.heap)
}

// NewBinaryHeap init new binaryHeap
func NewBinaryHeap(arr []int) *BinaryHeap {
	b := &BinaryHeap{heap: arr}
	// 构建二叉堆，所有非叶子节点依次下沉
	for i := (len(b.heap) - 2) / 2; i >= 0; i-- {
		b.DownAdjust(i)
	}
	return b
}
