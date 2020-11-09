package sort

// PriorityQueue struct
type PriorityQueue struct {
	queue []int64
	count int
}

func (pq PriorityQueue) Len() int {
	return pq.count
}

func (pq PriorityQueue) Sort() {

}

func (pq PriorityQueue) Exch(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq.queue[i] < pq.queue[j]
}

func (pq PriorityQueue) Swim(k int) {}

func (pq PriorityQueue) Sink(k int) {}

func (pq PriorityQueue) Insert(v int64) {}

func (pq PriorityQueue) DelMax(v int64) int64 {
	return 0
}

func (pq PriorityQueue) DelMin(v int64) int64 {
	return 0
}
