package _0211005

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	nums1 := nums[1:]
	nums2 := nums[:len(nums)-1]
	res1 := rob(nums1)
	res2 := rob(nums2)
	if res1 > res2 {
		return res1
	}
	return res2
}
