package _0211014

type ListNode struct {
	Val  int
	Next *ListNode
}

func kthToLast(head *ListNode, k int) int {
	dummyHead := &ListNode{Next: head}
	fast := dummyHead
	slow := dummyHead
	for k != 0 {
		if fast == nil {
			return -1
		}
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}
