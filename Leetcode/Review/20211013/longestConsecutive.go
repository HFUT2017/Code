package _0211013

func longestConsecutive(nums []int) int {
	hash := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		hash[nums[i]] = true
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		if hash[nums[i]+1] == true {
			continue
		}
		cur := nums[i]
		length := 1
		for hash[cur-1] == true {
			cur -= 1
			length++
		}
		res = max(res, length)
	}
	return res
}
