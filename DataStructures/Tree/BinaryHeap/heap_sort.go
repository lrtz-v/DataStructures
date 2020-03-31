package binaryheap

func (bh *BinaryHeap) sort() {

	for n := bh.size(); n > 1; n-- {
		bh.swap(1, n-1)
		bh.bubbleDown(1, n-2)
	}
}
