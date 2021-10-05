package _0211005

func fib(n int) int {
	dp:=make([]int,n+1)
	if n<=1{
		return n
	}
	dp[0]=0
	dp[1]=1
	for i:=2;i<=n;i++{
		dp[i]=dp[i-1]+dp[i-2]
	}
	return dp[n]
}
