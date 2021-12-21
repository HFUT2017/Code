package amazon

func findJudge(n int, trust [][]int) int {
	in := make([]int, n+1)
	out := make([]int, n+1)
	for _, t := range trust {
		out[t[0]]++
		in[t[1]]++
	}
	for i := 1; i <= n; i++ {
		if in[i] == n-1 && out[i] == 0 {
			return i
		}
	}
	return -1
}
