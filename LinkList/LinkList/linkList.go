package main

/*
	双向链表✅

	链表（Linked Lists）
		- [ ] size() —— 返回链表中数据元素的个数
        - [ ] empty() —— 若链表为空则返回一个布尔值 true
		- [ ] value_at(index) —— 返回第 n 个元素的值（从0开始计算）
		- [ ] push_front(value) —— 添加元素到链表的首部
		- [ ] pop_front() —— 删除首部元素并返回其值
		- [ ] push_back(value) —— 添加元素到链表的尾部
		- [ ] pop_back() —— 删除尾部元素并返回其值
		- [ ] front() —— 返回首部元素的值
		- [ ] back() —— 返回尾部元素的值
		- [ ] insert(index, value) —— 插入值到指定的索引，并把当前索引的元素指向到新的元素
		- [ ] erase(index) —— 删除指定索引的节点
		- [ ] value_n_from_end(n) —— 返回倒数第 n 个节点的值
		- [ ] reverse() —— 逆序链表
        - [ ] remove_value(value) —— 删除链表中指定值的第一个元素
*/

// Node struct 双向节点
type Node struct {
	Value interface{}
	Prev  *Node
	Next  *Node
}

// LinkedList struct 链表
type LinkedList struct {
	Length int
	Head   *Node
	Tail   *Node
}

func initNode(value interface{}, Prev *Node, Next *Node) *Node {
	return &Node{Value: value, Prev: Prev, Next: Next}
}

func initLinkedList() *LinkedList {
	return &LinkedList{Length: 0, Head: nil, Tail: nil}
}

func (l *LinkedList) empty() bool {
	return l.Length == 0
}

func (l *LinkedList) size() int {
	return l.Length
}

func (l *LinkedList) valueAt(index int) interface{} {
	nextNode := l.Head
	for i := 0; i < index; i++ {
		if nextNode.Next == nil {
			break
		}
		nextNode = nextNode.Next
	}
	return nextNode.Value
}

func (l *LinkedList) pushFront(value interface{}) {
	newNode := &Node{Value: value, Prev: nil, Next: l.Head}
	if l.Length == 0 {
		l.Head = newNode
		l.Tail = newNode
	} else if l.Length == 1 {
		l.Head.Prev = newNode
		l.Tail = l.Head
		l.Head = newNode
	} else {
		l.Head.Prev = newNode
		l.Head = newNode
	}
	l.Length++
}

func (l *LinkedList) pushBack(value interface{}) {
	newNode := &Node{Value: value, Prev: l.Tail, Next: nil}
	if l.Length == 0 {
		l.Head = newNode
		l.Tail = newNode
	} else if l.Length == 1 {
		l.Head.Next = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Length++
}

func (l *LinkedList) front() interface{} {
	return l.Head.Value
}

func (l *LinkedList) back() interface{} {
	return l.Tail.Value
}

func (l *LinkedList) insert(value interface{}, index int) {
	// 在指定位置插入节点
	node := l.Head
	for i := 0; i < index; i++ {
		if node.Next != nil {
			node = node.Next
		} else {
			break
		}
	}

	if node == l.Head {
		newNode := &Node{Value: value, Prev: nil, Next: node}
		node.Prev = newNode
		l.Head = newNode
	} else if node == l.Tail {
		newNode := &Node{Value: value, Prev: node.Prev, Next: node}
		node.Prev.Next = newNode
		node.Prev = newNode
	} else {
		newNode := &Node{Value: value, Prev: node.Prev, Next: node}
		node.Prev.Next = newNode
		node.Prev = newNode
	}

	l.Length++
}

func (l *LinkedList) erase(index int) {
	// 移除指定位置的节点
	node := l.Head
	for i := 0; i < index; i++ {
		if node.Next != nil {
			node = node.Next
		} else {
			break
		}
	}
	if node == l.Head {
		nextNode := l.Head.Next
		nextNode.Prev = nil
		l.Head = nextNode
	} else if node == l.Tail {
		PrevNode := l.Tail.Prev
		PrevNode.Next = nil
		l.Tail = PrevNode
	} else {
		PreNode := node.Prev
		NextNode := node.Next
		PreNode.Next = NextNode
		NextNode.Prev = PreNode
	}
	l.Length--
}

func (l *LinkedList) valueNFromEnd(index int) interface{} {
	return l.valueAt(l.Length - index)
}

func (l *LinkedList) reverse() {

	node := l.Head
	for i := 0; i < l.Length; i++ {
		next := node.Next
		node.Prev, node.Next = node.Next, node.Prev
		if next != nil {
			node = next
		}
	}

	l.Head, l.Tail = l.Tail, l.Head
}

func (l *LinkedList) removeValue(value interface{}) {

	node := l.Head
	for i := 0; i < l.Length; i++ {
		if node.Value == value {
			l.erase(i)
			break
		}
		if node.Next != nil {
			node = node.Next
		} else {
			break
		}
	}
}

func (l *LinkedList) clear() {
	l.Length = 0
	l.Head = nil
	l.Tail = nil
}

func (l *LinkedList) concat(k *LinkedList) {
	l.Tail.Next, k.Head.Prev = k.Head, l.Tail
	l.Tail = k.Tail
	l.Length += k.Length
}

func (l *LinkedList) forMap(f func(node *Node)) {
	for node := l.Head; node != nil; node = node.Next {
		f(node)
	}
}
