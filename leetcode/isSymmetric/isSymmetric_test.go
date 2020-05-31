package isSymmetric

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
		Val: 2,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 3,
		},
	},
}

func TestIsSymmetric(t *testing.T) {
	if !isSymmetric(root) {
		t.Fatalf("[*] isSymmetric Error!")
	}
}

func TestIsSymmetric2(t *testing.T) {
	case1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	if isSymmetric(case1) {
		t.Fatalf("[*] isSymmetric Error!")
	}
}
