package CBTInserter

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root *TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	return CBTInserter{}
}

func (this *CBTInserter) Insert(v int) int {
	return 0
}

func (this *CBTInserter) Get_root() *TreeNode {
	return nil
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */
