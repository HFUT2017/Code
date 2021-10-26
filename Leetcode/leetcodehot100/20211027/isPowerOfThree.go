package _0211027

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	for {
		if n%3 == 0 && n != 0 {
			n /= 3
			if n == 1 {
				return true
			}
		} else {
			return false
		}
	}
}
