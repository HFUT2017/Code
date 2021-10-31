package _0211031

func mySqrt(x int) int {
	left, right := 0, x
	res := 0
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
