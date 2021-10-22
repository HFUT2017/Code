package _0211022

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	length := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		length++
	}
	last := length - k%length
	if last == length {
		return head
	}
	cur.Next = head
	for last != 0 {
		cur = cur.Next
		last--
	}
	res := cur.Next
	cur.Next = nil
	return res
}
