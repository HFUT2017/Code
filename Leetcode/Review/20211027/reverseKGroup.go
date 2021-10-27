package _0211027

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{Next: head}
	pre := dummyHead
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummyHead.Next
			}
		}
		next := tail.Next
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return dummyHead.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
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
