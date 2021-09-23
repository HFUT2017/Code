package _0210922

func permute(nums []int) [][]int {
	used := make([]bool, len(nums))
	var dfs func(nums []int, temp []int, used []bool)
	res := [][]int{}
	dfs = func(nums []int, temp []int, used []bool) {
		if len(temp) == len(nums) {
			t := make([]int, len(temp))
			copy(t, temp)
			res = append(res, t)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] == true {
				continue
			}
			used[i] = true
			temp = append(temp, nums[i])
			dfs(nums, temp, used)
			used[i] = false
			temp = temp[:len(temp)-1]
		}
	}
	dfs(nums, []int{}, used)
	return res
}
