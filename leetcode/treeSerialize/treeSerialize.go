package treeSerialize

import (
	"encoding/json"
)

/*
请设计一个算法来实现二叉树的序列化与反序列化。
这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

示例:
你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"

说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。
*/

// TreeNode defination
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Order contains preOrder and inOrder scale
type Order struct {
	PreOrderList []int
	InOrderList []int
}

// Codec defination
type Codec struct {
	
}

// Constructor init Codec
func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	order := &Order{PreOrderList: []int{}, InOrderList: []int{}}
	order.preOrder(root)
	order.inOrder(root)
	d, _ := json.Marshal(order)
	return string(d)
}

func (c *Codec) deserialize(data string) *TreeNode {
	var order Order
	_ = json.Unmarshal([]byte(data), &order)
	return buildTree(order.PreOrderList, order.InOrderList)
}

func (o *Order) preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	o.PreOrderList = append(o.PreOrderList, root.Val)
	o.preOrder(root.Left)
	o.preOrder(root.Right)
}

func (o *Order) inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	o.inOrder(root.Left)
	o.InOrderList = append(o.InOrderList, root.Val)
	o.inOrder(root.Right)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	index := 0
	for ; index < len(inorder); index++ {
		if inorder[index] == root.Val {
			break
		}
	}

	iLeft := inorder[0:index]
	pLeft := preorder[1 : 1+len(iLeft)]
	root.Left = buildTree(pLeft, iLeft)

	iRight := inorder[index+1:]
	pRight := preorder[1+len(iLeft):]
	root.Right = buildTree(pRight, iRight)

	return root
}