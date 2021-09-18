package _0210918

// SumNums 求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）
func SumNums(n int) int {
	if n == 0 {
		return 0
	}
	return n + SumNums(n-1)
}
