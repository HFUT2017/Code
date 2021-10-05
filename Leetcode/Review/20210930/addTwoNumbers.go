package _0210930

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	t := 0
	sum := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum, t = (n1+n2+t)%10, (n1+n2+t)/10
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
	}
	if t != 0 {
		cur.Next = &ListNode{Val: t}
	}
	return head.Next
}
