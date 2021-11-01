package _0211101

func findTargetSumWays(nums []int, target int) int {
	var dfs func(index int, sum int)
	res := 0
	dfs = func(index int, sum int) {
		if sum == target {
			res++
		}
		if index == len(nums) {
			return
		}
		dfs(index+1, sum+nums[index])
		dfs(index+1, sum-nums[index])
	}
	dfs(0, 0)
	return res
}
