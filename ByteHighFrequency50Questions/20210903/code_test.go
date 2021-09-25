package _0210903

import "testing"

func TestTwoSum(t *testing.T) {
	var nums = []int{2, 7, 11, 159}
	var target = 9
	res := TwoSum(nums, target)
	for i := 0; i < len(res); i++ {
		println(res[i])
	}
}

//将数组转换成ListNode形式
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

func TestAddTwoNumbers(t *testing.T) {
	var nums1 = []int{2, 4, 3}
	var nums2 = []int{5, 6, 4}
	l1 := nums2ListNode(nums1)
	l2 := nums2ListNode(nums2)
	res := AddTwoNumbers(l1, l2)
	for {
		if res == nil {
			break
		}
		println(res.Val)
		res = res.Next
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	var s = "abcabcbb"
	println(LengthOfLongestSubstring(s))
}
