package _0211003

func findRepeatNumber(nums []int) int {
	for i:=0;i<len(nums);i++{
		if nums[i]==i{
			continue
		}
		if nums[nums[i]]==nums[i]{
			return nums[i]
		}
		for nums[i]!=i&&nums[nums[i]]!=nums[i]{
			nums[nums[i]],nums[i]=nums[i],nums[nums[i]]
		}
	}
	return 0
}
