package largestRectangleArea

import "testing"

func TestlargestRectangleArea(t *testing.T) {
	s := largestRectangleArea([]int{2,1,5,6,2,3})
	if s != 10 {
		t.Fatalf("[*] largestRectangleArea error! Expected 10 Got %d\n", s)
	}
	if s != largestRectangleArea2([]int{2,1,5,6,2,3}) {
		t.Fatal("[*] largestRectangleArea2 Error!")
	}
}
