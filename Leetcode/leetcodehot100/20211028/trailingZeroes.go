package _0211028

func trailingZeroes(n int) int {
	res := 0
	temp := 5
	for n >= temp {
		res += n / temp
		temp *= 5
	}
	return res
}
