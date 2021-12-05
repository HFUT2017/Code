package intership

func superPow(a int, b []int) int {
	mod := 1337
	tag := a % mod
	res := 1
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != 0 {
			res = res * pow(tag, b[i], mod)
		}
		tag = pow(tag, 10, mod)
	}
	return res
}

func pow(a, b, mod int) int {
	res := 0
	a = a % mod
	for b != 0 {
		if b&1 != 0 {
			res = res * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return res
}
