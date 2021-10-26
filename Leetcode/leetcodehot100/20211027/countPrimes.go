package _0211027

func countPrimes(n int) int {
	dp := make([]bool, n+1)
	res := 0
	for i := 2; i <= n; i++ {
		if dp[i] == true {
			continue
		}
		for j := i + i; j <= n; j += i {
			dp[j] = true
		}
	}
	for i := 2; i < n; i++ {
		if dp[i] == false {
			res++
		}
	}
	return res
}
