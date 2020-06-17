package treeSerialize

import (
	"testing"
)

func TestSerialize(t *testing.T) {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	
	obj := Constructor()
	data := obj.serialize(root)
	ans := obj.deserialize(data)
	t.Fatal(ans)
}
