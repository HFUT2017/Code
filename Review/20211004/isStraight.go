package _0211004

import "math"

func isStraight(nums []int) bool {
	max,min:=math.MinInt32,math.MaxInt32
	hash:=map[int]int{}
	for i:=0;i<len(nums);i++{
		hash[nums[i]]++
		if (hash[nums[i]]>1&&nums[i]!=0){
			return false
		}
		if nums[i]>max{
			max=nums[i]
		}
		if nums[i]<min&&nums[i]!=0{
			min=nums[i]
		}
	}
	return max-min<5
}