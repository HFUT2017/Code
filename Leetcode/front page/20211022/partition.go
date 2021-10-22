package _0211022

func partition(head *ListNode, x int) *ListNode {
	leftHead, rightHead := &ListNode{}, &ListNode{}
	left, right := leftHead, rightHead
	cur := head
	for cur != nil {
		if cur.Val < x {
			left.Next = &ListNode{Val: cur.Val}
			left = left.Next
		} else {
			right.Next = &ListNode{Val: cur.Val}
			right = right.Next
		}
		cur = cur.Next
	}
	left.Next = rightHead.Next
	return leftHead.Next
}
