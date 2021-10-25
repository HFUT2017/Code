package _0211026

func maxCoins(nums []int) int{
	dp:=make([][]int,len(nums)+2)
	for i:=0;i<len(dp);i++{
		dp[i]=make([]int,len(nums)+2)
	}
	value:=make([]int,len(nums)+2)
	value[0],value[len(value)-1]=1,1
	for i:=0;i<len(nums);i++{
		value[i+1]=nums[i]
	}
	for i:=len(nums)-1;i>=0;i--{
		for j:=i+2;j<=len(nums)+1;j++{
			for k:=i+1;k<j;k++{
				sum:=value[i]*value[k]*value[j]
				sum+=dp[i][k]+dp[k][j]
				dp[i][j]=max(dp[i][j],sum)
			}
		}
	}
	return dp[0][len(nums)+1]
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}
