package array

import "errors"

/*
	- ### 数组（Arrays）
		- 实现一个可自动调整大小的动态数组。
		- [ ] 介绍：
		- [ ] 实现一个动态数组（可自动调整大小的可变数组）
		- 可以使用 int 类型的数组，但不能使用其语法特性
		- 从大小为8或更大的数（使用2的倍数 —— 16、32、64、128）开始编写
		- [ ] size() —— 数组元素的个数
		- [ ] capacity() —— 可容纳元素的个数
		- [ ] is_empty()
		- [ ] at(index) —— 返回对应索引的元素，且若索引越界则愤然报错
		- [ ] push(item) ——
		- [ ] insert(index, item) —— 在指定索引中插入元素，并把后面的元素依次后移
		- [ ] prepend(item) —— 可以使用上面的 insert 函数，传参 index 为 0
		- [ ] pop() —— 删除在数组末端的元素，并返回其值
		- [ ] delete(index) —— 删除指定索引的元素，并把后面的元素依次前移
		- [ ] remove(item) —— 删除指定值的元素，并返回其索引（即使有多个元素）
		- [ ] find(item) —— 寻找指定值的元素并返回其中第一个出现的元素其索引，若未找到则返回 -1
		- [ ] resize(new_capacity)
		- 若数组的大小到达其容积，则变大一倍
		- 获取元素后，若数组大小为其容积的1/4，则缩小一半
		- [ ] 时间复杂度
		- 在数组末端增加/删除、定位、更新元素，只允许占 O(1) 的时间复杂度（平摊（amortized）去分配内存以获取更多空间）
		- 在数组任何地方插入/移除元素，只允许 O(n) 的时间复杂度
		- [ ] 空间复杂度
		- 因为在内存中分配的空间邻近，所以有助于提高性能
		- 空间需求 = （大于或等于 n 的数组容积）* 元素的大小。即便空间需求为 2n，其空间复杂度仍然是 O(n)
*/

const (
	defaultLength = 8
	stage         = 2
)

type array struct {
	arrLength   int   // 数组大小
	arrCapacity int   // 数组内的元素个数
	arrList     []int // 数组
}

func (a *array) bigger(nowCapacity int) error {
	// 1. check space 若数组的大小到达其容积，则变大一倍
	if nowCapacity < a.arrLength {
		return nil
	}

	// 2. resize
	a.arrLength = a.arrLength * stage
	destArray := make([]int, a.arrLength)

	// 3. copy data
	for index, item := range a.arrList {
		destArray[index] = item
	}
	a.arrList = destArray
	return nil
}

func (a *array) smaller(nowCapacity int) error {
	// 1. check space 获取元素后，若数组大小为其容积的1/4，则缩小一半
	if nowCapacity >= a.arrLength>>2 || a.arrLength == defaultLength {
		return nil
	}

	// 2. resize
	a.arrLength = int(a.arrLength / stage)
	destArray := make([]int, a.arrLength)

	// 3. copy data
	for index, item := range a.arrList {
		if item != 0 {
			destArray[index] = item
		}
	}
	a.arrList = destArray
	return nil
}

func (a *array) getSize() int {
	// 数组大小
	return a.arrLength
}

func (a *array) getCapacity() int {
	// 可容纳元素的个数
	return a.arrLength - a.arrCapacity
}

func (a *array) isEmpty() bool {
	// 是否为空
	// true: 空
	return a.arrCapacity == 0
}

func (a *array) at(index int) (int, error) {
	// 返回对应索引的元素，且若索引越界则愤然报错
	if index > a.arrLength-1 {
		return 0, errors.New("Out of range")
	}
	return a.arrList[index], nil
}

func (a *array) insert(item, index int) error {
	// 指定位置插入

	// 1. 检查index合法性
	if index < 0 && index > a.arrLength-1 {
		return errors.New("Index ourt of range")
	}

	// resize
	err := a.bigger(a.arrCapacity + 1)
	if err != nil {
		return errors.New("reSize Error")
	}

	// insert
	for i := a.arrLength - 1; i > index; i-- {
		if a.arrList[i-1] != 0 {
			a.arrList[i] = a.arrList[i-1]
			a.arrList[i-1] = 0
		}
	}
	a.arrList[index] = item

	a.arrCapacity++
	return nil
}

func (a *array) prepend(item int) error {
	// 首部插入
	a.insert(item, 0)
	return nil
}

func (a *array) push(item int) error {
	// 尾部插入
	a.insert(item, a.arrLength-1)
	return nil
}

func (a *array) pop() int {
	value := a.arrList[a.arrLength-1]
	a.arrList[a.arrLength-1] = 0
	a.arrCapacity--
	return value
}

func (a *array) delete(index int) error {
	// 删除指定位置的元素
	// 1. 检查index合法性
	if index < 0 && index > a.arrLength-1 {
		return errors.New("Index ourt of range")
	}

	a.arrList[index] = 0
	// delete
	for i := index; i < a.arrLength-1; i++ {
		a.arrList[i] = a.arrList[i+1]
		a.arrList[i+1] = 0
	}
	a.arrCapacity--

	// resize
	err := a.smaller(a.arrCapacity - 1)
	if err != nil {
		return errors.New("reSize Error")
	}

	return nil
}

func (a *array) remove(item int) []int {
	indexList := make([]int, 0, 0)
	for index, value := range a.arrList {
		if item == value {
			a.delete(index)
			indexList = append(indexList, index)
		}
	}
	return indexList
}

func (a *array) find(item int) int {
	for index, value := range a.arrList {
		if item == value {
			return index
		}
	}
	return -1
}

func initArray() *array {
	var initSlice = make([]int, defaultLength)
	return &array{arrLength: defaultLength, arrCapacity: 0, arrList: initSlice}
}
