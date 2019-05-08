package BinaryHeap

import (
	"fmt"
)

// Node of tree
type Node struct {
	Value int
}

// MaxHeap define
// 父节点的下标是index，那么它的左孩子下标就是 2*index+1；它的右孩子下标就是  2*index+2
type MaxHeap struct {
	Heap []*Node
	size int
}

// get index of parent of node at index k
func parent(k int) (int, bool) {
	if k == 0 {
		return 0, false
	}
	return k / 2, true
}

// get index of left child of node at index k
func left(k int) int {
	return 2*k + 1
}

// get index of right child of node at index k
func right(k int) int {
	return 2*k + 2
}

func (heap *MaxHeap) bubbleUp(k int) {
	p, ok := parent(k)
	if !ok {
		return // k is root node
	}
	if heap.Heap[k].Value > heap.Heap[p].Value {
		heap.Heap[p], heap.Heap[k] = heap.Heap[k], heap.Heap[p]
		heap.bubbleUp(p)
	}
}

func (heap *MaxHeap) min(a, b int) int {
	if heap.Heap[a].Value > heap.Heap[b].Value {
		return b
	}
	return a
}

func (heap *MaxHeap) bubbleDown(k int) {
	min := k
	left := left(k)
	for i := 0; i < 2; i++ {
		if (left + i) <= heap.size-1 {
			if heap.Heap[left+i].Value > heap.Heap[min].Value {
				min = left + i
			}
		}
	}
	if min != k {
		heap.Heap[k], heap.Heap[min] = heap.Heap[min], heap.Heap[k]
		heap.bubbleUp(k)
	}
}

// Insert x into the heap
func (heap *MaxHeap) Insert(x int) {
	heap.Heap = append(heap.Heap, &Node{Value: x})
	heap.bubbleUp(len(heap.Heap) - 1)
	heap.size++
}

// Print elements
func (heap *MaxHeap) Print(index int) {
	for _, value := range heap.Heap {
		fmt.Println(value.Value)
	}
}

func (heap *MaxHeap) exchange(indexA, indexB int) {
	if indexA == indexB {
		return
	}
	heap.Heap[indexA].Value, heap.Heap[indexB].Value = heap.Heap[indexB].Value, heap.Heap[indexA].Value
}

// Clear elements
func (heap *MaxHeap) Clear() {
	heap.Heap = make([]*Node, 0)
	heap.size = 0
}

func compare(nodeA, nodeB *Node) bool {
	return nodeA.Value >= nodeB.Value
}

// NewMaxHeap create a new NewMaxHeap with a set of elements
func NewMaxHeap(set []int) (*MaxHeap, error) {

	var maxHeap MaxHeap
	maxHeap.Heap = make([]*Node, len(set))

	for i, v := range set {
		maxHeap.Heap[i] = &Node{Value: v}
		maxHeap.bubbleUp(i)
		maxHeap.size++
	}
	return &maxHeap, nil
}

// 返回最大值但不移除元素
func (heap *MaxHeap) getMax() int {
	return heap.Heap[0].Value
}

// 返回存储的元素数量
func (heap *MaxHeap) getSize() int {
	return heap.size
}

// 若堆为空则返回 true
func (heap *MaxHeap) isEmpty() bool {
	return heap.size == 0
}

// 返回最大值并移除
func (heap *MaxHeap) extractMax() int {
	maxValue := heap.Heap[0].Value
	heap.exchange(0, heap.size-1)
	heap.Heap = heap.Heap[0 : heap.size-1]
	heap.size--

	for i := 0; i < heap.size/2; i++ {
		heap.bubbleDown(i)
	}

	return maxValue
}

// 删除指定索引的元素
func (heap *MaxHeap) remove(index int) int {
	maxValue := heap.Heap[index].Value
	heap.exchange(index, heap.size-1)
	heap.Heap = heap.Heap[0 : heap.size-1]
	heap.size--

	for i := index; i < heap.size/2; i++ {
		heap.bubbleDown(i)
	}

	return maxValue
}
