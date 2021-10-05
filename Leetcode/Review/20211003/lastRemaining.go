package _0211003

func lastRemaining(n int, m int) int {
	dp:=make([]int,n+1)
	for i:=2;i<=n;i++{
		dp[i]=(dp[i-1]+m)%i
	}
	return dp[n]
}
