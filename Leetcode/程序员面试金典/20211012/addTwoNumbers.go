package _0211012

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Next: nil}
	cur := dummyHead
	sum, cnt := 0, 0
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
		sum, cnt = (n1+n2+cnt)%10, (n1+n2+cnt)/10
		cur.Next = &ListNode{Val: sum, Next: nil}
		cur = cur.Next
	}
	if cnt != 0 {
		cur.Next = &ListNode{Val: cnt, Next: nil}
	}
	return dummyHead.Next
}
