package _0210912

func Subsets(nums []int) [][]int {
	res := [][]int{}
	var dfs func(i int, ans []int)
	dfs = func(i int, ans []int) {
		temp := make([]int, len(ans))
		copy(temp, ans)
		res = append(res, temp)
		for j := i; j < len(nums); j++ {
			ans = append(ans, nums[j])
			dfs(j+1, ans)
			ans = ans[:len(ans)-1]
		}
	}
	dfs(0, []int{})
	return res
}
