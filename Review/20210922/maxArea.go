package _0210922

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left, right, leftIndex, rightIndex, maxArea := height[0], height[len(height)-1], 0, len(height)-1, 0
	for leftIndex < rightIndex {
		maxArea = max(maxArea, (rightIndex-leftIndex)*min(left, right))
		if left < right {
			leftIndex++
			left = height[leftIndex]
		} else {
			rightIndex--
			right = height[rightIndex]
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
