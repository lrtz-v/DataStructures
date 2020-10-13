package maxTreeProduct

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var mod = 10 ^ 9 + 7
var maxMul int64

func maxProduct(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sumMap := make(map[*TreeNode]int)

	sum := nodeSum(root, sumMap)
	splitTree(root, sum, sumMap)
	return int(maxMul % 1000000007)
}

func nodeSum(root *TreeNode, sumMap map[*TreeNode]int) int {
	if root == nil {
		return 0
	}

	l := nodeSum(root.Left, sumMap)
	r := nodeSum(root.Right, sumMap)
	tmp := l + r + root.Val

	sumMap[root] = tmp
	return tmp
}

func splitTree(root *TreeNode, sum int, sumMap map[*TreeNode]int) {
	if root == nil {
		return
	} else if root.Left == nil && root.Right == nil {
		return
	}

	var tmp int
	if root.Left != nil {
		sl1 := sumMap[root.Left]
		sr1 := sum - sl1
		tmp = sr1 * sl1
	}
	if root.Right != nil {
		sr2 := sumMap[root.Right]
		sl2 := sum - sr2
		tmp = max(tmp, sr2*sl2)
	}

	tmpi64 := int64(tmp)
	if tmpi64 > maxMul {
		maxMul = tmpi64
	}
	splitTree(root.Left, sum, sumMap)
	splitTree(root.Right, sum, sumMap)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
