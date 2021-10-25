package _0211026

func maxSlidingWindow(nums []int, k int) []int {
	q := make([]int, 0, len(nums))
	res := make([]int, 0, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {
		for len(q) > 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if q[0] <= i-k {
			q = q[1:]
		}
		if i+1 >= k {
			res = append(res, nums[q[0]])
		}
	}
	return res
}
