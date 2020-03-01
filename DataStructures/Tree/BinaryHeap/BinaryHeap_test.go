package binaryheap

import (
	"testing"
)

func TestFindFirstEqual(t *testing.T) {
	// t.Skip()
	testData := []int{3, 6, 7, 8, 8, 10}
	i := findFirstEqual(testData,8)
	if i != 3 {
		t.Fatal("TestFindFirstEqual Failed")
	}
}

func TestFindLastEqual(t *testing.T) {
	// t.Skip()
	testData := []int{3, 6, 7, 8, 8, 10}
	i := findLastEqual(testData, 8)
	if i != 4 {
		t.Fatal("TestFindLastEqual Failed")
	}
}

func TestFindFirstBE(t *testing.T) {
	// t.Skip()
	testData := []int{1, 3, 5, 7, 7, 9}
	i := findFirstBE(testData, 6)
	if i != 3 {
		t.Fatal("TestFindFirstBE Failed")
	}

	i = findFirstBE(testData, 5)
	if i != 2 {
		t.Fatal("TestFindFirstBE Failed")
	}

	i = findFirstBE(testData, 7)
	if i != 3 {
		t.Fatal("TestFindFirstBE Failed")
	}
}

func TestFindLastLE(t *testing.T) {
	// t.Skip()
	testData := []int{1, 3, 5, 7, 7, 9}
	i := findLastLE(testData, 8)
	if i != 4 {
		t.Fatal("TestFindLastLE Failed")
	}

	i = findLastLE(testData, 5)
	if i != 2 {
		t.Fatal("TestFindLastLE Failed")
	}

	i = findLastLE(testData, 6)
	if i != 2 {
		t.Fatal("TestFindLastLE Failed")
	}
}

func TestFindEqualInCycleArr(t *testing.T) {
	// t.Skip()

	testData2 := []int{4, 5, 6, 7, 1, 2, 3}
	i := findEqualInCycleArr(testData2, 2)
	if i != 5 {
		t.Fatal("TestFindEqualInCycleArr Failed")
	}

	testData2 = []int{4, 5, 6, 1, 2, 3}
	i = findEqualInCycleArr(testData2, 2)
	if i != 4 {
		t.Fatal("TestFindEqualInCycleArr Failed")
	}

	testData := []int{5, 1, 3}
	i = findEqualInCycleArr(testData, 5)
	if i != 0 {
		t.Fatal("TestFindEqualInCycleArr Failed")
	}

	testData = []int{1, 3}
	i = findEqualInCycleArr(testData, 0)
	if i != -1 {
		t.Fatal("TestFindEqualInCycleArr Failed")
	}

	testData = []int{1, 3, 5}
	i = findEqualInCycleArr(testData, 0)
	if i != -1 {
		t.Fatal("TestFindEqualInCycleArr Failed")
	}
}
