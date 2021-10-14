package _0211014

func trap(height []int) int {
	left, right, leftMax, rightMax := 0, len(height)-1, 0, 0
	res := 0
	for left <= right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			}
			res += (leftMax - height[left])
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			}
			res += (rightMax - height[right])
			right--
		}
	}
	return res
}
