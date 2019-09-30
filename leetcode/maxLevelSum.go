package leetcode

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {

	parentNodeList := []*TreeNode{root}
	childNodeList := []*TreeNode{}

	maxLevel := 1
	max := root.Val
	level := 1

	for {

		if len(parentNodeList) == 0 {
			break
		}
		tmp := 0
		for _, node := range parentNodeList {
			tmp += node.Val
			if node.Left != nil {
				childNodeList = append(childNodeList, node.Left)
			}
			if node.Right != nil {
				childNodeList = append(childNodeList, node.Right)
			}
		}
		parentNodeList, childNodeList = childNodeList, []*TreeNode{}
		if tmp > max {
			max = tmp
			maxLevel = level
		}
		level++

	}

	return maxLevel
}

func maxLevelSum2(root *TreeNode) int {

	if root == nil {
		return 0
	}

	nodeList := []*TreeNode{root}

	maxLevel := 1
	max := root.Val
	level := 1

	for {
		tmpList := []*TreeNode{}

		if len(nodeList) == 0 {
			break
		}
		tmp := 0
		for _, node := range nodeList {
			tmp += node.Val
			if node.Left != nil {
				tmpList = append(tmpList, node.Left)
			}
			if node.Right != nil {
				tmpList = append(tmpList, node.Right)
			}
		}

		nodeList = tmpList
		if tmp > max {
			max = tmp
			maxLevel = level
		}
		level++
	}

	return maxLevel
}
