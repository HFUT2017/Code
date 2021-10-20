package _0211020

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		if height[left] < height[right] {
			res = max(res, height[left]*(right-left))
			left++
		} else {
			res = max(res, height[right]*(right-left))
			right--
		}
	}
	return res
}
