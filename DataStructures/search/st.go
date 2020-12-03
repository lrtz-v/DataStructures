package search

import (
	"DataStructures/entry"
	"time"
)

// ST 符号表API
type ST struct {
	seed     uint64 // hash seed
	sizemask int
}

func (st *ST) hasher(key entry.Key) uintptr {
	return key.Hash(st.seed)
}

func (st *ST) bucket(key entry.Key) uintptr {
	hash := st.hasher(key)
	return hash & uintptr(st.sizemask)
}

// NewST create new ST
func NewST() *ST {
	st := &ST{}
	st.seed = uint64(time.Now().Nanosecond())
	return st
}

// Put 将键值对存入表中（若值为空则将key删除）
func (st *ST) Put(key entry.Key, value entry.Value) {

}

// Get 获取key对应的value，不存在返回error
func (st *ST) Get(key entry.Key) (entry.Value, error) {
	return 0, nil
}

// Delete 删除key
func (st *ST) Delete(key entry.Key) {
}

// Contains 是否包含key
func (st *ST) Contains(key entry.Key) bool {
	return true
}

// IsEmpty 收否为空
func (st *ST) IsEmpty() bool {
	return true
}

// Size 返回键值对数量
func (st *ST) Size() int {
	return 0
}

// Keys 返回键的集合
func (st *ST) Keys() []entry.Key {
	return []entry.Key{}
}
