package skiplist

import (
	"math/rand"
	"strings"
)

const ZSKIPLIST_MAXLEVEL = 32 // 最大层数为32
const ZSKIPLIST_P = 0.25      // 被插入到高层的概率为1/4

type zskiplistLevel struct {
	forward *zskiplistNode // 指向链表下一个节点的指针，即前向指针
	span    int            // 表示这个前向指针跳跃过了多少个节点
}

type zskiplistNode struct {
	key      string            // key
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
	zsl.header = zslNodeInit(ZSKIPLIST_MAXLEVEL, 0, 0, "")
	for j := 0; j < ZSKIPLIST_MAXLEVEL; j++ {
		zsl.header.level[j].forward = nil
		zsl.header.level[j].span = 0
	}
	zsl.header.backward = nil
	zsl.tail = nil
	return zsl
}

// 节点初始化
func zslNodeInit(level int, score float32, val int, key string) *zskiplistNode {
	l := make([]*zskiplistLevel, level)
	for i := 0; i < level; i++ {
		l[i] = &zskiplistLevel{}
	}
	return &zskiplistNode{obj: val, score: score, backward: nil, level: l, key: key}
}

func (zsl *zskiplist) zslInsert(score float32, val int, key string) *zskiplistNode {
	update := make([]*zskiplistNode, ZSKIPLIST_MAXLEVEL)
	rank := make([]int, ZSKIPLIST_MAXLEVEL) // 用于计算span

	x := zsl.header

	for i := zsl.level; i >= 0; i-- {
		if i == zsl.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score && strings.Compare(x.level[i].forward.key, key) < 0)) {
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

	x = zslNodeInit(level, score, val, key)
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
	for rand.Int63()&0xFFFF < int64(p) {
		level += 1
	}
	if level < ZSKIPLIST_MAXLEVEL {
		return level
	} else {
		return ZSKIPLIST_MAXLEVEL
	}
}

func zslFree(zsl *zskiplist) {
	node := zsl.header.level[0].forward

	zsl.header = nil
	for node != nil {
		next := node.level[0].forward
		node = nil
		node = next
	}
	zsl = nil
}

func (zsl *zskiplist) zslDeleteNode(x *zskiplistNode, update []*zskiplistNode) {
	for i := 0; i < zsl.level; i++ {
		if update[i].level[i].forward == x {
			update[i].level[i].span += x.level[i].span - 1
			update[i].level[i].forward = x.level[i].forward
		} else {
			update[i].level[i].span -= 1
		}
	}
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x.backward
	} else {
		zsl.tail = x.backward
	}
	for zsl.level > 1 && zsl.header.level[zsl.level-1].forward == nil {
		zsl.level--
	}
	zsl.length--
}

func (zsl *zskiplist) zslDelete(score float32, key string) int {
	update := make([]*zskiplistNode, ZSKIPLIST_MAXLEVEL)

	x := zsl.header
	for i := zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score && strings.Compare(x.level[i].forward.key, key) < 0)) {
			x = x.level[i].forward
		}
		update[i] = x
	}

	x = x.level[0].forward
	if x == nil {
		return 0
	}
	for x != nil && score == x.score && strings.Compare(x.key, key) == 0 {
		zsl.zslDeleteNode(x, update)
		x = x.level[0].forward
	}
	return 1
}

func (zsl *zskiplist) zslGetRank(score float32, key string) int {
	rank := 0
	x := zsl.header
	for i := zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score && strings.Compare(x.level[i].forward.key, key) <= 0)){
			rank += x.level[i].span
			x = x.level[i].forward
		}

		/* x might be equal to zsl->header, so test if obj is non-NULL */
        if x.key != "" && strings.Compare(x.key, key) == 0 {
            return rank
        }

	}
	return 0
}

// Finds an element by its rank
func (zsl *zskiplist) zslGetElementByRank(rank int) *zskiplistNode {
	traversed := 0
	x := zsl.header
	for i := zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && (traversed+x.level[i].span) <= rank {
			traversed += x.level[i].span
			x = x.level[i].forward
		}
		if traversed == rank {
			return x
		}
	}
	return nil
}
