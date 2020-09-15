package generateTrees


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {

	if n == 0 {
		return []*TreeNode{}
	}

	return doGenerate(1, n)
}

func doGenerate(start, end int) []*TreeNode {

	res := []*TreeNode{}

	if end < start {
		res = append(res, nil)
	}

	for i := start; i <= end; i++ {
		leftRes := doGenerate(start, i-1)
		rightRes := doGenerate(i+1, end)

		for _, l := range leftRes {
			for _, r := range rightRes {
				node := &TreeNode{Val: i}
				node.Left = l
				node.Right = r
				res = append(res, node)
			}
		}

	}

	return res
}
