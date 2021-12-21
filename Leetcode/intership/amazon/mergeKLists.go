package amazon

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	return mergeTwoLists(mergeKLists(lists[:mid]), mergeKLists(lists[mid:]))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur, cur1, cur2 := dummyHead, l1, l2
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			cur.Next = &ListNode{Val: cur1.Val}
			cur1 = cur1.Next
		} else {
			cur.Next = &ListNode{Val: cur2.Val}
			cur2 = cur2.Next
		}
		cur = cur.Next
	}
	if cur1 != nil {
		cur.Next = cur1
	} else {
		cur.Next = cur2
	}
	return dummyHead.Next
}
