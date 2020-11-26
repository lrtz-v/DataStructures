package splay

import (
	"fmt"
)

// SplayNode node
type SplayNode struct {
	Value               int
	LeftNode, RightNode *SplayNode
}

// NewSplayTreeNode create node
func NewSplayTreeNode(value int) *SplayNode {
	return &SplayNode{
		Value: value,
	}
}

// SplayTree tree
type SplayTree struct {
	Root *SplayNode
}

// NewSplayTree create tree
func NewSplayTree() *SplayTree {
	return &SplayTree{
		Root: nil,
	}
}

// RightRotate A utility function to right rotate subtree rooted with y
// See the diagram given above.
func RightRotate(x *SplayNode) *SplayNode {
	y := x.LeftNode
	x.LeftNode = y.RightNode
	y.RightNode = x
	return y
}

// LeftRotate A utility function to left rotate subtree rooted with x
// See the diagram given above.
func LeftRotate(x *SplayNode) *SplayNode {
	y := x.RightNode
	x.RightNode = y.LeftNode
	y.LeftNode = x
	return y
}

// preOrder print
func preOrder(root *SplayNode) {
	if root != nil {
		fmt.Println(root.Value)
		preOrder(root.LeftNode)
		preOrder(root.RightNode)
	}
}

// splay tree
func splay(root *SplayNode, value int) *SplayNode {
	if root == nil {
		return nil
	} else if root.Value == value {
		return root
	}

	if root.Value > value {
		// Key is not in tree
		if root.LeftNode == nil {
			return root
		}

		// Zig-Zig (Left Left)
		if root.LeftNode.Value > value {
			root.LeftNode.LeftNode = splay(root.LeftNode.LeftNode, value)
			root = RightRotate(root)
		} else {
			// Zig-Zag (Left Right)
			root.LeftNode.RightNode = splay(root.LeftNode.RightNode, value)
			if root.LeftNode.RightNode != nil {
				root.LeftNode = LeftRotate(root.LeftNode)
			}
		}

		if root.LeftNode == nil {
			return root
		}
		return RightRotate(root)

	}
	// ============================================================================
	// Key is not in tree
	if root.RightNode == nil {
		return root
	}

	// Zag-Zag (Right Right)
	if root.RightNode.Value < value {
		root.RightNode.RightNode = splay(root.RightNode.RightNode, value)
		root = LeftRotate(root)
	} else {
		// Zag-Zig (Right Left)
		root.RightNode.LeftNode = splay(root.RightNode.LeftNode, value)
		if root.RightNode.LeftNode != nil {
			root.RightNode = RightRotate(root.RightNode)
		}
	}

	if root.RightNode == nil {
		return root
	}
	return LeftRotate(root)
}

// SplaySearch search
func (s *SplayTree) SplaySearch(value int) *SplayNode {
	return splay(s.Root, value)
}

// SplayInsert insert
func (s *SplayTree) SplayInsert(value int) {

	if s.Root == nil {
		s.Root = NewSplayTreeNode(value)
		return
	}

	root := splay(s.Root, value)

	if root.Value == value {
		return
	}

	node := NewSplayTreeNode(value)
	if node.Value > root.Value {
		node.LeftNode = root
		node.RightNode = root.RightNode
		root.RightNode = nil
	} else {
		node.RightNode = root
		node.LeftNode = root.LeftNode
		root.LeftNode = nil
	}
	s.Root = node
}

// SplayDelete delete
func (s *SplayTree) SplayDelete(value int) {
	if s.Root == nil {
		return
	}
	root := splay(s.Root, value)
	if value != root.Value {
		return
	}

	if root.LeftNode == nil {
		root = root.RightNode
	} else {
		temp := root
		root = splay(root.LeftNode, value)
		root.RightNode = temp.RightNode
	}

}
