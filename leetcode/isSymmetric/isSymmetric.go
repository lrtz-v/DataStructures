package isSymmetric

import "fmt"

/*
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3


But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3


Follow up: Solve it both recursively and iteratively.
*/

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	pre := []int{}
	after := []int{}

	preOrder(root, &pre)
	afterOrder(root, &after)

	fmt.Println(pre)
	fmt.Println(after)
	l := len(pre)
	for i := 0; i < len(pre); i++ {
		if pre[i] != after[l-i-1] {
			return false
		}
	}
	return true
}

func preOrder(root *TreeNode, pre *[]int) {
	if root == nil {
		*pre = append(*pre, -1)
		return
	}
	*pre = append(*pre, root.Val)
	preOrder(root.Left, pre)
	preOrder(root.Right, pre)
}

func afterOrder(root *TreeNode, after *[]int) {
	if root == nil {
		*after = append(*after, -1)
		return
	}
	afterOrder(root.Left, after)
	afterOrder(root.Right, after)
	*after = append(*after, root.Val)
}

func inOrder(root *TreeNode, in *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, in)
	*in = append(*in, root.Val)
	inOrder(root.Right, in)
}

func isSymmetric2(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
