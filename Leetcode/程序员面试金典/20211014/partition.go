package _0211014

func partition(head *ListNode, x int) *ListNode {
	dummyHead := &ListNode{}
	left := dummyHead
	dummyHeadRight := &ListNode{}
	right := dummyHeadRight
	for head != nil {
		if head.Val < x {
			left.Next = &ListNode{Val: head.Val, Next: nil}
			left = left.Next
		} else {
			right.Next = &ListNode{Val: head.Val, Next: nil}
			right = right.Next
		}
		head = head.Next
	}
	left.Next = dummyHeadRight.Next
	return dummyHead.Next
}
