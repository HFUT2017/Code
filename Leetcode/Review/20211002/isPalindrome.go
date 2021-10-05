package _0211002

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x==0{
		return true
	}
	y := 0
	z := x
	for x != 0 {
		y = y*10 + x%10
		x /= 10
	}
	return z == y
}