package ch1

type WeightedQuickUnion struct {
	ids []int
	sz  []int // 各个根节点所对应的分量的大小
}

func (u *WeightedQuickUnion) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *WeightedQuickUnion) Find(p int) int {
	for p != u.ids[p] {
		p = u.ids[p]
	}
	return p
}

// 构造的森林中的任意节点的深度最多为lgN，将两个数归并时，选择将节点数小的树连接到节点数大的树上，这样小的树的深度只增加1，平均深度更小
func (u *WeightedQuickUnion) Union(p, q int) {
	i := u.Find(p)
	j := u.Find(q)
	if i == j {
		return
	}
	if u.sz[i] < u.sz[j] {
		u.ids[i] = j
		u.sz[j] += u.sz[i]
	} else {
		u.ids[j] = i
		u.sz[i] += u.sz[j]
	}
}
