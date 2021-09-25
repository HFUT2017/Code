package _0210907

// FindKthNumber 给定整数 n 和 k，找到 1 到 n 中字典序第 k 小的数字。
func FindKthNumber(n, k int) int {
	getcount := func(prefix int) int {
		cur, next := prefix, prefix+1
		count := 0
		for cur <= next {
			count += min(next, n+1) - cur
			cur *= 10
			next *= 10
		}
		return count
	}
	preNum := 1
	parent := 1
	for preNum < k {
		count := getcount(parent)
		if count+preNum > k {
			preNum++
			parent *= 10
		} else {
			preNum += count
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
