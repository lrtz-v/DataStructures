package search

import "DataStructures/entry"

// SortedST 有序符号表API
type SortedST struct {
}

// Put 将键值对存入表中（若值为空则将key删除）
func (st *SortedST) Put(key entry.Key, value entry.Value) {

}

// Get 获取key对应的value，不存在返回error
func (st *SortedST) Get(key entry.Key) (entry.Value, error) {
	return 0, nil
}

// Delete 删除key
func (st *SortedST) Delete(key entry.Key) {
}

// Contains 是否包含key
func (st *SortedST) Contains(key entry.Key) bool {
	return true
}

// IsEmpty 收否为空
func (st *SortedST) IsEmpty() bool {
	return true
}

// Size 返回键值对数量
func (st *SortedST) Size() int {
	return 0
}

// RangeSize 返回[lo, hi]之间键的数量
func (st *SortedST) RangeSize(lo, hi entry.Key) []entry.Key {
	return []entry.Key{}
}

// Keys 返回键的集合
func (st *SortedST) Keys() []entry.Key {
	return []entry.Key{}
}

// RangeKeys 返回[lo, hi]之间的所有键，已排序
func (st *SortedST) RangeKeys(lo, hi entry.Key) []entry.Key {
	return []entry.Key{}
}

// Min 返回最小值
func (st *SortedST) Min() int {
	return 0
}

// Max 返回最大值
func (st *SortedST) Max() int {
	return 0
}

// Floor 返回小于等于key的最大值
func (st *SortedST) Floor(key entry.Key) entry.Value {
	return 0
}

// Ceiling 返回大于等于key的最小值
func (st *SortedST) Ceiling(key entry.Key) entry.Key {
	return entry.Key("")
}

// Rank 返回小于key的键的数量
func (st *SortedST) Rank(key entry.Key) int {
	return 0
}

// Select 返回排名为k的键
func (st *SortedST) Select(k int) *entry.Entry {
	return &entry.Entry{}
}

// DelMin 删除最小键
func (st *SortedST) DelMin() {

}

// DelMax 删除最大值
func (st *SortedST) DelMax() {

}
