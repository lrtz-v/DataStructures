package reversePrintLinkList

// ListNode is definition node of LinkList
type ListNode struct {
     Val int
     Next *ListNode
}
func reversePrint(head *ListNode) []int {
	nums := []int{}

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	reverse(nums)
	return nums
}

func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}