package lowestcommonancestor

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	root := TreeNode{
		Val: 3, 
		Left: &TreeNode{
			Val: 5, 
			Left: &TreeNode{
				Val: 6, 
				Left: &TreeNode{},
				Right: &TreeNode{},
			},
			Right: &TreeNode{
				Val: 2, 
				Left: &TreeNode{
					Val: 7, 
				},
				Right: &TreeNode{
					Val: 4, 
				},
			},
		},
		Right: &TreeNode{
			Val: 1, 
			Left: &TreeNode{
				Val: 0, 
				Left: &TreeNode{},
				Right: &TreeNode{},},
			Right: &TreeNode{
				Val: 8, 
				Left: &TreeNode{},
				Right: &TreeNode{},
			},
		},
	}
	
	n := lowestCommonAncestor(&root, root.Left, root.Left.Right.Right)
	if n.Val != 5 {
		t.Fatalf("Wrong Ancestor! Expected 5, Got %d", n.Val)
	}
}

func TestLowestCommonAncestor2(t *testing.T) {
	root := TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val: 0, 
			Left: &TreeNode{
				Val: -2, 
				Left: &TreeNode{
					Val: 8,
				},
			},
			Right: &TreeNode{
				Val: 4, 
			},
		},
		Right: &TreeNode{
			Val: 3, 
		},
	}
	
	n := lowestCommonAncestor(&root, root.Left.Left.Left, root.Left.Right)
	if n.Val != 0 {
		t.Fatalf("Wrong Ancestor! Expected 0, Got %d", n.Val)
	}
}
