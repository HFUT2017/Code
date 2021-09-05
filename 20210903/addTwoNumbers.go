package _0210903

type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers 给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。请你将两个数相加，并以相同形式返回一个表示和的链表。
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cnt := 0
	var head *ListNode
	var tail *ListNode
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + cnt
		sum, cnt = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
		if cnt != 0 {
			tail.Next = &ListNode{Val: cnt}
		}
	}
	return head

}
