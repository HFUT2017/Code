package _0211102

import "sort"

func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	res := 0
	hash := map[int]bool{}
	for _, value := range nums {
		hash[value] = true
	}
	for i := 0; i < len(nums); i++ {
		if hash[nums[i]+1] == true {
			continue
		}
		t := 0
		start := nums[i]
		for hash[start] == true {
			start = start - 1
			t++
		}
		res = max(res, t)
	}
	return res
}
