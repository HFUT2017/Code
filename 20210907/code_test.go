package _0210907

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	println(MaxSubArray(nums))
}

func buildTree(value []int, i int) *TreeNode {
	if i >= len(value) || value[i] == -1 {
		return nil
	}
	node := &TreeNode{
		Val:   value[i],
		Left:  buildTree(value, 2*i+1),
		Right: buildTree(value, 2*i+2),
	}
	return node
}

func TestLevelOrder(t *testing.T) {
	value := []int{3, 9, 20, -1, -1, 15, 7}
	root := buildTree(value, 0)
	res := LevelOrder(root)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%d ", res[i][j])
		}
		println()
	}
}

func TestZigzagLevelOrder(t *testing.T) {
	value := []int{3, 9, 20, -1, -1, 15, 7}
	root := buildTree(value, 0)
	res := ZigzagLevelOrder(root)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%d ", res[i][j])
		}
		println()
	}
}

func TestRightSideView(t *testing.T) {
	value := []int{1, 2, 3, -1, 5, -1, 4}
	root := buildTree(value, 0)
	res := RightSideView(root)
	for i := 0; i < len(res); i++ {
		fmt.Printf("%d ", res[i])
	}
}

func TestFindKthNumber(t *testing.T) {
	println(FindKthNumber(13, 2))
}

func TestNextPermutation(t *testing.T) {
	nums := []int{1, 2, 3}
	NextPermutation(nums)
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}

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

func TestMergeKLists(t *testing.T) {
	lists := []*ListNode{}
	nums1 := []int{1, 4, 5}
	nums2 := []int{1, 3, 4}
	nums3 := []int{2, 6}
	lists = append(lists, nums2ListNode(nums1))
	lists = append(lists, nums2ListNode(nums2))
	lists = append(lists, nums2ListNode(nums3))
	res := MergeKLists(lists)
	for res != nil {
		println(res.Val)
		res = res.Next
	}

}
