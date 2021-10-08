package _0211007

//给你一个二进制字符串数组 strs 和两个整数 m 和 n 。 请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
//func findMaxForm(strs []string, m int, n int) int {
//	zero,one:=0,0
//	res:=0
//	var getCount func(str string)(int,int)
//	getCount=func(str string)(int,int){
//		zero,one:=0,0
//		for i:=0;i<len(str);i++{
//			if str[i]=='0'{
//				zero++
//			}else{
//				one++
//			}
//		}
//		return zero,one
//	}
//
//	var check func(str string) bool
//	check=func(str string) bool{
//		zeroT,oneT:=getCount(str)
//		if zero-zeroT>=m&&one-oneT>=n{
//			return true
//		}
//		return false
//	}
//
//	left,right:=0,0
//	for right<len(strs){
//		zeroT,oneT:=getCount(strs[right])
//		right++
//		zero,one=zero+zeroT,one+oneT
//		if zero>=m&&one>=n{
//			for left<right&&check(strs[left]){
//				zeroTT,oneTT:=getCount(strs[left])
//				zero,one=zero-zeroTT,one-oneTT
//				left++
//			}
//			res=right-left
//		}
//	}
//	return res
//}

//三维动态规划
func findMaxForm(strs []string, m int, n int) int {
	var getCount func(str string) (int, int)
	getCount = func(str string) (int, int) {
		zero, one := 0, 0
		for i := 0; i < len(str); i++ {
			if str[i] == '0' {
				zero++
			} else {
				one++
			}
		}
		return zero, one
	}

	dp := make([][][]int, len(strs)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i := 0; i < len(strs); i++ {
		zero, one := getCount(strs[i])
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i+1][j][k] = dp[i][j][k]
				if j >= zero && k >= one {
					dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j-zero][k-one]+1)
				}
			}
		}
	}

	return dp[len(strs)][m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

