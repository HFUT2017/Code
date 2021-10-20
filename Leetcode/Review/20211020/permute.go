package _0211020

func permute(nums []int) [][]int {
	res := [][]int{}
	var dfs func(temp []int)
	used := make([]bool, len(nums))
	dfs = func(temp []int) {
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
			dfs(temp)
			used[i] = false
			temp = temp[:len(temp)-1]
		}
	}
	dfs([]int{})
	return res
}
