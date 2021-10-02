package _0211002

func longestConsecutive(nums []int) int {
	hash := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		hash[nums[i]] = true
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		temp := 0
		for hash[nums[i]-temp] == true {
			temp++
		}
		if temp > res {
			res = temp
		}
	}
	return res
}
