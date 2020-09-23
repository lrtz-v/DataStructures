package flatten

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	tmp := root

	for tmp != nil {
		if tmp.Child != nil {
			next := tmp.Next
			child := tmp.Child
			start, end := flat(child)

			end.Next = next
			if next != nil {
				next.Prev = end
			}
			tmp.Child = nil
			tmp.Next = start
			start.Prev = tmp
		}
		tmp = tmp.Next
	}

	return root
}

func flat(root *Node) (*Node, *Node) {
	tmp := root
	for tmp.Next != nil {
		tmp = tmp.Next
	}

	return root, tmp
}
