package redblacktree

import (
	"fmt"
	"reflect"
)

const (
	red   = false
	black = true
)

// RBNode of tree
type RBNode struct {
	key                 interface{}
	value               interface{}
	color               bool
	left, right, parent *RBNode
}

// RBTree define
type RBTree struct {
	root    *RBNode
	keyType reflect.Type
	valType reflect.Type
	size    int
	compare func(a, b interface{}) bool
}

func (RB *RBTree) init(key, val interface{}, compare func(a, b interface{}) bool) {
	RB.root = &RBNode{key: key, value: val, color: false}
	RB.compare = compare
	RB.keyType = reflect.TypeOf(key)
	RB.valType = reflect.TypeOf(val)
}

func (RB *RBTree) insert() {

}

/*
 * 左旋示意图：对节点x进行左旋
 *     p                       p
 *    /                       /
 *   x                       y
 *  / \                     / \
 * lx  y      ----->       x  ry
 *    / \                 / \
 *   ly ry               lx ly
 * 左旋做了三件事：
 * 1. 将y的左子节点赋给x的右子节点,并将x赋给y左子节点的父节点(y左子节点非空时)
 * 2. 将x的父节点p(非空时)赋给y的父节点，同时更新p的子节点为y(左或右)
 * 3. 将y的左子节点设为x，将x的父节点设为y
 */
func (RB *RBTree) leftRotate(x *RBNode) {
	// 1
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}

	// 2
	y.parent = x.parent
	if x.parent == nil {
		RB.root = y
	} else {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	}

	// 3
	y.left = x
	x.parent = y

}

/*
 * 右旋示意图：对节点y进行右旋
 *        p                   p
 *       /                   /
 *      y                   x
 *     / \                 / \
 *    x  ry   ----->      lx  y
 *   / \                     / \
 * lx  rx                   rx ry
 * 右旋做了三件事：
 * 1. 将x的右子节点赋给y的左子节点,并将y赋给x右子节点的父节点(x右子节点非空时)
 * 2. 将y的父节点p(非空时)赋给x的父节点，同时更新p的子节点为x(左或右)
 * 3. 将x的右子节点设为y，将y的父节点设为x
 */
func (RB *RBTree) rightRotate(y *RBNode) {
	x := y.left
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}

	x.parent = y.parent
	if y.parent == nil {
		RB.root = x
	} else {
		if y == y.parent.left {
			y.parent.left = x
		} else {
			y.parent.right = x
		}
	}

	x.right = y
	y.parent = x

}

// PreOrderTraversal travel (前序遍历)
// root, left, right
func (RB *RBTree) PreOrderTraversal(node *RBNode) {
	if node == nil {
		return
	}

	fmt.Println(node.value)
	if node.left != nil {
		RB.PreOrderTraversal(node.left)
	}
	if node.right != nil {
		RB.PreOrderTraversal(node.right)
	}
}

// InOrderTraversal Traversal 中序遍历
// Left, Root, Right
func (RB *RBTree) InOrderTraversal(node *RBNode) {
	if node == nil {
		return
	}

	if node.left != nil {
		RB.InOrderTraversal(node.left)
	}
	fmt.Println(node.value)
	if node.right != nil {
		RB.InOrderTraversal(node.right)
	}
}

// PostOrderTraversal Traversal 后序遍历
// left, right, root
func (RB *RBTree) PostOrderTraversal(node *RBNode) {
	if node == nil {
		return
	}

	if node.left != nil {
		RB.InOrderTraversal(node.left)
	}
	if node.right != nil {
		RB.InOrderTraversal(node.right)
	}
	fmt.Println(node.value)
}

func colorFlip(node *RBNode) {
	node.color = !node.color
	node.left.color = !node.left.color
	node.right.color = !node.right.color
}
