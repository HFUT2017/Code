package intership

func IntegerReplacement(n int) int {
	var dfs func(n int) int
	dfs = func(n int) int {
		if n == 1 || n == 0 {
			return 0
		}
		if n%2 == 0 {
			return dfs(n/2) + 1
		} else {
			return min(dfs(n+1)+1, dfs(n-1)+1)
		}
	}
	return dfs(n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
