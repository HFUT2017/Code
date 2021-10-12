package _0211012

func majorityElement(nums []int) int {
	cnt := 0
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			cnt++
			continue
		}
		if cnt == 0 {
			res = nums[i]
		} else {
			cnt--
		}
	}
	cnt = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == res {
			cnt++
		}
	}
	if 2*cnt > len(nums) {
		return res
	}
	return -1
}
