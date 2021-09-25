package _0210906

import "math"

// Reverse 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
func Reverse(x int) int {
	var res int
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		dight := x % 10
		x /= 10
		res = res*10 + dight
	}
	return res
}
