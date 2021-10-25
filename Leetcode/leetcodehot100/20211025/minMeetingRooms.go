package _0211025

import "sort"

func minMeetingRooms(intervals [][]int) int {
	nums := []int{}
	for _, value := range intervals {
		nums = append(nums, value[0]*10+2)
		nums = append(nums, value[1]*10+1)
	}
	sort.Ints(nums)
	res := 0
	nowNeedRoom := 0
	for _, value := range nums {
		if value%10 == 1 {
			nowNeedRoom--
		} else {
			nowNeedRoom++
		}
		if nowNeedRoom > res {
			res = nowNeedRoom
		}
	}
	return res
}
