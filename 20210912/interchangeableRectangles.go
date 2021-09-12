package _0210912

func InterchangeableRectangles(rectangles [][]int) int64 {
	hashtable := map[[2]int]int{}
	res := int64(0)
	for i := 0; i < len(rectangles); i++ {
		div := divisor(rectangles[i][0], rectangles[i][1])
		rectangles[i][0] /= div
		rectangles[i][1] /= div
		if hashtable[[2]int{rectangles[i][0], rectangles[i][1]}] != 0 {
			res += int64(hashtable[[2]int{rectangles[i][0], rectangles[i][1]}])
			hashtable[[2]int{rectangles[i][0], rectangles[i][1]}]++
		} else {
			hashtable[[2]int{rectangles[i][0], rectangles[i][1]}] = 1
		}
	}
	return res
}

func divisor(min, max int) int {
	maxDivisor := 0
	complement := max % min
	if complement != 0 {
		maxDivisor = divisor(complement, min)
	} else {
		maxDivisor = min
	}
	return maxDivisor
}
