package rotateRight

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL

1-2-3-4-5
5-1-2-3-4
4-5-1-2-3

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL

0-1-2
2-0-1
1-2-0
0-1-2

输入: 1->2->NULL, k = 1s
输出: 2->1->NULL

1-2
2-1
1-2

*/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	} else if k == 0 {
		return head
	}

	tail := head
	n := 1
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	tail.Next = head

	for i :=  n - k % n; i > 0; i-- {
		tail, head = tail.Next, head.Next
	}
	tail.Next = nil

	return head
}
