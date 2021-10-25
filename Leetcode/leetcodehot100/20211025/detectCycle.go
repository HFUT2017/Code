package _0211025

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			cur := head
			for cur != slow {
				slow = slow.Next
				cur = cur.Next
			}
			return slow
		}
	}
	return nil
}
