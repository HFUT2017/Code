package _0211017

func solveNQueens(n int) [][]string {
	if n == 1 {
		return [][]string{{"Q"}}
	}
	if n < 4 {
		return nil
	}
	var (
		k     = n - 1
		Pi1d2 = make([]bool, n)
		Pi1d4 = make([]bool, n<<1-1)
		Pi3d4 = make([]bool, n<<1-1)
		cb    = make([][]byte, n)
		res   [][]string
		dfs   func(int)
	)
	dfs = func(i int) {
		if i == n {
			s := make([]string, n)
			for i := range s {
				s[i] = string(cb[i])
			}
			res = append(res, s)
			return
		}
		for j := 0; j < n; j++ {
			if Pi1d2[j] || Pi1d4[i+j] || Pi3d4[i-j+k] {
				continue
			}
			cb[i][j] = 'Q'
			Pi1d2[j], Pi1d4[i+j], Pi3d4[i-j+k] = true, true, true
			dfs(i + 1)
			cb[i][j] = '.'
			Pi1d2[j], Pi1d4[i+j], Pi3d4[i-j+k] = false, false, false
		}
	}
	for i := range cb {
		cb[i] = make([]byte, n)
		for j := range cb[i] {
			cb[i][j] = '.'
		}
	}
	dfs(0)
	return res
}
