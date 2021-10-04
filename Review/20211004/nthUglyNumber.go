package _0211004

func nthUglyNumber(n int) int {
	dp:=make([]int,n+1)
	dp[1]=1
	index2,index3,index5:=1,1,1
	for i:=2;i<=n;i++{
		dp[i]=min(dp[index2]*2,min(dp[index3]*3,dp[index5]*5))
		if dp[i]==dp[index2]*2{
			index2++
		}
		if dp[i]==dp[index3]*3{
			index3++
		}
		if dp[i]==dp[index5]*5{
			index5++
		}
	}
	return dp[n]
}

func min(a,b int) int{
	if a<b{
		return a
	}
	return b
}
