package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Skip()
	list1 := []int{1, 3, 3, 5, 9, 0, 6}
	target1 := 14
	result1 := twoSum(list1, target1)
	if list1[result1[0]]+list1[result1[1]] != target1 {
		t.Fatalf("%d + %d did not equal %d\n", list1[result1[0]], list1[result1[1]], target1)
	}

	list1 = []int{585, 579, 245, 512, 37, 919, 219, 49, 565, 243, 944, 963, 894, 78, 457, 109, 976, 168, 340, 446, 618, 910, 278, 118, 453, 258, 727, 54, 25, 788, 195, 457, 81, 763, 510, 865, 313, 615, 26, 419, 185, 789, 218, 275, 636, 6, 231, 71, 143, 506, 460, 534, 801, 870, 531, 961, 409, 618, 657, 194, 896, 331, 118, 143, 896, 639, 579, 99, 304, 247, 582, 261, 551, 309, 411, 219, 939, 99, 265, 334, 590, 698, 903, 909, 256, 127, 749, 385, 939, 753, 700, 313, 799, 915, 954, 706, 518, 303, 360, 903}
	target1 = 1842
	result1 = twoSum(list1, target1)
	if list1[result1[0]]+list1[result1[1]] != target1 {
		t.Fatalf("%d + %d did not equal %d\n", list1[result1[0]], list1[result1[1]], target1)
	}
}

// go test -bench=. -run=none -benchtime=10s
func BenchmarkSprintf(b *testing.B) {
	b.Skip()
	list1 := []int{585, 579, 245, 512, 37, 919, 219, 49, 565, 243, 944, 963, 894, 78, 457, 109, 976, 168, 340, 446, 618, 910, 278, 118, 453, 258, 727, 54, 25, 788, 195, 457, 81, 763, 510, 865, 313, 615, 26, 419, 185, 789, 218, 275, 636, 6, 231, 71, 143, 506, 460, 534, 801, 870, 531, 961, 409, 618, 657, 194, 896, 331, 118, 143, 896, 639, 579, 99, 304, 247, 582, 261, 551, 309, 411, 219, 939, 99, 265, 334, 590, 698, 903, 909, 256, 127, 749, 385, 939, 753, 700, 313, 799, 915, 954, 706, 518, 303, 360, 903}
	target1 := 1842
	result1 := twoSum(list1, target1)
	if list1[result1[0]]+list1[result1[1]] != target1 {
		b.Fatalf("%d + %d did not equal %d\n", list1[result1[0]], list1[result1[1]], target1)
	}
}
