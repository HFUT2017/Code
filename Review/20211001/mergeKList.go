package _0211001

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKList(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	return merge(mergeKList(lists[:mid]), mergeKList(lists[mid:]))
}

func merge(l1, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	return dummyHead.Next
}
