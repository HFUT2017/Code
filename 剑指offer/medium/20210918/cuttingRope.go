package _0210918

// CuttingRope 给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？
func CuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}
	res := 1
	for n > 4 {
		res *= 3
		n -= 3
	}
	return res * n
}
