package _0211003

func reversePrint(head *ListNode) []int {
	res:=[]int{}
	for head!=nil{
		res=append(res,head.Val)
		head=head.Next
	}
	for i:=0;i<len(res)/2;i++{
		res[i],res[len(res)-i-1]=res[len(res)-i-1],res[i]
	}
	return res
}
