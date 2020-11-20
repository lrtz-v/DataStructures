package sort

import (
	"testing"
)

func TestBucketSort(t *testing.T) {
	data := []int64{1, 0, 10, 9, 8, 3, 5, 100}
	bucket := Bucket{}
	l := bucket.Sort(data)
	if !bucket.IsSorted(l) {
		t.Fatal(l)
		t.Fatal("Bucket Sort Failed")
	}
}
