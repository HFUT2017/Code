package _0210906

// MaxArea 在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
func MaxArea(height []int) int {
	left, right, res := 0, len(height)-1, 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		res = max(res, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
