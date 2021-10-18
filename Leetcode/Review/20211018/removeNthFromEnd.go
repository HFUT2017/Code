package _0211018

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	node := findKNode(dummyHead, n+1)
	if node == nil {
		return nil
	}
	if node.Next == nil {
		return head
	}
	node.Next = node.Next.Next
	return dummyHead.Next
}

func findKNode(head *ListNode, k int) *ListNode {
	dummyHead := head
	fast := dummyHead
	slow := dummyHead
	for k != 0 {
		if fast == nil {
			return nil
		}
		fast = fast.Next
		k = k - 1
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
