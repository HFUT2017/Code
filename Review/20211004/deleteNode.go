package _0211004

func deleteNode(head *ListNode, val int) *ListNode {
	dummyHead:=&ListNode{Next:head}
	cur:=dummyHead
	for cur.Next.Val!=val{
		cur=cur.Next
	}
	cur.Next=cur.Next.Next
	return dummyHead.Next
}
