package _0211014

func waysToStep(n int) int {
	if n <= 2 {
		return n
	}
	if n == 3 {
		return 4
	}
	a, b, c := 1, 2, 4
	for i := 4; i <= n; i++ {
		t := (a + b + c) % 1000000007
		a, b, c = b, c, t
	}
	return c
}
