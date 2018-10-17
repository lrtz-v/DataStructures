package RedBlackTree

// RBNode of tree
type RBNode struct {
	Value  int
	color  bool
	left   *RBNode
	right  *RBNode
	parent *RBNode
}

// RBTree define
type RBTree struct {
	root *RBNode
	size int
}

func compareToNode(nodeA, nodeB *RBNode) bool {
	return nodeA.Value < nodeB.Value
}

func compareToValue(node *RBNode, value int) int {
	return node.Value - value
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
 * 左旋示意图：对节点y进行右旋
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

func (RB *RBTree) changeColor(nodes ...*RBNode) {
	for _, node := range nodes {
		node.color = !node.color
	}
}

// 1.每个节点不是红色就是黑色的；
// 2.根节点总是黑色的；
// 3.如果节点是红色的，则它的子节点必须是黑色的（反之不一定）；
// 4.从根节点到叶节点或空子节点的每条路径，必须包含相同数目的黑色节点（即相同的黑色高度）。
func (RB *RBTree) insertFixUp(node *RBNode) {
	for {
		parent := node.parent
		if !(node.parent != nil && node.parent.color) {
			break
		}
		gparent := parent.parent
		if parent == gparent.left {
			uncle := gparent.right
			if uncle != nil && uncle.color {
				RB.changeColor(parent, uncle, gparent)
				node = gparent
				continue
			}

			if node == parent.right {
				RB.leftRotate(parent)
				node, parent = parent, node
			}
			RB.changeColor(parent, gparent)
			RB.rightRotate(gparent)
		} else {
			uncle := gparent.left
			if uncle != nil && uncle.color {
				RB.changeColor(parent, uncle, gparent)
				node = gparent
				continue
			}

			if node == parent.left {
				RB.rightRotate(parent)
				node, parent = parent, node
			}
			RB.changeColor(parent, gparent)
			RB.leftRotate(gparent)
		}

		RB.changeColor(RB.root)
	}
}

// Insert element to RBTree
func (RB *RBTree) Insert(value int) {
	node := &RBNode{Value: value, color: true}

	var current *RBNode
	x := RB.root

	//1. 找到插入的位置
	for {
		if x == nil {
			break
		}
		current = x
		if compareToNode(node, x) {
			x = x.left
		} else {
			x = x.right
		}
	}
	node.parent = current

	if current != nil {
		if compareToNode(node, current) {
			current.left = node
		} else {
			current.right = node
		}
	} else {
		RB.root = node
	}

	RB.insertFixUp(node)
}

// Search node
func (RB *RBTree) Search(value int) *RBNode {
	node := RB.root
	for {
		if node == nil {
			break
		}
		tmp := compareToValue(node, value)
		if tmp < 0 {
			node = node.right
		} else if tmp > 0 {
			node = node.left
		} else {
			return node
		}
	}
	return node
}

// NewRBTree create a new NewRBTree with a set of elements
func NewRBTree(set []int) (*RBTree, error) {

	return nil, nil
}
