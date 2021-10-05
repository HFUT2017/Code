package _0210908

// Permute 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以按任意顺序 返回答案。
func Permute(nums []int) [][]int {
	var result [][]int
	var backtrack func(nums, pathNums []int, used []bool)
	backtrack = func(nums, pathNums []int, used []bool) {
		if len(nums) == len(pathNums) {
			tmp := make([]int, len(nums))
			copy(tmp, pathNums)
			result = append(result, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				used[i] = true
				pathNums = append(pathNums, nums[i])
				backtrack(nums, pathNums, used)
				pathNums = pathNums[:len(pathNums)-1]
				used[i] = false
			}
		}
	}
	backtrack(nums, []int{}, make([]bool, len(nums)))
	return result
}
