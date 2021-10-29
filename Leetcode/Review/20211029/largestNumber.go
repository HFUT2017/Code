package _0211029

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	res := ""
	sort.Slice(nums, func(i, j int) bool {
		return strconv.Itoa(nums[i])+strconv.Itoa(nums[j]) > strconv.Itoa(nums[j])+strconv.Itoa(nums[i])
	})
	for _, value := range nums {
		res += strconv.Itoa(value)
	}
	if res[0] == '0' {
		return "0"
	}
	return res
}
