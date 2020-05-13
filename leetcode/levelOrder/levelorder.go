package levelorder

// TreeNode defined here
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	var nodeLevelVal [][]int
	if root == nil {
		return nodeLevelVal
	}

	nodes := []*TreeNode{root}

	for len(nodes) > 0 {
		var nextLevel []*TreeNode
		var val []int

		for _, node := range nodes {
			val = append(val, node.Val)

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		nodeLevelVal = append(nodeLevelVal, val)
		nodes = nextLevel
	}

	return nodeLevelVal
}
