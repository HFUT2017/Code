package _0210917

import (
	"sort"
	"strconv"
)

func MinNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return strconv.Itoa(nums[i])+strconv.Itoa(nums[j]) < strconv.Itoa(nums[j])+strconv.Itoa(nums[i])
	})
	res := ""
	for i := 0; i < len(nums); i++ {
		res += strconv.Itoa(nums[i])
	}
	return res
}
