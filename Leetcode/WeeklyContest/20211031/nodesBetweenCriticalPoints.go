package _0211031

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	nums := []int{}
	cur := head
	for cur != nil {
		nums = append(nums, cur.Val)
		cur = cur.Next
	}
	res := []int{}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			res = append(res, i)
		} else if nums[i] < nums[i-1] && nums[i] < nums[i+1] {
			res = append(res, i)
		}
	}
	if len(res) < 2 {
		return []int{-1, -1}
	}
	minDistance, maxDistance := math.MaxInt32, res[len(res)-1]-res[0]
	for i := 0; i < len(res)-1; i++ {
		minDistance = min(minDistance, res[i+1]-res[i])
	}
	return []int{minDistance, maxDistance}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
