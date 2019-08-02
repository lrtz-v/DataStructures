package leetcode

/**
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/

// ListNode node
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 16 ms	5 MB
	var rootNode *ListNode
	nodeList := rootNode
	tmpVal := 0
	for {
		if l1 == nil && l2 == nil && tmpVal > 0 {
			nodeList.Next = &ListNode{Val: tmpVal}
			break
		}
		val1 := 0
		val2 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		tmp := val1 + val2 + tmpVal
		tmpVal = 0

		if tmp >= 10 {
			tmp = tmp % 10
			tmpVal = 1
		}

		if rootNode == nil {
			nodeList = &ListNode{Val: tmp}
			rootNode = nodeList
		} else {
			nodeList.Next = &ListNode{Val: tmp}
			nodeList = nodeList.Next
		}
		if tmpVal == 0 && l1 == nil && l2 == nil {
			break
		}
	}
	return rootNode
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	// 16 ms	5.3 MB
	var resList []int
	tmpVal := 0
	for {
		if l1 == nil && l2 == nil && tmpVal > 0 {
			resList = append(resList, tmpVal)
			break
		}
		val1 := 0
		val2 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		tmp := val1 + val2 + tmpVal
		tmpVal = 0

		if tmp >= 10 {
			tmp = tmp % 10
			tmpVal = 1
		}

		resList = append(resList, tmp)

		if tmpVal == 0 && l1 == nil && l2 == nil {
			break
		}
	}
	return createNodeList(resList)
}

func addTwoNumbers4(l1 *ListNode, l2 *ListNode) *ListNode {
	// 20 ms	4.7 MB
	longList := l1
	shortList := l2
	for {
		if l1 == nil && l2 != nil {
			longList, shortList = shortList, longList
			break
		} else if l2 == nil {
			break
		}
		if l1.Next == nil && l2.Next == nil {
			break
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	l1, l2 = longList, shortList

	for {

		longListNodeVal := 0
		shortListNodeVal := 0

		if longList != nil {
			longListNodeVal = longList.Val
		}
		if shortList != nil {
			shortListNodeVal = shortList.Val
		}

		tmp := longListNodeVal + shortListNodeVal

		if tmp >= 10 {
			tmp = tmp % 10
			if longList.Next != nil {
				longList.Next.Val++
			} else {
				longList.Next = &ListNode{Val: 1}
			}
		}
		longList.Val = tmp

		if longList.Next != nil {
			longList = longList.Next
			if shortList != nil {
				shortList = shortList.Next
			}
		} else {
			break
		}
	}

	return l1
}

func createNodeList(nums []int) *ListNode {
	var rootNode *ListNode
	nodeList := rootNode

	for _, val := range nums {
		if rootNode == nil {
			nodeList = &ListNode{Val: val}
			rootNode = nodeList
		} else {
			nodeList.Next = &ListNode{Val: val}
			nodeList = nodeList.Next
		}
	}
	return rootNode
}

func printNodeList(nodeList *ListNode) {
	tmpList := nodeList

	for {
		if tmpList == nil {
			break
		}
		// fmt.Printf("Val: %d\n", tmpList.Val)
		tmpList = tmpList.Next
	}
	// fmt.Println(".....")
}
