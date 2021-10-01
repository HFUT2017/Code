package _0211001

func reorderList(head *ListNode) {
	mid := findMidList(head)
	next := mid.Next
	l1 := reverseList(next)
	head = mergeTwoList(head, l1)
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	for l1 != nil && l2 != nil {
		l1Temp := l1.Next
		l1.Next = l2
		l2Temp := l2.Next
		l2.Next = l1.Next
		l1 = l1Temp
		l2 = l2Temp
	}
	return l1
}

func findMidList(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
