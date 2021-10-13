package _0211013

func permute(nums []int) [][]int {
	res := [][]int{}
	temp := []int{}
	used := make([]bool, len(nums))
	var dfs func()
	dfs = func() {
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
			temp = append(temp, nums[i])
			used[i] = true
			dfs()
			used[i] = false
			temp = temp[:len(temp)-1]
		}
	}
	dfs()
	return res
}
