package recoverFromPreorder

import (
	"testing"
)

var root = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 4,
		},
	},
	Right: &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 7,
		},
	},
}

func TestRecoverFromPreorder(t *testing.T) {
	root := recoverFromPreorder("1-2--3--4-5--6--7")
	t.Fatal(root)
	
}
