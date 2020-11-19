package sort

import (
	"testing"
)

func TestNewMaxPriorityQueue(t *testing.T) {
	mq := NewMaxPriorityQueue([]int64{})
	if len(mq.queue) != 1 || mq.count != 0 {
		t.Fatal("[*] Create empty MaxPriorityQueue failed")
	}
}

func TestInsert(t *testing.T) {
	mq := NewMaxPriorityQueue([]int64{})

	mq.Insert(int64(1))
	if len(mq.queue) != 2 || mq.count != 1 {
		t.Fatalf("[*] Insert failed, length: %d, count: %d\n", len(mq.queue), mq.count)
	}

	mq.Insert(int64(4))
	if len(mq.queue) != 3 || mq.count != 2 {
		t.Fatalf("[*] Insert failed, length: %d, count: %d\n", len(mq.queue), mq.count)
	}

	mq.Insert(int64(2))
	if len(mq.queue) != 4 || mq.count != 3 {
		t.Fatalf("[*] Insert failed, length: %d, count: %d\n", len(mq.queue), mq.count)
	}

	mq.Insert(int64(9))
	if len(mq.queue) != 5 || mq.count != 4 {
		t.Fatalf("[*] Insert failed, length: %d, count: %d\n", len(mq.queue), mq.count)
	}

	mq.Insert(int64(10))
	if len(mq.queue) != 6 || mq.count != 5 {
		t.Fatalf("[*] Insert failed, length: %d, count: %d\n", len(mq.queue), mq.count)
	}
}

func TestGetMax(t *testing.T) {
	mq := NewMaxPriorityQueue([]int64{})
	max, err := mq.GetMax()
	if err == nil {
		t.Fatal("Should be empty")
	}
	mq.Insert(int64(1))
	max, _ = mq.GetMax()
	if max != 1 {
		t.Fatalf("[*] GetMax failed, Expected 1, Got: %d\n", max)
	}
	mq.Insert(int64(4))
	max, _ = mq.GetMax()
	if max != 4 {
		t.Fatalf("[*] GetMax failed, Expected 4, Got: %d\n", max)
	}
	mq.Insert(int64(2))
	max, _ = mq.GetMax()
	if max != 4 {
		t.Fatalf("[*] GetMax failed, Expected 4, Got: %d\n", max)
	}
	mq.Insert(int64(9))
	max, _ = mq.GetMax()
	if max != 9 {
		t.Fatalf("[*] GetMax failed, Expected 9, Got: %d\n", max)
	}
	mq.Insert(int64(10))
	max, _ = mq.GetMax()
	if max != 10 {
		t.Fatalf("[*] GetMax failed, Expected 10, Got: %d\n", max)
	}
}

func TestDelMax(t *testing.T) {
	mq := NewMaxPriorityQueue([]int64{})
	var max int64
	v, err := mq.DelMax()
	if err == nil {
		t.Fatal("Should be empty")
	}
	mq.Insert(int64(1))
	mq.Insert(int64(4))
	mq.Insert(int64(2))
	mq.Insert(int64(9))
	mq.Insert(int64(10))

	v, _ = mq.DelMax()
	if v != 10 {
		t.Fatalf("Delete Max failed. Expected 10, Got %d\n", v)
	}
	max, _ = mq.GetMax()
	if max != 9 {
		t.Fatalf("[*] GetMax failed, Expected 9, Got: %d\n", max)
	}

	v, _ = mq.DelMax()
	if v != 9 {
		t.Fatalf("Delete Max failed. Expected 9, Got %d\n", v)
	}
	max, _ = mq.GetMax()
	if max != 4 {
		t.Fatalf("[*] GetMax failed, Expected 4, Got: %d\n", max)
	}

	v, _ = mq.DelMax()
	if v != 4 {
		t.Fatalf("Delete Max failed. Expected 4, Got %d\n", v)
	}
	max, _ = mq.GetMax()
	if max != 2 {
		t.Fatalf("[*] GetMax failed, Expected 2, Got: %d\n", max)
	}

	v, _ = mq.DelMax()
	if v != 2 {
		t.Fatalf("Delete Max failed. Expected 2, Got %d\n", v)
	}
	max, _ = mq.GetMax()
	if max != 1 {
		t.Fatalf("[*] GetMax failed, Expected 1, Got: %d\n", max)
	}

	v, _ = mq.DelMax()
	if v != 1 {
		t.Fatalf("Delete Max failed. Expected 1, Got %d\n", v)
	}
	max, err = mq.GetMax()
	if err == nil {
		t.Fatal("Should be empty")
	}

	v, err = mq.DelMax()
	if err == nil {
		t.Fatal("Should be empty")
	}

	if len(mq.queue) != 1 || mq.count != 0 {
		t.Fatal("Should be empty")
	}
}
