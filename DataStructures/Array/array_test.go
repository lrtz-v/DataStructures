package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	t.Skip()
	array := initArray()
	fmt.Println("array: ", array)
	fmt.Println("array len: ", array.getSize())
	fmt.Println("array capacity: ", array.getCapacity())

	array.insert(111, 1)
	array.insert(222, 2)
	array.insert(333, 3)
	array.insert(444, 4)
	array.insert(555, 5)
	array.insert(666, 6)
	array.insert(777, 7)
	array.insert(888, 8)
	array.insert(999, 9)
	fmt.Println("array: ", array)
	fmt.Println("array len: ", array.getSize())
	fmt.Println("array capacity: ", array.getCapacity())

	array.delete(0)
	array.delete(1)
	array.delete(2)
	array.delete(3)
	array.delete(4)
	array.delete(5)
	fmt.Println("array: ", array)
	fmt.Println("array len: ", array.getSize())
	fmt.Println("array capacity: ", array.getCapacity())

}
