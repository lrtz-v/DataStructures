package core

import (
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"math"
	"sort"
	"sync"
	"sync/atomic"
)

const (
	// 虚拟节点命名格式
	hostReplicaFormat = `%s%d`
)

var (
	// 每台物理服务器虚拟节点个数
	defaultReplicaNum = 10

	// 负载因子
	loadBoundFactor = 0.25

	// 默认Hash函数
	defaultHashFunc = func(key string) uint64 {
		out := sha512.Sum512([]byte(key))
		return binary.LittleEndian.Uint64(out[:])
	}
)

// 一致性Hash结构定义
type Consistent struct {
	// 虚拟节点数量
	replicaNum int

	// 总负载
	// https://research.googleblog.com/2017/04/consistent-hashing-with-bounded-loads.html
	totalLoad int64

	// Hash函数
	hashFunc func(key string) uint64

	// 服务服务器名称与对应Host的映射
	hostMap map[string]*Host

	// 虚拟节点对应的物理服务器映射
	replicaHostMap map[uint64]string

	// Hash环，可使用其他数据结构
	sortedHostsHashSet []uint64

	// 锁
	sync.RWMutex
}

func NewConsistent(replicaNum int, hashFunc func(key string) uint64) *Consistent {
	if replicaNum <= 0 {
		replicaNum = defaultReplicaNum
	}

	if hashFunc == nil {
		hashFunc = defaultHashFunc
	}

	return &Consistent{
		replicaNum:         replicaNum,
		totalLoad:          0,
		hashFunc:           hashFunc,
		hostMap:            make(map[string]*Host),
		replicaHostMap:     make(map[uint64]string),
		sortedHostsHashSet: make([]uint64, 0),
	}
}

func (c *Consistent) RegisterHost(hostName string) error {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.hostMap[hostName]; ok {
		return ErrHostAlreadyExists
	}

	c.hostMap[hostName] = &Host{
		Name:      hostName,
		LoadBound: 0,
	}

	// 添加虚拟节点
	for i := 0; i < c.replicaNum; i++ {
		hashedIdx := c.hashFunc(fmt.Sprintf(hostReplicaFormat, hostName, i))
		c.replicaHostMap[hashedIdx] = hostName
		c.sortedHostsHashSet = append(c.sortedHostsHashSet, hashedIdx)
	}

	// 虚拟节点排序
	sort.Slice(c.sortedHostsHashSet, func(i int, j int) bool {
		if c.sortedHostsHashSet[i] < c.sortedHostsHashSet[j] {
			return true
		}
		return false
	})

	return nil
}

func (c *Consistent) UnregisterHost(hostName string) error {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.hostMap[hostName]; !ok {
		return ErrHostNotFound
	}

	delete(c.hostMap, hostName)

	for i := 0; i < c.replicaNum; i++ {
		hashedIdx := c.hashFunc(fmt.Sprintf(hostReplicaFormat, hostName, i))
		delete(c.replicaHostMap, hashedIdx)
		c.delHashIndex(hashedIdx)
	}

	return nil
}

// 虚拟节点删除
func (c *Consistent) delHashIndex(val uint64) {
	idx := -1
	l := 0
	r := len(c.sortedHostsHashSet) - 1
	for l <= r {
		m := (l + r) / 2
		if c.sortedHostsHashSet[m] == val {
			idx = m
			break
		} else if c.sortedHostsHashSet[m] < val {
			l = m + 1
		} else if c.sortedHostsHashSet[m] > val {
			r = m - 1
		}
	}
	if idx != -1 {
		c.sortedHostsHashSet = append(c.sortedHostsHashSet[:idx], c.sortedHostsHashSet[idx+1:]...)
	}
}

// 更新负载
func (c *Consistent) UpdateLoad(host string, load int64) {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.hostMap[host]; !ok {
		return
	}
	c.totalLoad = c.totalLoad - c.hostMap[host].LoadBound + load
	c.hostMap[host].LoadBound = load
}

// 获取服务器列表
func (c *Consistent) Hosts() []string {
	c.RLock()
	defer c.RUnlock()

	hosts := make([]string, 0)
	for k := range c.hostMap {
		hosts = append(hosts, k)
	}
	return hosts
}

// 查询key所在服务器
func (c *Consistent) GetKey(key string) (string, error) {
	hashedKey := c.hashFunc(key)
	idx := c.searchKey(hashedKey)
	return c.replicaHostMap[c.sortedHostsHashSet[idx]], nil
}

// 找到key所在、且负载较小的服务器
func (c *Consistent) GetKeyLeast(key string) (string, error) {
	c.RLock()
	defer c.RUnlock()

	if len(c.replicaHostMap) == 0 {
		return "", ErrHostNotFound
	}

	hashedKey := c.hashFunc(key)
	idx := c.searchKey(hashedKey) // Find the first host that may serve the key

	i := idx
	for {
		host := c.replicaHostMap[c.sortedHostsHashSet[i]]
		loadChecked, err := c.checkLoadCapacity(host)
		if err != nil {
			return "", err
		}
		if loadChecked {
			return host, nil
		}
		i++

		// if idx goes to the end of the ring, start from the beginning
		if i >= len(c.replicaHostMap) {
			i = 0
		}
	}
}

// 增加负载
func (c *Consistent) Inc(hostName string) {
	c.Lock()
	defer c.Unlock()

	atomic.AddInt64(&c.hostMap[hostName].LoadBound, 1)
	atomic.AddInt64(&c.totalLoad, 1)
}

// 降低负载
func (c *Consistent) Done(host string) {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.hostMap[host]; !ok {
		return
	}
	atomic.AddInt64(&c.hostMap[host].LoadBound, -1)
	atomic.AddInt64(&c.totalLoad, -1)
}

// 获取全部负载
func (c *Consistent) GetLoads() map[string]int64 {
	c.RLock()
	defer c.RUnlock()

	loads := make(map[string]int64)
	for k, v := range c.hostMap {
		loads[k] = atomic.LoadInt64(&v.LoadBound)
	}
	return loads
}

// 获取单节点最大负载
func (c *Consistent) MaxLoad() int64 {
	if c.totalLoad == 0 {
		c.totalLoad = 1
	}

	var avgLoadPerNode float64
	avgLoadPerNode = float64(c.totalLoad / int64(len(c.hostMap)))
	if avgLoadPerNode == 0 {
		avgLoadPerNode = 1
	}
	avgLoadPerNode = math.Ceil(avgLoadPerNode * (1 + loadBoundFactor))
	return int64(avgLoadPerNode)
}

// 获取key所在虚拟节点
func (c *Consistent) searchKey(key uint64) int {
	idx := sort.Search(len(c.sortedHostsHashSet), func(i int) bool {
		return c.sortedHostsHashSet[i] >= key
	})

	if idx >= len(c.sortedHostsHashSet) {
		idx = 0
	}

	return idx
}

// 检查是否在负载范围内
func (c *Consistent) checkLoadCapacity(host string) (bool, error) {
	if c.totalLoad < 0 {
		c.totalLoad = 0
	}

	var avgLoadPerNode float64
	avgLoadPerNode = float64((c.totalLoad + 1) / int64(len(c.hostMap)))
	if avgLoadPerNode == 0 {
		avgLoadPerNode = 1
	}
	avgLoadPerNode = math.Ceil(avgLoadPerNode * (1 + loadBoundFactor))

	candidateHost, ok := c.hostMap[host]
	if !ok {
		return false, ErrHostNotFound
	}

	if float64(candidateHost.LoadBound)+1 <= avgLoadPerNode {
		return true, nil
	}

	return false, nil
}
