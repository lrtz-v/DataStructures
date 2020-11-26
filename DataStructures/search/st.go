package search

// ST 符号表API
type ST struct {
}

// Put 将键值对存入表中（若值为空则将key删除）
func (st *ST) Put(key string, value int) {

}

// Get 获取key对应的value，不存在返回error
func (st *ST) Get(key string) (int, error) {
	return 0, nil
}

// Delete 删除key
func (st *ST) Delete(key string) {
}

// Contains 是否包含key
func (st *ST) Contains(key string) bool {
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
func (st *ST) Keys() []string {
	return []string{}
}
