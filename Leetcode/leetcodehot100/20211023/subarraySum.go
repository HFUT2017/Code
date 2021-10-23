package _0211023

func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	hash := map[int]int{}
	hash[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := hash[pre-k]; ok {
			count += hash[pre-k]
		}
		hash[pre] += 1
	}
	return count
}
