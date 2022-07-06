package ch1

type UnionFind struct {
	ids []int
}

func (u *UnionFind) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *UnionFind) Find(p int) int {
	return u.ids[p]
}

// 使用ids数组的值代表触点，值相同代表相连，每次union，都需要把值刷一遍
func (u *UnionFind) Union(p, q int) {
	pId := u.Find(p)
	qId := u.Find(q)
	if pId == qId {
		return
	}
	for i := 0; i < len(u.ids); i++ {
		if u.ids[i] == pId {
			u.ids[i] = qId
		}
	}
}
