package _0211027

import "testing"

func nums2List(nums []int) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return dummyHead.Next
}
func TestDeleteDuplicates(t *testing.T) {
	head := nums2List([]int{1, 2, 3, 3, 4, 4, 5})
	res := DeleteDuplicates(head)
	for res != nil {
		println(res.Val)
		res = res.Next
	}
}

func TestDeleteDuplicates2(t *testing.T) {
	head := nums2List([]int{1, 2, 3, 3, 4, 4, 5})
	res := DeleteDuplicates2(head)
	for res != nil {
		println(res.Val)
		res = res.Next
	}
}
