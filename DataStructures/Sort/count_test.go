package sort

import (
	"testing"
)

func TestCountSort(t *testing.T) {
	data := []int64{1, 4, 9, 0, 2, 10, 100}
	c := Count{}
	l := c.Sort(data)
	if !c.IsSorted(l) {
		t.Fatal(l)
		t.Fatal("Count Sort failed")
	}
}
