package intership

func subArrayRanges(nums []int) int64 {
	res := int64(0)
	for i := 0; i < len(nums); i++ {
		minNum, maxNum := nums[i], nums[i]
		for j := i; j < len(nums); j++ {
			if nums[j] > maxNum {
				maxNum = nums[j]
			} else if nums[j] < minNum {
				minNum = nums[j]
			}
			res += int64(maxNum - minNum)
		}
	}
	return res
}
