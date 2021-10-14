package _0211014

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	hash := map[int]bool{head.Val: true}
	pos := head
	for pos.Next != nil {
		cur := pos.Next
		if !hash[cur.Val] {
			hash[cur.Val] = true
			pos = pos.Next
		} else {
			pos.Next = pos.Next.Next
		}
	}
	pos.Next = nil
	return head
}
