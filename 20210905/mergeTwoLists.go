package _0210905

type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoLists 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	var head = &ListNode{
		Val:  -1,
		Next: nil,
	}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
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
