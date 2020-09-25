package lowestcommonancestor

// TreeNode Definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	if p.Val > q.Val {
		p, q = q, p
	}

	if q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
