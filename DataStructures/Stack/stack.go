package Stack

import "errors"

type node struct {
	Value interface{}
	Prev  *node
	Next  *node
}

// Stack . 栈 先进后出
type Stack struct {
	Length int
	Head   *node
	Tail   *node
}

// NewStack init new Stack
func NewStack() *Stack {
	return &Stack{}
}

// Push add an element to Stack
func (q *Stack) Push(value interface{}) {
	newNode := &node{Value: value}
	defer q.increase()

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
}

// Pop get an element from Stack
func (q *Stack) Pop() (interface{}, error) {
	if q.Length == 0 {
		return nil, errors.New("Empty Stack")
	}
	defer q.decrease()

	tmpNode := q.Tail
	if q.Length == 1 {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Tail = tmpNode.Prev
	}
	return tmpNode.Value, nil
}

func (q *Stack) decrease() {
	q.Length--
}

func (q *Stack) increase() {
	q.Length++
}

func (q *Stack) size() int {
	return q.Length
}

// Empty isEmpty
func (q *Stack) Empty() bool {
	return q.Length == 0
}
