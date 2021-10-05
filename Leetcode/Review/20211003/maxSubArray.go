package _0211003

import "math"

func maxSubArray(nums []int) int {
	res:=math.MinInt32
	dp:=make([]int,len(nums))
	for i:=0;i<len(nums);i++{
		if i==0{
			dp[i]=nums[i]
			res=dp[i]
			continue
		}
		dp[i]=max(dp[i-1]+nums[i],nums[i])
		res=max(res,dp[i])
	}
	return res
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}
