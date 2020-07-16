package buildtree

/*
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
*/

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    indexMap := make(map[int]int)
    for i := 0; i < len(inorder); i++ {
        indexMap[inorder[i]] = i
    }
    return build(preorder, inorder, 0, len(preorder) - 1, 0, len(inorder) - 1, indexMap)
}

func build(preorder, inorder []int, preorderStart, preorderEnd, inorderStart, inorderEnd int, indexMap map[int]int) *TreeNode {
    if preorderStart > preorderEnd {
        return nil
    }

    root := &TreeNode{Val: preorder[preorderStart]}
    if preorderStart == preorderEnd {
        return root
    }
    rootIndex := indexMap[preorder[preorderStart]]
    leftNodes := rootIndex - inorderStart
    rightNodes := inorderEnd - rootIndex
    root.Left = build(preorder, inorder, preorderStart + 1, preorderStart + leftNodes, inorderStart, rootIndex - 1, indexMap)
    root.Right = build(preorder, inorder, preorderEnd - rightNodes + 1, preorderEnd, rootIndex + 1, inorderEnd, indexMap)

    return root
}
