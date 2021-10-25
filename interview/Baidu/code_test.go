package Baidu

import "testing"

func nums2List(nums []int) *ListNode{
	dummyHead:=&ListNode{}
	cur:=dummyHead
	for i:=0;i<len(nums);i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur=cur.Next
	}
	return dummyHead.Next
}

func TestReverseListNode(t *testing.T) {
	nums:=[]int{1,2,3,4,5}
	l1:=nums2List(nums)
	l2:=ReverseListNode(l1,2)
	for l2!=nil{
		println(l2.Val)
		l2=l2.Next
	}
}

func reverseString(str string) string{
	res:=""
	for i:=len(str)-1;i>=0;i--{
		res+=string(str[i])
	}
	println(res)

	return res
}
func TestBigIntergerAdd(t *testing.T) {
	nums1:="12400"
	//nums2:="456"
	nums1=reverseString(nums1)
	//nums2=reverseString(nums2)
	//println(BigIntergerAdd(nums1,nums2))
}