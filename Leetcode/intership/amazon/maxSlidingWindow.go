package amazon

func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if queue[0] <= i-k {
			queue = queue[1:]
		}
		if i+1 >= k {
			res = append(res, nums[queue[0]])
		}
	}
	return res
}
