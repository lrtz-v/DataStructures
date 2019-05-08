package BinarySearchTree

import (
	"fmt"

	"github.com/DataStructures/DataStructures/Queue"
)

/*	Binary Tree
	a binary tree is a tree data structure in which each node has at most two children,
	which are referred to as the left child and the right child
*/

// DisplayLength display length
var DisplayLength = 132

// TreeNode --> node
type TreeNode struct {
	Value          int
	ParentNode     *TreeNode
	LeftChildNode  *TreeNode
	RightChildNode *TreeNode
}

// NewNode is a function for creating a new node
func NewNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

// BigThan n > node
func (n *TreeNode) BigThan(node *TreeNode) bool {
	return n.Value > node.Value
}

// BinaryTree define
type BinaryTree struct {
	Root *TreeNode
}

// NewBinaryTree create a new BinaryTree
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

// Insert a node to BinaryTree
func (b *BinaryTree) Insert(value int) {
	temNode := NewNode(value)
	// tree has no node
	if b.Root == nil {
		b.Root = temNode
		return
	}
	ParentNode := b.Root
	for {
		if temNode.BigThan(ParentNode) {
			if ParentNode.RightChildNode == nil {
				ParentNode.RightChildNode = temNode
				temNode.ParentNode = ParentNode
				return
			}
			ParentNode = ParentNode.RightChildNode
			continue
		} else {
			if ParentNode.LeftChildNode == nil {
				ParentNode.LeftChildNode = temNode
				temNode.ParentNode = ParentNode
				return
			}
			ParentNode = ParentNode.LeftChildNode
			continue
		}
	}

}

// Search a value from tree
func (b *BinaryTree) Search(value int) *TreeNode {
	temNode := NewNode(value)
	ParentNode := b.Root
	for {
		if value == ParentNode.Value {
			return ParentNode
		}
		if temNode.BigThan(ParentNode) {
			if ParentNode.RightChildNode != nil {
				ParentNode = ParentNode.RightChildNode
			}
		} else {
			if ParentNode.LeftChildNode != nil {
				ParentNode = ParentNode.LeftChildNode
			}
		}

		if ParentNode.LeftChildNode == nil && ParentNode.RightChildNode == nil {
			break
		}
	}
	return nil
}

// Delete node
func (b *BinaryTree) Delete(value int) {
	/*
		Case 1 : The node to delete is a leaf node
		Case 2 : The node to delete is a node with one child.
			connect the right pointer of the parent (of the deleted node) to single child.
		Case 3: The node to delete is a node with two children
			find the largest node in the left sub tree (15) or smallest node in the right sub tree (45)
			and replace the delete node with that node and then delete that node.
	*/
	AimNode := b.Search(value)
	AimParentNode := AimNode.ParentNode

	if AimNode.LeftChildNode == nil && AimNode.RightChildNode == nil {
		if AimParentNode != nil && AimParentNode.LeftChildNode == AimNode {
			AimParentNode.LeftChildNode = nil
		} else {
			AimParentNode.RightChildNode = nil
		}
	} else if AimNode.LeftChildNode == nil || AimNode.RightChildNode == nil {
		SingleChildNode := AimNode.LeftChildNode
		if AimNode.RightChildNode != nil {
			SingleChildNode = AimNode.RightChildNode
		}
		if AimParentNode != nil && AimParentNode.LeftChildNode == AimNode {
			AimParentNode.LeftChildNode = SingleChildNode
		} else {
			AimParentNode.RightChildNode = SingleChildNode
		}
	} else {
		// 查询左子树的最大数 / 右子树的最小值
		if AimNode.LeftChildNode.RightChildNode != nil {
			subNode := AimNode.LeftChildNode
			for subNode.RightChildNode != nil {
				subNode = subNode.RightChildNode
			}
			// 交换值
			AimNode.Value = subNode.Value
			// 删除subNode
			if subNode.LeftChildNode == nil {
				subNode.ParentNode.RightChildNode = nil
			} else {
				subNode.ParentNode.RightChildNode = subNode.LeftChildNode
			}
		} else if AimNode.RightChildNode.LeftChildNode != nil {
			subNode := AimNode.RightChildNode
			for subNode.LeftChildNode != nil {
				subNode = subNode.LeftChildNode
			}
			// 交换值
			AimNode.Value = subNode.Value
			// 删除subNode
			if subNode.RightChildNode == nil {
				subNode.ParentNode.LeftChildNode = nil
			} else {
				subNode.ParentNode.LeftChildNode = subNode.RightChildNode
			}
		}
	}
}

// Clear tree
func (b *BinaryTree) Clear() {
	b.Root = nil
}

// DFS Recursion(递归)
// PreOrderTraversal Traversal 前序遍历
// root, left, right
func (b *BinaryTree) PreOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Value)
	if root.LeftChildNode != nil {
		b.PreOrderTraversal(root.LeftChildNode)
	}
	if root.RightChildNode != nil {
		b.PreOrderTraversal(root.RightChildNode)
	}
}

// DFS Recursion(递归)
// InOrderTraversal Traversal 中序遍历
// Left, Root, Right
func (b *BinaryTree) InOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	if root.LeftChildNode != nil {
		b.InOrderTraversal(root.LeftChildNode)
	}
	fmt.Println(root.Value)
	if root.RightChildNode != nil {
		b.InOrderTraversal(root.RightChildNode)
	}
}

// DFS Recursion(递归)
// PostOrderTraversal Traversal 后序遍历
// left, right, root
func (b *BinaryTree) PostOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	if root.LeftChildNode != nil {
		b.InOrderTraversal(root.LeftChildNode)
	}
	if root.RightChildNode != nil {
		b.InOrderTraversal(root.RightChildNode)
	}
	fmt.Println(root.Value)
}

// LevelTraversal 广度遍历(层序遍历)
func (b *BinaryTree) LevelTraversal() {
	if b.Root == nil {
		return
	}
	queue := Queue.NewQueue()
	queue.Push(b.Root)
	for !queue.Empty() {
		val, err := queue.Pop()
		if err != nil {
			continue
		}
		node := val.(*TreeNode)
		fmt.Println(node.Value)
		if node.LeftChildNode != nil {
			queue.Push(node.LeftChildNode)
		}
		if node.RightChildNode != nil {
			queue.Push(node.RightChildNode)
		}
	}
}

// LevelTraversalRecursion 广度遍历(层序遍历) Recursion
func (b *BinaryTree) LevelTraversalRecursion(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == 1 {
		fmt.Println(root.Value)
	} else if level > 1 {
		b.LevelTraversalRecursion(root.LeftChildNode, level-1)
		b.LevelTraversalRecursion(root.RightChildNode, level-1)
	}
}

// GetMax get max node
func (b *BinaryTree) GetMax() int {
	node := b.Root
	if node == nil {
		return 0
	}
	var MaxValue int

	MaxValue = node.Value

	for node.RightChildNode != nil {
		node = node.RightChildNode
		MaxValue = node.Value
	}
	return MaxValue
}

// GetMin get min node
func (b *BinaryTree) GetMin() int {

	node := b.Root
	if node == nil {
		return 0
	}
	var MinValue int

	MinValue = node.Value

	for node.LeftChildNode != nil {
		node = node.LeftChildNode
		MinValue = node.Value
	}
	return MinValue
}

// GetSuccessor get node follows
func (b *BinaryTree) GetSuccessor(value int) int {
	// 返回给定值的后继者，若没有则返回-1
	AimNode := b.Search(value)
	if AimNode.RightChildNode != nil {
		return AimNode.RightChildNode.Value
	} else if AimNode.LeftChildNode != nil {
		return AimNode.LeftChildNode.Value
	} else {
		return -1
	}
}

// IsInTree is in tree
func (b *BinaryTree) IsInTree(value int) bool {
	AimNode := b.Search(value)
	if AimNode != nil {
		return true
	}
	return false
}

// GetDepth get the distance of node to root
func (b *BinaryTree) GetDepth(node *TreeNode) int {
	level := 1
	for node.ParentNode != nil {
		level++
		node = node.ParentNode
	}
	return level
}

// Max get max num bewteen x and y
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// GetHight get the distance of node to leaf
func (b *BinaryTree) GetHight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return Max(b.GetHight(node.LeftChildNode), b.GetHight(node.RightChildNode)) + 1
}

// Revert a BinaryTree
func (b *BinaryTree) Revert(root *TreeNode) {
	if root == nil {
		return
	}
	tmp := &TreeNode{}
	tmp = root.LeftChildNode
	root.LeftChildNode = root.RightChildNode
	root.RightChildNode = tmp
	if root.LeftChildNode != nil {
		b.Revert(root.LeftChildNode)
	}
	if root.RightChildNode != nil {
		b.Revert(root.RightChildNode)
	}
}
