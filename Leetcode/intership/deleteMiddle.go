package intership

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	len := 0
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := head
	for cur != nil {
		len++
		cur = cur.Next
	}
	if len <= 1 {
		return nil
	}
	cur = head
	len = len/2 - 1
	for len != 0 {
		cur = cur.Next
		len--
	}
	cur.Next = cur.Next.Next
	return dummyHead.Next
}
