package _0211020

func findKthNumber(n int, k int) int {
	var count func(prefix int) int
	count = func(prefix int) int {
		now, next, res := prefix, prefix+1, 0
		for now <= n {
			res += min(n+1, next) - now
			now *= 10
			next *= 10
		}
		return res
	}

	preNum := 1
	now := 1
	for preNum < k {
		if preNum+count(now) > k {
			now *= 10
			preNum++
		} else {
			preNum += count(now)
			now++
		}
	}
	return now
}
