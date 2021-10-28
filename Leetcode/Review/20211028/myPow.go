package _0211028

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1.0 / x
		n *= -1
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	sum := float64(1)
	for n != 0 {
		if n&1 == 1 {
			sum *= x
		}
		n >>= 1
		x *= x
	}
	return sum
}
