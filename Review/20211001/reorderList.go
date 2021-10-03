package _0211001

func reorderList(head *ListNode)  {
	if head==nil{
		return
	}
	mid:=midSearch(head)
	l1:=head
	l2:=mid.Next
	mid.Next=nil
	l2=reverseR(l2)
	mergeR(l1,l2)
}

func midSearch(head *ListNode) *ListNode{
	slow,fast:=head,head.Next
	for fast!=nil&&fast.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
	}
	return slow
}

func reverseR(head *ListNode) *ListNode{
	var pre *ListNode
	for head!=nil{
		next:=head.Next
		head.Next=pre
		pre=head
		head=next
	}
	return pre
}

func mergeR(l1,l2 *ListNode){
	var l1Temp,l2Temp *ListNode
	for l1!=nil&&l2!=nil{
		l1Temp=l1.Next
		l2Temp=l2.Next
		l1.Next=l2
		l1=l1Temp
		l2.Next=l1
		l2=l2Temp
	}
}
