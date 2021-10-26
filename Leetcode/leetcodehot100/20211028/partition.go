package _0211028

func partition(s string) [][]string {
	dp := make([][]int, len(s))
	res := [][]string{}
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s))
	}
	var check func(i, j int) int
	check = func(i, j int) int {
		if i >= j {
			return 1
		}
		if dp[i][j] != 0 {
			return dp[i][j]
		}
		dp[i][j] = -1
		if s[i] == s[j] {
			dp[i][j] = check(i+1, j-1)
		}
		return dp[i][j]
	}
	splits := []string{}
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(dp) {
			res = append(res, append([]string(nil), splits...))
			return
		}
		for j := i; j < len(s); j++ {
			if check(i, j) > 0 {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	return res
}
