package BDA2

import "testing"

func TestMergeTwoNUms(t *testing.T) {
	nums1:=[]int{1,2,4}
	nums2:=[]int{3,5,6}
	nums:=MergeTwoNUms(nums1,nums2)
	for _,v:=range nums{
		println(v)
	}
}


func TestFindKElement(t *testing.T) {
	nums:=[]int{5,1,3,4,2}
	println(FindKElement(nums,3))
}