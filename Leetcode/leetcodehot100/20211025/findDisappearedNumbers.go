package _0211025

func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == i+1 {
			continue
		}
		for nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for index, value := range nums {
		if value != index+1 {
			res = append(res, index+1)
		}
	}
	return res
}
