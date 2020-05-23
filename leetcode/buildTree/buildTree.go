package buildTree

import "fmt"

/*
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
*/

 // TreeNode Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
        return nil
    }

	root := &TreeNode{Val: preorder[0]}

	index := 0 
	for ; index < len(inorder); index++ {
		if inorder[index] == root.Val {
			break
		}
	}

	iLeft := inorder[0: index]
	pLeft := preorder[1: 1+len(iLeft)]
	root.Left = buildTree(pLeft, iLeft)

	iRight := inorder[index+1:]
	pRight := preorder[1+len(iLeft): ]
	root.Right = buildTree(pRight, iRight)

	return root
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Println(root.Val)
	inOrder(root.Right)
}

func afterOrder(root *TreeNode) {
	if root == nil {
		return
	}
	afterOrder(root.Left)
	afterOrder(root.Right)
	fmt.Println(root.Val)
}
