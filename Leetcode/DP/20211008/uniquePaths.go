package _0211008

//错误解法：回溯
//func uniquePaths(m int, n int) int {
//	res:=0
//	var dfs func(i,j int)
//	dfs=func(i,j int){
//		if i<0||j<0||i>=m||j>=n{
//			return
//		}
//		if i==m-1&&j==n-1{
//			res++
//		}
//		dfs(i+1,j)
//		dfs(i,j+1)
//	}
//	dfs(0,0)
//	return res
//}

//正确解法:dp
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j]=dp[i-1][j]+dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
