package bubble

import (
	"fmt"
	"testing"
)

func TestBubble1(t *testing.T) {
	// t.Skip()
	testArray := []int{5, 8, 6, 3, 9, 2, 1, 7}
	fmt.Println("TestBubble1...")
	bubbleSort(testArray)
	fmt.Println(testArray)
}

func TestBubble2(t *testing.T) {
	// t.Skip()
	testArray := []int{3, 4, 2, 1, 5, 6, 7, 8}
	fmt.Println("TestBubble2...")
	bubbleSort(testArray)
	fmt.Println(testArray)
}

func TestCockTailSort(t *testing.T) {
	t.Skip()
	// testArray := []int{3, 4, 2, 1, 5, 6, 7, 8}
	testArray := []int{5, 8, 6, 3, 9, 2, 1, 7}
	fmt.Println("cockTailSort...")
	cockTailSort(testArray)
	fmt.Println(testArray)
}
