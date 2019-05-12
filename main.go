package main

import (
	"fmt"

	"github.com/DataStructures/leetcode"
)

func main() {
	testArr := []int{1, 2, 3, 5, 4}
	testArr, err := leetcode.FindNearestNumber(testArr)
	if err != nil {
		panic(err)
	}
	fmt.Println(testArr)
}
