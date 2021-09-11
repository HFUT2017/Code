package _0210911

func MiddleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	if fast == nil {
		return nil
	}
	for {
		if fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
			if fast == nil {
				return slow
			}
		} else {
			return slow
		}
	}
}
