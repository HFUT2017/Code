package amazon

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	cur1, cur2, c := l1, l2, 0
	for cur1 != nil || cur2 != nil {
		n1, n2 := 0, 0
		if cur1 != nil {
			n1 = cur1.Val
			cur1 = cur1.Next
		}
		if cur2 != nil {
			n2 = cur2.Val
			cur2 = cur2.Next
		}
		cur.Next = &ListNode{Val: (n1 + n2 + c) % 10}
		c = (n1 + n2 + c) / 10
		cur = cur.Next
	}
	if c != 0 {
		cur.Next = &ListNode{Val: c}
	}
	return dummyHead.Next
}
