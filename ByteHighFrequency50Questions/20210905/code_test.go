package _0210905

import "testing"

func nums2ListNode(nums []int) *ListNode {
	var head *ListNode
	var tail *ListNode
	for _, x := range nums {
		if head == nil {
			head = &ListNode{Val: x}
			tail = head
		} else {
			tail.Next = &ListNode{Val: x}
			tail = tail.Next
		}
	}
	return head
}

func TestMergeTwoLists(t *testing.T) {
	nums1 := []int{1, 2, 4}
	nums2 := []int{1, 3, 4}
	l1 := nums2ListNode(nums1)
	l2 := nums2ListNode(nums2)
	head := MergeTwoLists(l1, l2)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}

func TestLongestPalindrome(t *testing.T) {
	var s = "babad"
	println(LongestPalindrome(s))
}

func TestReverseKGroup(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	head := nums2ListNode(nums)
	res := ReverseKGroup(head, 2)
	for res != nil {
		println(res.Val)
		res = res.Next
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	num1 := []int{1, 3}
	num2 := []int{2}
	println(FindMedianSortedArrays(num1, num2))
}
