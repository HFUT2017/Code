package _0210908

import (
	"fmt"
	"testing"
)

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

func TestReverseBetween(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	head := nums2ListNode(nums)
	left, right := 2, 4
	res := ReverseBetween(head, left, right)
	for res != nil {
		println(res.Val)
		res = res.Next
	}
}

func TestMinDistance(t *testing.T) {
	word1 := "horse"
	word2 := "ros"
	println(MinDistance(word1, word2))
}

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	res := Permute(nums)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%d ", res[i][j])
		}
		println()
	}
}
