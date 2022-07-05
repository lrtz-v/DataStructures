package queue

type Queue interface {
	Enqueue(e any)
	Dequeue()
	Empty() bool
	Size() int
}

// 下压栈，先进后出
type Stack interface {
	Push(e any)
	pop() any
	Empty() bool
	Size() int
}

// 先进先出
type FiFoQueue interface {
	Queue
}
