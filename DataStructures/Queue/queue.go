package Queue

import "errors"

type node struct {
	Value interface{}
	Prev  *node
	Next  *node
}

// Queue .
type Queue struct {
	Length int
	Head   *node
	Tail   *node
}

// NewQueue init new queue
func NewQueue() *Queue {
	return &Queue{Length: 0}
}

// Push add an element to queue
func (q *Queue) Push(value interface{}) {
	newNode := &node{Value: value}
	if q.Length == 0 {
		q.Head = newNode
		q.Tail = newNode
	} else if q.Length == 1 {
		newNode.Prev = q.Head
		q.Head.Next = newNode
		q.Tail = newNode
	} else {
		newNode.Prev = q.Tail
		q.Tail.Next = newNode
		q.Tail = newNode
	}
	q.Length++
}

// Pop get an element from queue
func (q *Queue) Pop() (interface{}, error) {
	if q.Length == 0 {
		return nil, errors.New("Empty Queue")
	}
	defer q.decrease()

	tmpNode := q.Head
	if q.Length == 1 {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Head = tmpNode.Next
	}
	return tmpNode.Value, nil
}

func (q *Queue) decrease() {
	q.Length--
}

func (q *Queue) size() int {
	return q.Length
}

// Empty isEmpty
func (q *Queue) Empty() bool {
	return q.Length == 0
}
