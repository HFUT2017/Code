package _0211003

func getKthFromEnd(head *ListNode, k int) *ListNode {
	dummyHead:=&ListNode{Next:head}
	fast:=dummyHead
	for k!=0{
		if fast!=nil{
			fast=fast.Next
			k--
		}else{
			return nil
		}
	}
	slow:=dummyHead
	for fast!=nil{
		fast=fast.Next
		slow=slow.Next
	}
	return slow
}
