package _0211028

func countSmaller(nums []int) []int {
	var c []int
	update := func(x, v int) {
		for ; x < len(c); x += x & -x {
			c[x] += v
		}
	}
	query := func(x int) int {
		res := 0
		for ; x > 0; x -= x & -x {
			res += c[x]
		}
		return res
	}
	c = make([]int, 10000*2+1)
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		x := nums[i] + 10001
		res[i] = query(x - 1)
		update(x, 1)
	}
	return res
}
