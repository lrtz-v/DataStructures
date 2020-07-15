package reversePrintLinkList

import (
	"strconv"
	"strings"
	"testing"
)

func TestReversePrint(t *testing.T) {

	var root *ListNode
	root = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
			},
		},
	}

	expectedList := []int{2, 3, 1}
	list := reversePrint(root)
	if len(expectedList) != len(list) {
		t.Fatalf("[*] TestReversePrint Error! Expetected length is 3, Got %d", len(list))
	}

	for i := range expectedList {
		if expectedList[i] != list[i] {
			t.Fatalf("[*] TestReversePrint Error! Expetected is [2, 3, 1], Got [%s]", toString(list))
		}
	}

}

func toString(list []int) string {
	var builder strings.Builder

	for _, i := range list {
		builder.WriteString(strconv.Itoa(i))
		builder.WriteString(", ")
	}
	return builder.String()
}
