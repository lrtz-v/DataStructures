package paxos

import (
	"context"
	"fmt"
	"math/rand"
	"sort"
	"sync/atomic"
	"time"
)

var (
	RoundNum int64
)

type Server struct {
	NodeSize   int     //  节点数
	Quorum     int     // 半数
	PaxosNodes []*Node // paxos节点
}

func NewPaxosServer(size int) *Server {
	RoundNum = 0
	paxosServer := &Server{NodeSize: size, PaxosNodes: []*Node{}, Quorum: 1 + size/2}
	for i := 0; i < size; i++ {
		paxosServer.PaxosNodes = append(paxosServer.PaxosNodes, NewNode())
	}
	return paxosServer
}

func (p *Server) randomHalf() []int {
	ids := []int{}
	indexes := make(map[int]int)
	for i := 0; i < p.Quorum; i++ {
		for {
			nodeIndex := rand.Intn(p.NodeSize)
			if _, ok := indexes[nodeIndex]; ok {
				continue
			} else {
				indexes[nodeIndex] = nodeIndex
				break
			}
		}
	}

	for l := range indexes {
		ids = append(ids, l)
	}
	return ids
}

// 读取key最新的值
func (p *Server) Get(key string) int {
	// todo
	return 0
}

// 设置key的值为value，并更新版本
func (p *Server) Set(key string, value int) error {
	atomic.AddInt64(&RoundNum, 1)
	currentRnd := RoundNum
	phrase1Responses := []Phrase1Response{}

	phrase1ResponseChan := make(chan Phrase1Response, p.NodeSize)
	parentContext := context.Background()
	ctx, cancelFunc := context.WithTimeout(parentContext, time.Duration(3)*time.Second)
	defer cancelFunc()
	for i := 0; i < p.NodeSize; i++ {
		go func(ctx context.Context, index int, ch chan Phrase1Response) {
			ch <- p.PaxosNodes[index].Phrase1(key, value, currentRnd)
		}(ctx, i, phrase1ResponseChan)
	}

	select {
	case res := <-phrase1ResponseChan:
		phrase1Responses = append(phrase1Responses, res)
	case <-ctx.Done():
		fmt.Println("phrase 1 timout ", ctx.Err())
	}

	// 相应节点少于半数，流程失败
	if len(phrase1Responses) < p.Quorum {
		return PaxosNodeLessThanQuorum
	}

	// 如果返回的last rnd大于当前rnd，退出
	// 从所有应答中选出vrnd最大的值
	for _, v := range phrase1Responses {
		if v.Err != nil {
			return Phrase1Failed
		}
		if v.LastRnd > currentRnd {
			return CurrentRndNumInvalid
		}
	}
	sort.Slice(phrase1Responses, func(i, j int) bool {
		if phrase1Responses[i].Value == nil {
			return true
		}
		return phrase1Responses[i].Value.VRnd < phrase1Responses[j].Value.VRnd
	})
	if phrase1Responses[len(phrase1Responses)-1].Value != nil {
		// 如果收到了某个应答包含被写入的v和vrnd, 这时, Proposer X 必须假设有其他客户端(Proposer) 正在运行, 虽然X不知道对方是否已经成功结束, 但任何已经写入的值都不能被修改!
		// 所以X必须保持原有的值. 于是X将看到的最大vrnd对应的v作为X的phase-2将要写入的值
		value = phrase1Responses[len(phrase1Responses)-1].Value.Value
	} else {
		// 如果所有的应答中，value都是空，可以直接写入自己的value，可以进行phrase-2写入
	}

	phrase2Responses := []bool{}
	phrase2ResponseChan := make(chan bool, p.NodeSize)
	ctx, cancelFunc = context.WithTimeout(parentContext, time.Duration(3)*time.Second)
	defer cancelFunc()
	for i := 0; i < p.NodeSize; i++ {
		go func(ctx context.Context, index int, ch chan bool) {
			ch <- p.PaxosNodes[index].Phrase2(key, value, currentRnd)
		}(ctx, i, phrase2ResponseChan)
	}

	select {
	case res := <-phrase2ResponseChan:
		phrase2Responses = append(phrase2Responses, res)
	case <-ctx.Done():
		fmt.Println("phrase 2 timout ", ctx.Err())
	}

	return nil
}

// 指定key的值增加value
func (p *Server) Inc(key string, value int) {
	// todo
}
