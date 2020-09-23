package reorderList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {

	if head == nil {
		return
	}

	nodes := []*ListNode{}
	tmp := head
	for tmp != nil {
		nodes = append(nodes, tmp)
		tmp = tmp.Next
	}

	i, j := 0, len(nodes)-1
	for i < j {

		next := nodes[i].Next
		nodes[i].Next = nodes[j]
		nodes[j].Next = next

		i++
		j--
	}
	nodes[i].Next = nil

}
