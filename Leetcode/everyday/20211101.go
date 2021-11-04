package everyday

func distributeCandies(candyType []int) int {
	hash := map[int]int{}
	res := 0
	for _, value := range candyType {
		hash[value]++
		if hash[value] == 1 {
			res++
		}
	}
	if len(candyType)/2 > res {
		return res
	}
	return len(candyType) / 2
}
