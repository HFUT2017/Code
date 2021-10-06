package _0211006

////错误解法，只考虑了连续子序列之和
//func canPartition(nums []int) bool {
//	sum:=0
//	if len(nums)<2{
//		return false
//	}
//	for i:=0;i<len(nums);i++{
//		sum+=nums[i]
//	}
//	if sum%2==1{
//		return false
//	}
//	target:=sum/2
//	sort.Ints(nums)
//	left,right:=0,0
//	for right<len(nums){
//		if target==0{
//			return true
//		}
//		if target>0{
//			target-=nums[right]
//			right++
//		}else if target<0{
//			target+=nums[left]
//			left++
//		}
//	}
//	return false
//
//}


//正确解法：0 1背包问题
func canPartition(nums []int) bool {
	if len(nums)<2{
		return false
	}
	sum,max:=0,0
	for _,value:=range nums{
		sum+=value
		if value>max{
			max=value
		}
	}
	if sum%2==1{
		return false
	}
	target:=sum/2
	if max>target{
		return false
	}
	dp:=make([][]bool,len(nums))
	for i:=0;i<len(dp);i++{
		dp[i]=make([]bool,target+1)
	}
	for i:=0;i<len(dp);i++{
		dp[i][0]=true
	}
	dp[0][nums[0]]=true
	for i:=1;i<len(dp);i++{
		value:=nums[i]
		for j:=1;j<=target;j++{
			if j>=value{
				dp[i][j]=dp[i-1][j]||dp[i-1][j-value]
			}else{
				dp[i][j]=dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][target]
}