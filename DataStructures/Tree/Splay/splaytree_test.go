package splay

import (
	"testing"
)

func TestSplayTreeSearch(t *testing.T) {
	t.Skip()
	tree := NewSplayTree()
	tree.Root = NewSplayTreeNode(100)

	tree.Root.LeftNode = NewSplayTreeNode(50)
	tree.Root.RightNode = NewSplayTreeNode(200)
	tree.Root.LeftNode.LeftNode = NewSplayTreeNode(40)
	tree.Root.LeftNode.LeftNode.LeftNode = NewSplayTreeNode(30)
	tree.Root.LeftNode.LeftNode.LeftNode.LeftNode = NewSplayTreeNode(20)

	root := tree.SplaySearch(20)
	preOrder(root)
}

func TestSplayTreeInsert(t *testing.T) {
	t.Skip()
	tree := NewSplayTree()
	tree.Root = NewSplayTreeNode(100)

	tree.Root.LeftNode = NewSplayTreeNode(50)
	tree.Root.RightNode = NewSplayTreeNode(200)
	tree.Root.LeftNode.LeftNode = NewSplayTreeNode(40)
	tree.Root.LeftNode.LeftNode.LeftNode = NewSplayTreeNode(30)
	tree.Root.LeftNode.LeftNode.LeftNode.LeftNode = NewSplayTreeNode(20)

	tree.SplayInsert(25)
	preOrder(tree.Root)
}

func TestSplayTreeDelete(t *testing.T) {
	// t.Skip()
	tree := NewSplayTree()
	tree.Root = NewSplayTreeNode(6)

	tree.Root.LeftNode = NewSplayTreeNode(1)
	tree.Root.RightNode = NewSplayTreeNode(9)
	tree.Root.LeftNode.RightNode = NewSplayTreeNode(4)
	tree.Root.LeftNode.RightNode.LeftNode = NewSplayTreeNode(2)
	tree.Root.RightNode.LeftNode = NewSplayTreeNode(7)

	tree.SplayDelete(4)
	preOrder(tree.Root)
}
