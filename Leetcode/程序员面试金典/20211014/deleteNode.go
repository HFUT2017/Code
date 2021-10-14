package _0211014

func deleteNode(node *ListNode) {
	if node == nil || node.Next == nil {
		node = nil
		return
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
