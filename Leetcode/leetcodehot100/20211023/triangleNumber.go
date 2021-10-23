package _0211023

import "sort"

func triangleNumber(nums []int) (ans int) {
	sort.Ints(nums)
	for i := len(nums) - 1; i > 0; i-- {
		for k, j := i-1, 0; j < k; k-- {
			for {
				if j < k && nums[k]+nums[j] <= nums[i] {
					j++
				} else {
					break
				}
			}
			ans += k - j
		}
	}
	return ans
}
