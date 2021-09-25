package _0210923

//超时
//func findKthNumber(n int, k int) int {
//	ans := []string{}
//	for i := 1; i <= n; i++ {
//		ans = append(ans, strconv.Itoa(i))
//	}
//	tsort.Strings(ans)
//	res, _ := strconv.Atoi(ans[k-1])
//	return res
//}

//字典树
func findKthNumber(n int, k int) int {
	var count = func(prefix int) int {
		now, next := prefix, prefix+1
		num := 0
		for now <= n {
			num += min(next, n+1) - now
			next *= 10
			now *= 10
		}
		return num
	}

	preNum := 1
	parent := 1
	for preNum < k {
		num := count(parent)
		if preNum+num > k {
			preNum++
			parent *= 10
		} else {
			preNum += num
			parent++
		}
	}
	return parent

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}