package amazon

func trap(height []int) int {
	ans := 0
	leftIndex, rightIndex := 0, len(height)-1
	leftMaxHeight, rightMaxHeight := 0, 0
	for leftIndex < rightIndex {
		if height[leftIndex] < height[rightIndex] {
			if height[leftIndex] > leftMaxHeight {
				leftMaxHeight = height[leftIndex]
			}
			ans += leftMaxHeight - height[leftIndex]
			leftIndex++
		} else {
			if height[rightIndex] > rightMaxHeight {
				rightMaxHeight = height[rightIndex]
			}
			ans += rightMaxHeight - height[rightIndex]
			rightIndex--
		}
	}
	return ans
}
