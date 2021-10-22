package _0211023

func combine(n int, k int) [][]int {
	res := [][]int{}
	var dfs func(temp []int, index int)
	dfs = func(temp []int, index int) {
		if len(temp) == k {
			t := make([]int, k)
			copy(t, temp)
			res = append(res, t)
			return
		}
		for i := index; i <= n; i++ {
			temp = append(temp, i)
			dfs(temp, i+1)
			temp = temp[:len(temp)-1]
		}
	}
	dfs([]int{}, 1)
	return res
}
