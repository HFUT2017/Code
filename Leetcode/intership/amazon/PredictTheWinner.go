package amazon

func PredictTheWinner(nums []int) bool {
	var dfs func(temp []int, start, end int, turn int) int
	dfs = func(temp []int, start, end int, turn int) int {
		if start == end {
			return nums[start] * turn
		}
		scoreStart := temp[start]*turn + dfs(temp, start+1, end, -turn)
		scoreEnd := temp[end]*turn + dfs(temp, start, end-1, turn)
		return max(scoreStart*turn, scoreEnd*turn) * turn
	}
	return dfs(nums, 0, len(nums)-1, 1) >= 0
}
