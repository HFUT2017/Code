package _0211001

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{Next: head}
	pre := dummyHead
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	tail := dummyHead
	for i := 0; i < right; i++ {
		tail = tail.Next
	}
	next := tail.Next
	pre.Next, tail = reverseB(pre.Next, tail)
	tail.Next = next
	return dummyHead.Next
}

func reverseB(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	cur := head
	for pre != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return tail, head
}
