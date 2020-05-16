package reversekgroup

import "testing"

func TestReverseKGroup(t *testing.T) {
	root := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	r := reverseKGroup(&root, 2)

	res := []int{2, 1, 4, 3, 5}
	for _, v := range res {
		if r.Val != v {
			t.Fatalf("[*]TestReverseKGroup Error! Expected %d, Got %d\n", v, r.Val)
		}
		r = r.Next
	}
}
