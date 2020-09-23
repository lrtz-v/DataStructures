package rotateRight

import "testing"

func TestRotateRight(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	res := rotateRight(head, 2)
	if res == nil {
		t.Fatal("TestRotateRight Failed. Got: nil")
	}
	if res.Val != 4 {
		t.Fatalf("TestRotateRight Failed. Got: %d", res.Val)
	}
}

func TestRotateRight2(t *testing.T) {
	head := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	res := rotateRight(head, 4)
	if res == nil {
		t.Fatal("TestRotateRight Failed. Got: nil")
	}
	if res.Val != 2 {
		t.Fatalf("TestRotateRight Failed. Got: %d", res.Val)
	}
}

func TestRotateRight3(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	res := rotateRight(head, 1)
	if res == nil {
		t.Fatal("TestRotateRight Failed. Got: nil")
	}
	if res.Val != 2 {
		t.Fatalf("TestRotateRight Failed. Got: %d", res.Val)
	}
}
