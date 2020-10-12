package mintime

import "testing"

func testMinTime(t *testing.T) {

	edges := [][]int{
		[]int{0,1}, 
		[]int{0,2}, 
		[]int{1,4}, 
		[]int{1,5}, 
		[]int{2,3}, 
		[]int{2,6},
	}
	hasApple := []bool{false,false,true,false,true,true,false}

	s := minTime(7, edges, hasApple)
	if s != 8 {
		t.Fatalf("\n[*] testMinTime Failed. Expected 8, Got: %d\n", s)
	}
}
