package _0211020

func trap(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			}
			res += leftMax - height[left]
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			}
			res += rightMax - height[right]
			right--
		}
	}
	return res
}
