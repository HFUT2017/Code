package _0211029

func integerBreak(n int) int {
	if n == 2 || n == 3 {
		return n - 1
	}
	var helper func(n int) int
	helper = func(n int) int {
		if n <= 3 {
			return n
		}
		return max(2*helper(n-2), 3*helper(n-3))
	}
	return helper(n)
}
