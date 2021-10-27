package _0211027

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			val := cur.Val
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
