package _0211029

func candy(ratings []int) int {
	res := 0
	left := make([]int, len(ratings))
	right := 0
	for i := 0; i < len(ratings); i++ {
		if i == 0 {
			left[i] = 1
			continue
		}
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	for i := len(ratings) - 1; i >= 0; i-- {
		if i == len(ratings)-1 {
			right = 1
		} else {
			if ratings[i] > ratings[i+1] {
				right++
			} else {
				right = 1
			}
		}
		res += max(right, left[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
