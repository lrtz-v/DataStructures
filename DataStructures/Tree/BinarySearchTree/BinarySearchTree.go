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
	Value      int
	ParentNode *TreeNode
	Left       *TreeNode
	Right      *TreeNode
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
			if ParentNode.Right == nil {
				ParentNode.Right = temNode
				temNode.ParentNode = ParentNode
				return
			}
			ParentNode = ParentNode.Right
			continue
		} else {
			if ParentNode.Left == nil {
				ParentNode.Left = temNode
				temNode.ParentNode = ParentNode
				return
			}
			ParentNode = ParentNode.Left
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
			if ParentNode.Right != nil {
				ParentNode = ParentNode.Right
			}
		} else {
			if ParentNode.Left != nil {
				ParentNode = ParentNode.Left
			}
		}

		if ParentNode.Left == nil && ParentNode.Right == nil {
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

	if AimNode.Left == nil && AimNode.Right == nil {
		if AimParentNode != nil && AimParentNode.Left == AimNode {
			AimParentNode.Left = nil
		} else {
			AimParentNode.Right = nil
		}
	} else if AimNode.Left == nil || AimNode.Right == nil {
		SingleChildNode := AimNode.Left
		if AimNode.Right != nil {
			SingleChildNode = AimNode.Right
		}
		if AimParentNode != nil && AimParentNode.Left == AimNode {
			AimParentNode.Left = SingleChildNode
		} else {
			AimParentNode.Right = SingleChildNode
		}
	} else {
		// 查询左子树的最大数 / 右子树的最小值
		if AimNode.Left.Right != nil {
			subNode := AimNode.Left
			for subNode.Right != nil {
				subNode = subNode.Right
			}
			// 交换值
			AimNode.Value = subNode.Value
			// 删除subNode
			if subNode.Left == nil {
				subNode.ParentNode.Right = nil
			} else {
				subNode.ParentNode.Right = subNode.Left
			}
		} else if AimNode.Right.Left != nil {
			subNode := AimNode.Right
			for subNode.Left != nil {
				subNode = subNode.Left
			}
			// 交换值
			AimNode.Value = subNode.Value
			// 删除subNode
			if subNode.Right == nil {
				subNode.ParentNode.Left = nil
			} else {
				subNode.ParentNode.Left = subNode.Right
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
	if root.Left != nil {
		b.PreOrderTraversal(root.Left)
	}
	if root.Right != nil {
		b.PreOrderTraversal(root.Right)
	}
}

// DFS Recursion(递归)
// InOrderTraversal Traversal 中序遍历
// Left, Root, Right
func (b *BinaryTree) InOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		b.InOrderTraversal(root.Left)
	}
	fmt.Println(root.Value)
	if root.Right != nil {
		b.InOrderTraversal(root.Right)
	}
}

// DFS Recursion(递归)
// PostOrderTraversal Traversal 后序遍历
// left, right, root
func (b *BinaryTree) PostOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		b.InOrderTraversal(root.Left)
	}
	if root.Right != nil {
		b.InOrderTraversal(root.Right)
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
		if node.Left != nil {
			queue.Push(node.Left)
		}
		if node.Right != nil {
			queue.Push(node.Right)
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
		b.LevelTraversalRecursion(root.Left, level-1)
		b.LevelTraversalRecursion(root.Right, level-1)
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

	for node.Right != nil {
		node = node.Right
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

	for node.Left != nil {
		node = node.Left
		MinValue = node.Value
	}
	return MinValue
}

// GetSuccessor get node follows
func (b *BinaryTree) GetSuccessor(value int) int {
	// 返回给定值的后继者，若没有则返回-1
	AimNode := b.Search(value)
	if AimNode.Right != nil {
		return AimNode.Right.Value
	} else if AimNode.Left != nil {
		return AimNode.Left.Value
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
	return Max(b.GetHight(node.Left), b.GetHight(node.Right)) + 1
}

// Revert a BinaryTree
func (b *BinaryTree) Revert(root *TreeNode) {
	if root == nil {
		return
	}
	tmp := &TreeNode{}
	tmp = root.Left
	root.Left = root.Right
	root.Right = tmp
	if root.Left != nil {
		b.Revert(root.Left)
	}
	if root.Right != nil {
		b.Revert(root.Right)
	}
}

/**
 * Definition for a binary tree node.
 在一个 m*n 的二维字符串数组中输出二叉树，并遵守以下规则：

行数 m 应当等于给定二叉树的高度。
列数 n 应当总是奇数。
根节点的值（以字符串格式给出）应当放在可放置的第一行正中间。根节点所在的行与列会将剩余空间划分为两部分（左下部分和右下部分）。你应该将左子树输出在左下部分，右子树输出在右下部分。左下和右下部分应当有相同的大小。即使一个子树为空而另一个非空，你不需要为空的子树输出任何东西，但仍需要为另一个子树留出足够的空间。然而，如果两个子树都为空则不需要为它们留出任何空间。
每个未使用的空间应包含一个空的字符串""。
使用相同的规则输出子树。
二叉树的高度在范围 [1, 10] 中。

 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func PrintTree(root *TreeNode) [][]string {
	data := make([][]string, 10)
	level := 0
	nodeList := []*TreeNode{root}

	for level < 10 {
		eleNum := slot + 2 ^ level
	}

	return data[0:level]
}
