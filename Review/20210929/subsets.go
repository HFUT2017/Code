package _0210929

func subsets(nums []int) [][]int {
	res := [][]int{}
	var backTrack func(nums []int, temp []int, index int)
	backTrack = func(nums []int, temp []int, index int) {
		t := make([]int, len(temp))
		copy(t, temp)
		res = append(res, t)
		for i := index; i < len(nums); i++ {
			temp = append(temp, nums[i])
			backTrack(nums, temp, i+1)
			temp = temp[:len(temp)-1]
		}
	}
	backTrack(nums, []int{}, 0)
	return res
}
