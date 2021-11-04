package everyday

func ByteDance(x float64) float64 {
	left, right := float64(0), float64(x)
	for right-left >= 0.0001 {
		mid := left + (right-left)/2.0
		if x >= mid*mid {
			left = mid
		} else {
			right = mid
		}
	}
	return (left + right) / 2
}
