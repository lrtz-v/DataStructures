package binaryheap

func (bh *BinaryHeap) sort() {

	for n := bh.size(); n > 1; n-- {
		bh.swap(1, n-1)
		bh.bubbleDown(1, n-2)
	}
}

func (bh *BinaryHeap) sink(k, n int) {
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
