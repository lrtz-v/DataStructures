package maxScoreSightseeingPair

import "testing"

func TestMaxScoreSightseeingPair(t *testing.T) {
	var n int
	n = maxScoreSightseeingPair([]int{8,1,5,2,6})
	if n != 11 {
		t.Fatalf("[*] TestMaxScoreSightseeingPair Error! Expected 11 Got %d.", n)
	}
	n = maxScoreSightseeingPair([]int{2,2,2})
	if n != 3 {
		t.Fatalf("[*] TestMaxScoreSightseeingPair Error! Expected 3 Got %d.", n)
	}
	n = maxScoreSightseeingPair([]int{2,1})
	if n != 2 {
		t.Fatalf("[*] TestMaxScoreSightseeingPair Error! Expected 2 Got %d.", n)
	}
	n = maxScoreSightseeingPair([]int{1, 3, 5})
	if n != 7 {
		t.Fatalf("[*] TestMaxScoreSightseeingPair Error! Expected 7 Got %d.", n)
	}
}
