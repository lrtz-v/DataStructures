package quik

import (
	"fmt"
	"testing"
)

func TestSort1(t *testing.T) {
	testData := []int{3, 6, 7, 8, 2, 10, 9}
	quickSort2(testData, 0, len(testData)-1)
	fmt.Println(testData)
}

func TestSort2(t *testing.T) {
	testData := []int{3, 3, 3, 3, 3, 4}
	quickSort2(testData, 0, len(testData)-1)
	fmt.Println(testData)
}
