package _0211023

func majorityElement(nums []int) []int {
	hash := map[int]int{}
	ans := map[int]bool{}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++
		if hash[nums[i]] > len(nums)/3 {
			ans[nums[i]] = true
		}
	}
	for index, _ := range ans {
		res = append(res, index)
	}
	return res
}
