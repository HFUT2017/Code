package intership

func maxDistance(colors []int) int {
	res := 0
	for i := 0; i < len(colors); i++ {
		for j := len(colors) - 1; j > i; j-- {
			if colors[j] != colors[i] {
				res = max(res, j-i)
			}
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
