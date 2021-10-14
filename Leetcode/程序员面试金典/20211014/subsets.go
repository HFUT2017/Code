package _0211014

func subsets(nums []int) [][]int {
	res := [][]int{}
	var backTrack func(temp []int, index int)
	backTrack = func(temp []int, index int) {
		t := make([]int, len(temp))
		copy(t, temp)
		res = append(res, t)
		for i := index; i < len(nums); i++ {
			temp = append(temp, nums[i])
			backTrack(temp, i+1)
			temp = temp[:len(temp)-1]
		}
	}
	backTrack([]int{}, 0)
	return res
}
