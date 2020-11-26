package intrusivelinklist

import (
	"fmt"
	"testing"
)

func TestIntrusiveList(t *testing.T) {
	type E struct {
		List
		data int
	}

	l := List{}
	e4 := &E{data: 4}
	e1 := &E{data: 1}
	l.PushBack(e4)
	l.PushFront(e1)
	l.InsertBefore(&E{data: 3}, e4)
	l.InsertAfter(&E{data: 2}, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("e: %+v\n", e)
		fmt.Printf("data: %d\n", e.(*E).data)
	}
}
