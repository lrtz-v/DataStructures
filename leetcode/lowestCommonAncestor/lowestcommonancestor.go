package lowestcommonancestor

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

// TreeNode Definition
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildParentMap(n, p *TreeNode, m map[int]*TreeNode) {
	if n == nil {
		return
	}
	m[n.Val] = p
	buildParentMap(n.Left, n, m)
	buildParentMap(n.Right, n, m)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	parentMap := make(map[int]*TreeNode)
	buildParentMap(root, nil, parentMap)
	
	pathMap := make(map[int]*TreeNode)
	pathMap[p.Val] = p
	
	for pParent := parentMap[p.Val]; pParent != nil; {
		pathMap[pParent.Val] = pParent
		pParent = parentMap[pParent.Val]
	}

	if _, ok := pathMap[q.Val]; ok {
		return q
	}
	pathMap[q.Val] = q

	for qParent := parentMap[q.Val]; qParent != nil; {		
		if _, ok := pathMap[qParent.Val]; ok {
			return qParent
		}
		qParent = parentMap[qParent.Val]
	}

	return root
}
