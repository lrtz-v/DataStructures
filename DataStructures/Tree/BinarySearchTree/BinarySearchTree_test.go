package BinarySearchTree

import (
	"fmt"
	"testing"
)

func TestBinaryTreeRevert(t *testing.T) {
	t.Skip()
	// 11 15 10 13 3 14 12 8 15 3 1 13 2 9 2
	// testData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// testData := []int{1, 2, 3, 4, 5, 6, 7}
	testData := []int{4, 2, 7, 1, 3, 6, 9}
	BinaryTree := NewBinaryTree()
	for _, data := range testData {
		BinaryTree.Insert(data)
	}
	BinaryTree.PreOrderTraversal(BinaryTree.Root)

	BinaryTree.Revert(BinaryTree.Root)
	fmt.Println("revert done")
	BinaryTree.PreOrderTraversal(BinaryTree.Root)

	BinaryTree.Clear()
}

func TestBinaryTreeLevelRevert(t *testing.T) {
	t.Skip()
	testData := []int{4, 2, 7, 1, 3, 6, 9}
	BinaryTree := NewBinaryTree()
	for _, data := range testData {
		BinaryTree.Insert(data)
	}
	BinaryTree.LevelTraversal()
	BinaryTree.Clear()
}

func TestBinaryTreeLevelRevert2(t *testing.T) {
	// t.Skip()
	testData := []int{4, 2, 7, 1, 3, 6, 9}
	BinaryTree := NewBinaryTree()
	for _, data := range testData {
		BinaryTree.Insert(data)
	}
	BinaryTree.LevelTraversal()
	fmt.Println("...")
	high := BinaryTree.GetHight(BinaryTree.Root)
	for i := 1; i < high+1; i++ {
		BinaryTree.LevelTraversalRecursion(BinaryTree.Root, i)
	}
	BinaryTree.Clear()
}
