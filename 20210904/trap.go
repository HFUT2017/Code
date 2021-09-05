package _0210904

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IntMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Trap 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
func Trap(height []int) int {
	res, left, right, maxLeft, maxRight := 0, 0, len(height) - 1, 0, 0
	for left <= right {
		if height[left] <= height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			}else{
				res += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] > maxRight {
				maxRight = height[right]
			}else{
				res += maxRight - height[right]
			}
			right--
		}
	}
	return res
}
