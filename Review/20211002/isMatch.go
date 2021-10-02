package _0211002

func isMatch(s string, p string) {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(p)+1)
	}

}
