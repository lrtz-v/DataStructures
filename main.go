package main

import (
	"fmt"

	"github.com/DataStructures/leetcode"
)

func main() {
	testArr := []int{1, 2, 3, 5, 4}
	testArr = leetcode.FindNearestNumber(testArr)
	fmt.Println(testArr)
}
