package _0211102

func findKthNumber(n int, k int) int {
	count := func(prefix int) int {
		now, next := prefix, prefix+1
		sum := 0
		for now <= n {
			sum += min(next, n+1) - now
			now *= 10
			next *= 10
		}
		return sum
	}

	prenum := 1
	now := 1
	for prenum < k {
		num := count(now)
		if prenum+num > k {
			prenum++
			now *= 10
		} else {
			prenum += num
			now++
		}
	}
	return now
}
