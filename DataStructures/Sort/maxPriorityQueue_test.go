package sort

import (
	"testing"
)

func TestNewMaxPriorityQueue(t *testing.T) {
	mq := NewMaxPriorityQueue([]int64{})
	if len(mq.queue) != 1 || mq.count != 0 {
		t.Fatal(mq.queue)
	}
}

func TestMaxPriorityQueue(t *testing.T) {
	mq := NewMaxPriorityQueue([]int64{})
	mq.Insert(int64(1))
	mq.Insert(int64(4))
	mq.Insert(int64(2))
	mq.Insert(int64(9))
	mq.Insert(int64(10))
	t.Fatal(mq.queue)
}
