package _0210908

type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseBetween 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
func ReverseBetween(head *ListNode, left, right int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}
