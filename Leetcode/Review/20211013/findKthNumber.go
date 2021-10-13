package _0211013

func findKthNumber(n int, k int) int {
	var count func(prefix int) int
	count = func(prefix int) int {
		now, next, res := prefix, prefix+1, 0
		for now <= n {
			res += min(next, n+1) - now
			now *= 10
			next *= 10
		}
		return res
	}
	preNum := 1
	now := 1
	for preNum < k {
		num := count(now)
		if num+preNum > k {
			preNum++
			now *= 10
		} else {
			preNum += num
			now++
		}
	}
	return now
}
