package _0210917

// NthUglyNumber 我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
func NthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	index2, index3, index5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[index2]*2, min(dp[index3]*3, dp[index5]*5))
		if dp[i] == dp[index2]*2 {
			index2++
		}
		if dp[i] == dp[index3]*3 {
			index3++
		}
		if dp[i] == dp[index5]*5 {
			index5++
		}
	}
	return dp[n]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
