package skiplist

const ZSKIPLIST_MAXLEVEL = 32

type zskiplistLevel struct {
	forward *zskiplistLevel // 指向链表下一个节点的指针，即前向指针
	span    int            // 表示这个前向指针跳跃过了多少个节点
}

type zskiplistNode struct {
	obj      *interface{}      // 存放该节点的数据
	score    float32          // 数据对应的分数值，zset通过分数对数据升序排列
	backward *zskiplistNode    // 指向链表上一个节点的指针，即后向指针
	level    []zskiplistLevel // 表示跳表中的一层
}

type zskiplist struct {
	header *zskiplistNode // 头指针
	tail   *zskiplistNode // 尾指针
	length int64         // 跳表的长度，不包括头指针
	level  int           // 跳表的层数
}

func zslInit() *zskiplist {
	zsl := &zskiplist{level: 1, length: 0}
	zsl.header = zslNodeInit(ZSKIPLIST_MAXLEVEL, 0)
	for j := 0; j < ZSKIPLIST_MAXLEVEL; j++ {
		zsl.header.level[j].forward = nil
		zsl.header.level[j].span = 0
	}
	zsl.header.backward = nil
	zsl.tail = nil
	return zsl
}

func zslNodeInit(level int, score float32) *zskiplistNode {
	l := make([]zskiplistLevel, level)
	return &zskiplistNode{obj:nil, score:score, backward:nil, level:l}
}
