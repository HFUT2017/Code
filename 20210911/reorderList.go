package _0210911

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}

func ReorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := MiddleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = ReverseList(l2)
	mergeList(l1, l2)
}
