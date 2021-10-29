package _0211029

//暴力
//func lexicalOrder(n int) []int {
//	res := []int{}
//	for i := 1; i <= n; i++ {
//		res = append(res, i)
//	}
//	sort.Slice(res, func(i, j int) bool {
//		return strconv.Itoa(res[i]) < strconv.Itoa(res[j])
//	})
//	return res
//}

func lexicalOrder(n int) []int {
	res := []int{}
	var dfs func(val int)
	dfs = func(val int) {
		if val > n {
			return
		}
		res = append(res, val)
		for j := 0; j <= 9; j++ {
			dfs(val*10 + j)
		}
	}
	for i := 1; i < 10; i++ {
		dfs(i)
	}
	return res
}
