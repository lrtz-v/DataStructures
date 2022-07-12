package paxos

type Element struct {
	Key   string // 键
	Value int    // 值
	VRnd  int64  // vrnd 跟v是一对, 它记录了在哪个Round中v被写入了
}

// Paxos节点，每个节点维护存储内容（这里用hashTable简单代替）、上一轮Rnd
type Node struct {
	HashTable map[string]*Element
	LastRnd   int64 // Acceptor记住的最后一次进行写前读取的Proposer(客户端)是谁, 以此来决定谁可以在后面真正把一个值写到存储中
}

type Phrase1Response struct {
	LastRnd int64
	Value   *Element
	Err     error
}

// 阶段一
// return:
// 			last rnd
// 			value of key
// 			value rnd
func (n *Node) Phrase1(key string, value int, rnd int64) Phrase1Response {
	// 如果请求rnd比last rnd小，直接拒绝请求
	if n.LastRnd > 0 && n.LastRnd > rnd {
		return Phrase1Response{LastRnd: n.LastRnd, Value: nil, Err: RndInvalid}
	}
	// 将请求的rnd保存下来，只接受这个rnd的phrase-2请求
	n.LastRnd = rnd

	// 返回，带上之前的last rnd和值，以及值对应的rnd
	if _, ok := n.HashTable[key]; ok {
		return Phrase1Response{LastRnd: n.LastRnd, Value: n.HashTable[key], Err: nil}
	} else {
		return Phrase1Response{LastRnd: n.LastRnd, Value: nil, Err: nil}
	}
}

func (n *Node) Phrase2(key string, value int, rnd int64) bool {
	if rnd != n.LastRnd {
		return false
	}
	if _, ok := n.HashTable[key]; ok {
		n.HashTable[key].Value = value
		n.HashTable[key].VRnd = rnd
	} else {
		n.HashTable[key] = &Element{Key: key, Value: value, VRnd: rnd}
	}
	return true
}

func NewNode() *Node {
	return &Node{HashTable: make(map[string]*Element)}
}

func (n *Node) Get(key string) (*Element, error) {
	if v, ok := n.HashTable[key]; !ok {
		return nil, KeyNotExists
	} else {
		return v, nil
	}
}

// 设置key的值为value，并更新版本
func (n *Node) Set(key string, value int) {
	if _, ok := n.HashTable[key]; ok {
		n.HashTable[key].Value = value

	} else {
		n.HashTable[key] = &Element{Key: key, Value: value}
	}
}

// 指定key的值增加value
func (n *Node) Inc(key string, value int) {
	if _, ok := n.HashTable[key]; ok {
		n.HashTable[key].Value += value
	} else {
		n.HashTable[key] = &Element{Key: key, Value: value}
	}
}
