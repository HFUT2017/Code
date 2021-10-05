package _0210906

import "sort"

// ThreeSum 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
func ThreeSum(nums []int) [][]int {
	res:=make([][]int,0)
	sort.Ints(nums)
	for first:=0;first<len(nums);first++{
		if first>0&&nums[first]==nums[first-1]{
			continue
		}
		third:=len(nums)-1
		target:=-1*nums[first]
		for second:=first+1;second<len(nums);second++{
			if second>first+1&&nums[second]==nums[second-1]{
				continue
			}
			for second<third&&nums[second]+nums[third]>target{
				third--
			}
			if second==third{
				break
			}
			if nums[second]+nums[third]==target{
				res=append(res,[]int{nums[first],nums[second],nums[third]})
			}
		}
	}
	return res

}
