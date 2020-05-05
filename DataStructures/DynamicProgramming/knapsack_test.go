package dynamic

import (
	"testing"
)


func TestKnapsack(t *testing.T) {
	itemsWightList := []int{2, 2, 4, 6, 3}
	itemsNum := 5
	packageMaxWeight := 9
	maxWeight := knapsack(itemsWightList, itemsNum, packageMaxWeight)
	if maxWeight != 9 {
		t.Fatalf("[*]ERROR! Expected 9, Got %d", maxWeight)
	}
}

func TestKnapsack2(t *testing.T) {
	itemsWightList := []int{2, 2, 4, 6, 3}
	itemsNum := 5
	packageMaxWeight := 9
	maxWeight := knapsack2(itemsWightList, itemsNum, packageMaxWeight)
	if maxWeight != 9 {
		t.Fatalf("[*]ERROR! Expected 9, Got %d", maxWeight)
	}
}
