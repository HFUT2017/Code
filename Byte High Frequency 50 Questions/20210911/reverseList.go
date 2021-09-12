package _0210911

type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}
