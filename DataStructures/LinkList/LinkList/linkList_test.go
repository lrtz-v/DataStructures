package linklist

import (
	"fmt"
	"testing"
)

func printLinkList(l *LinkedList) {
	node := l.Head
	for i := 0; i < l.Length; i++ {
		fmt.Println("Value: ", node.Value)
		node = node.Next
	}
}

func TestLinkList(t *testing.T) {
	t.Skip()
	// -3 -2 -1 0
	l := initLinkedList()
	l.pushBack(0)
	l.pushFront(-1)
	l.pushFront(-2)
	l.pushFront(-3)
	// printLinkList(l)
	l.insert(9, l.Length)
	// printLinkList(l)

	counter := 0
	f := func(node *Node) {
		counter += node.Value.(int)
	}

	l.forMap(f)
	fmt.Println(counter)
}
