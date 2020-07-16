package buildtree

import "testing"

var root = &TreeNode{
	Val: 3,
	Left: &TreeNode{
		Val: 9,
	},
	Right: &TreeNode{
		Val: 20,
		Left: &TreeNode{
			Val: 15,
		},
		Right: &TreeNode{
			Val: 7,
		},
	},
}

var testNode = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 5,
		},
	},
	Right: &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 7,
		},
	},
}


func TestBuildTree(t *testing.T) {
	buildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7})
}
