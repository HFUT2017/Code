package _0211002

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	mid := findMidList(head)
	return mergeTwoListNode(sortList(head), sortList(mid))
}

//func findMidList(head *ListNode) *ListNode {
//	fast, slow := head, head
//	for fast != nil {
//		if fast.Next != nil {
//			fast = fast.Next.Next
//			slow = slow.Next
//		} else {
//			return slow
//		}
//	}
//	return slow
//}
//
//func mergeTwoListNode(l1, l2 *ListNode) *ListNode {
//	dummyHead := &ListNode{}
//	cur := dummyHead
//	for l1 != nil && l2 != nil {
//		if l1.Val < l2.Val {
//			cur.Next = l1
//			l1 = l1.Next
//		} else {
//			cur.Next = l2
//			l2 = l2.Next
//		}
//		cur = cur.Next
//	}
//	if l1 != nil {
//		cur.Next = l1
//	} else {
//		cur.Next = l2
//	}
//	return dummyHead.Next
//}
