package findNearestNumber

import (
	"testing"
)

func TestFindNearestNumber(t *testing.T) {
	t.Skip()
	testArr := []int{1, 2, 3, 5, 4}
	testArr, err := FindNearestNumber(testArr)
	if err != nil {
		t.Errorf("findd err")
	}
	//fmt.Println(testArr)
}

func TestFindNearestNumber2(t *testing.T) {
	t.Skip()
	testArr := []int{1, 3, 8, 9, 8}
	testArr, err := FindNearestNumber(testArr)
	if err != nil {
		t.Errorf("findd err")
	}
	//fmt.Println(testArr)
}
