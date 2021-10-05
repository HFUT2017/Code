package _0210907

type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeKLists 请你将所有链表合并到一个升序链表中，返回合并后的链表。
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	return mergeLists(MergeKLists(lists[0:mid]), MergeKLists(lists[mid:]))
}

func mergeLists(l1, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  -1,
		Next: nil,
	}
	cur := head
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
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return head.Next
}
