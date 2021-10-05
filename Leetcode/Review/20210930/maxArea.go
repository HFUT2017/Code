package _0210930

func maxArea(height []int) int {
	left, right, res := 0, len(height)-1, 0
	for left < right {
		length := min(height[left], height[right])
		if length*(right-left) > res {
			res = length * (right - left)
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
