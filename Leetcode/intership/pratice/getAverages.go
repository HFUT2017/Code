package intership

import "strconv"

func getIndex(index, value int) string {
	return strconv.Itoa(index) + strconv.Itoa(value)
}
func GetAverages(nums []int, k int) []int {
	res := []int{}
	hash := map[string]int{}
	left, right := make([]int, len(nums)), make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			left[i] = nums[i]
		} else {
			left[i] = left[i-1] + nums[i]
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			right[i] = nums[i]
		} else {
			right[i] = right[i+1] + nums[i]
		}
	}
	for i := k; i < len(nums)-k; i++ {
		if i-k-1 >= 0 {
			hash[getIndex(i, nums[i])] = left[i] - left[i-k-1]
		} else {
			hash[getIndex(i, nums[i])] = left[i]
		}
		if i+k+1 < len(nums) {
			hash[getIndex(i, nums[i])] = hash[getIndex(i, nums[i])] + right[i] - right[i+k+1] - nums[i]
		} else {
			hash[getIndex(i, nums[i])] = hash[getIndex(i, nums[i])] + right[i] - nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if i < k || i >= len(nums)-k {
			res = append(res, -1)

		} else {
			res = append(res, hash[getIndex(i, nums[i])]/(2*k+1))
		}
	}
	return res
}
