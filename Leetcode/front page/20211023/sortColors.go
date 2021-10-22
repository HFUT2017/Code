package _0211023

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	for i := 0; i <= right; i++ {
		for ; i <= right && nums[i] == 2; right-- {
			nums[i], nums[right] = nums[right], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
}
