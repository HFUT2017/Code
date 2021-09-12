package _0210912

func Candy(ratings []int) int {
	left := make([]int, len(ratings))
	left[0] = 1
	res, right := 0, 0
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	for i := len(ratings) - 1; i >= 0; i-- {
		if i <len(ratings)-1&& ratings[i] > ratings[i+1] {
			right = right + 1
		} else {
			right = 1
		}
		res += max(left[i], right)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
