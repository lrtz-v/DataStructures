package skiplist

import "math/rand"

const ZSKIPLIST_MAXLEVEL = 32 // 最大层数为32
const ZSKIPLIST_P = 0.25      // 被插入到高层的概率为1/4

type zskiplistLevel struct {
	forward *zskiplistNode // 指向链表下一个节点的指针，即前向指针
	span    int            // 表示这个前向指针跳跃过了多少个节点
}

type zskiplistNode struct {
	obj      int               // 存放该节点的数据
	score    float32           // 数据对应的分数值，zset通过分数对数据升序排列
	backward *zskiplistNode    // 指向链表上一个节点的指针，即后向指针
	level    []*zskiplistLevel // 表示跳表中的一层
}

type zskiplist struct {
	header *zskiplistNode // 头指针
	tail   *zskiplistNode // 尾指针
	length int            // 跳表的长度，不包括头指针
	level  int            // 跳表的层数
}

// 跳表初始化
func zslInit() *zskiplist {
	zsl := &zskiplist{level: 1, length: 0}
	zsl.header = zslNodeInit(ZSKIPLIST_MAXLEVEL, 0, 0)
	for j := 0; j < ZSKIPLIST_MAXLEVEL; j++ {
		zsl.header.level[j].forward = nil
		zsl.header.level[j].span = 0
	}
	zsl.header.backward = nil
	zsl.tail = nil
	return zsl
}

// 节点初始化
func zslNodeInit(level int, score float32, val int) *zskiplistNode {
	l := make([]*zskiplistLevel, level)
	for i := 0; i < level; i++ {
		l[i] = &zskiplistLevel{}
	}
	return &zskiplistNode{obj: val, score: score, backward: nil, level: l}
}

func (zsl *zskiplist) insert(score float32, val int) *zskiplistNode {
	update := make([]*zskiplistNode, ZSKIPLIST_MAXLEVEL)
	rank := make([]int, ZSKIPLIST_MAXLEVEL) // 用于计算span

	x := zsl.header

	for i := zsl.level; i >= 0; i-- {
		if i == zsl.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}
		for x.level[i].forward != nil && x.level[i].forward.score <= score {
			rank[i] += x.level[i].span
			x = x.level[i].forward
		}
		// 记录更新路径
		update[i] = x
	}

	// 随机出要插入的节点的层数。
	level := zsl.zslRandomLevel()
	if level > zsl.level {
		for i := zsl.level; i < level; i++ {
			rank[i] = 0
			update[i] = zsl.header
			update[i].level[i].span = zsl.length
		}
		zsl.level = level
	}

	x = zslNodeInit(level, score, val)
	for i := 0; i < level; i++ {
		x.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = x

		x.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}

	for i := level; i < zsl.level; i++ {
		update[i].level[i].span++
	}

	if update[0] == zsl.header {
		x.backward = nil
	} else {
		x.backward = update[0]
	}
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x
	} else {
		zsl.tail = x
	}
	zsl.length++
	return x
}

// 随机出要插入的节点的层数。
func (zsl *zskiplist) zslRandomLevel() int {
	level := 1
	p := ZSKIPLIST_P * 0xFFFF
	for {
		randInt := rand.Int63()&0xFFFF
		if randInt < int64(p) {
			level += 1
		} else {
			break
		}
	}
	if level < ZSKIPLIST_MAXLEVEL {
		return level
	} else {
		return ZSKIPLIST_MAXLEVEL
	}
}
