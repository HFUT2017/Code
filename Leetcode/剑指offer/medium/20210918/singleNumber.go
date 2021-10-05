package _0210918

// SingleNumber 在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。
func SingleNumber(nums []int) int {
	hash := map[int]int{}
	for _, value := range nums {
		hash[value]++
	}
	for _, value := range nums {
		if hash[value] == 1 {
			return value
		}
	}
	return -1
}
