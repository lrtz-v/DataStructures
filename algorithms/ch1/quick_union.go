package ch1

type QuickUnion struct {
	ids []int
}

func (u *QuickUnion) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

// 访问数组的次数为1 加上给定触点所对应的节点的深度的两倍
func (u *QuickUnion) Find(p int) int {
	for p != u.ids[p] {
		p = u.ids[p]
	}
	return p
}

// 使用ids数组的值代表连接到的触点，形成一个节点树，通过Find操作，寻找根节点，根节点相同即相连
// 树的深度在构造过程中可能会退化
func (u *QuickUnion) Union(p, q int) {
	pRoot := u.Find(p)
	qRoot := u.Find(q)
	if pRoot == qRoot {
		return
	}
	u.ids[pRoot] = qRoot
}
