package amazon

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		if i == 0 {
			res[i] = 1
		} else {
			res[i] = res[i-1] * nums[i-1]
		}
	}
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			right = 1
		} else {
			right *= nums[i+1]
		}
		res[i] *= right
	}
	return res
}
