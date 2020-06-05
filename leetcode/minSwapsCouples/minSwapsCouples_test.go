package minSwapsCouples

import (
	"testing"
)

func TestMinSwapsCouples(t *testing.T) {
	t.Skip()
	row := []int{0, 2, 1, 3}
	times := minSwapsCouples(row)
	if times != 1 {
		t.Fatal("times err")
	}
}

func TestMinSwapsCouples2(t *testing.T) {
	t.Skip()
	row := []int{3, 2, 0, 1}
	times := minSwapsCouples(row)
	if times != 0 {
		t.Fatal("times err")
	}
}

func TestMinSwapsCouples3(t *testing.T) {
	t.Skip()
	row := []int{5, 3, 4, 2, 1, 0}
	times := minSwapsCouples(row)
	if times != 1 {
		t.Fatal("times err")
	}
}
