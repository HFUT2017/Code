package _0210904

import "testing"

func TestTrap(t *testing.T) {
	var height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	println(Trap(height))
}

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

func TestReverseList(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5}
	head := nums2ListNode(nums)
	res := ReverseList(head)
	for res != nil {
		println(res.Val)
		res = res.Next
	}
}

func TestLRUCache(t *testing.T) {
	lru := Constructor(3)
	lru.Put(1, 1)
	println(lru.Get(1))
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	println(lru.Get(1))
}
