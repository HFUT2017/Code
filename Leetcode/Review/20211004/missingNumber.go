package _0211004

func missingNumber(nums []int) int {
	for i:=0;i<len(nums);i++{
		if nums[i]==i{
			continue
		}
		for nums[i]<len(nums)&&nums[nums[i]]!=nums[i]{
			nums[nums[i]],nums[i]=nums[i],nums[nums[i]]
		}
	}

	for i:=0;i<len(nums);i++{
		if nums[i]!=i{
			return i
		}
	}
	return len(nums)
}
