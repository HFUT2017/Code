package _0211008

func divisorGame(n int) bool {
	dp := make([]bool, n+5)
	dp[1] = false
	dp[2] = true
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && dp[i-j] == false {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
